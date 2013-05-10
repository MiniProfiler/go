/*
 * Copyright (c) 2013 Matt Jibson <matt.jibson@gmail.com>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package miniprofiler

import (
	"code.google.com/p/tcgl/identifier"
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"runtime/debug"
	"strings"
	"time"
)

func newGuid() string {
	return identifier.NewUUID().String()
}

type Profile struct {
	Id                   string
	Name                 string
	start                time.Time
	Started              int64
	MachineName          string
	Root                 *Timing
	ClientTimings        *ClientTimings
	DurationMilliseconds float64
	CustomLinks          map[string]string

	w http.ResponseWriter
	r *http.Request

	current *Timing
}

// NewProfile creates a new Profile with given name.
// For use only by miniprofiler extensions.
func NewProfile(w http.ResponseWriter, r *http.Request, name string) *Profile {
	p := &Profile{
		CustomLinks: make(map[string]string),
		w:           w,
		r:           r,
	}

	if Enable(r) {
		p.Id = newGuid()
		p.Name = name
		p.start = time.Now()
		p.MachineName = MachineName()
		p.Root = &Timing{
			Id: newGuid(),
		}
		p.current = p.Root

		w.Header().Add("X-MiniProfiler-Ids", fmt.Sprintf("[\"%s\"]", p.Id))
	}

	return p
}

// Finalize finalizes a Profile and Store()s it.
// For use only by miniprofiler extensions.
func (p *Profile) Finalize() {
	if p.Root == nil {
		return
	}

	u := p.r.URL
	if !u.IsAbs() {
		u.Host = p.r.Host
		if p.r.TLS == nil {
			u.Scheme = "http"
		} else {
			u.Scheme = "https"
		}
	}
	p.Root.Name = p.r.Method + " " + u.String()

	p.Started = p.start.Unix()*1000
	p.DurationMilliseconds = Since(p.start)
	p.Root.DurationMilliseconds = p.DurationMilliseconds

	Store(p.r, p)
}

// ProfileFromJson returns a Profile from JSON data.
func ProfileFromJson(b []byte) *Profile {
	p := Profile{}
	json.Unmarshal(b, &p)
	return &p
}

// Json converts a profile to JSON.
func (p *Profile) Json() []byte {
	b, _ := json.Marshal(p)
	return b
}

// Step adds a new child node with given name.
// f should generally be in a closure.
func (p *Profile) Step(name string, f func()) {
	if p.Root != nil {
		t := &Timing{
			Id:                newGuid(),
			Name:              name,
			StartMilliseconds: Since(p.start),
		}
		p.current.Children = append(p.current.Children, t)
		t.parent = p.current
		p.current = t

		f()

		t.DurationMilliseconds = Since(p.start) - t.StartMilliseconds
		p.current = p.current.parent
	} else {
		f()
	}
}

// AddCustomTiming adds a new CustomTiming with given type to the current node.
func (p *Profile) AddCustomTiming(callType, executeType string, start, duration float64, command string) {
	if p.Root == nil {
		return
	}
	t := p.current
	if t.CustomTimings == nil {
		t.CustomTimings = make(map[string][]*CustomTiming)
	}
	s := &CustomTiming{
		Id:                   newGuid(),
		StartMilliseconds:    start,
		DurationMilliseconds: duration,
		CommandString:        html.EscapeString(command),
		StackTraceSnippet:    getStackSnippet(),
		ExecuteType:          executeType,
	}
	t.CustomTimings[callType] = append(t.CustomTimings[callType], s)
}

func getStackSnippet() string {
	stack := debug.Stack()
	lines := strings.Split(string(stack), "\n")
	var snippet []string
	for i := 0; i < len(lines); i++ {
		idx := strings.LastIndex(lines[i], " ")
		if idx == -1 {
			break
		}

		if i+1 < len(lines) && strings.HasPrefix(lines[i+1], "\t") {
			i++
			snip := strings.TrimSpace(lines[i])
			snip = strings.Split(snip, ":")[0]
			sp := strings.Split(snip, ".")
			snip = sp[len(sp)-1]
			if strings.Contains(snip, "miniprofiler") || strings.HasPrefix(snip, "_func_") || snip == "ServeHTTP" || snip == "ProfileRequest" {
				continue
			}
			snippet = append(snippet, snip)
		}
	}

	return strings.Join(snippet[2:], " ")
}

type Timing struct {
	Id                   string
	Name                 string
	DurationMilliseconds float64
	StartMilliseconds    float64
	Children             []*Timing
	CustomTimings        map[string][]*CustomTiming

	parent *Timing
}

type CustomTiming struct {
	Id                             string
	ExecuteType                    string
	CommandString                  string
	StackTraceSnippet              string
	StartMilliseconds              float64
	DurationMilliseconds           float64
	FirstFetchDurationMilliseconds float64
}

type ClientTimings struct {
	RedirectCount int64
	Timings       []*ClientTiming
}

func (c *ClientTimings) Len() int           { return len(c.Timings) }
func (c *ClientTimings) Less(i, j int) bool { return c.Timings[i].Start < c.Timings[j].Start }
func (c *ClientTimings) Swap(i, j int)      { c.Timings[i], c.Timings[j] = c.Timings[j], c.Timings[i] }

type ClientTiming struct {
	Name     string
	Start    int64
	Duration int64
}
