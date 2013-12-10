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

package miniprofiler_traffic

import (
	"strings"

	"github.com/MiniProfiler/go/miniprofiler"
	"github.com/pilu/traffic"
)

type Middleware struct{}

func (c *Middleware) ServeHTTP(w traffic.ResponseWriter, r *traffic.Request, next traffic.NextMiddlewareFunc) {
	ok := miniprofiler.Enable(r.Request)
	if ok && strings.HasPrefix(r.Request.URL.Path, miniprofiler.PATH) {
		miniprofiler.MiniProfilerHandler(w, r.Request)
		return
	}
	p := miniprofiler.NewProfile(w, r.Request, r.URL.Path)
	w.SetVar("miniprofiler", p.Includes())
	w.SetVar("miniprofiler_timer", p)
	if nextMiddleware := next(); nextMiddleware != nil {
		nextMiddleware.ServeHTTP(w, r, next)
	}
	if ok {
		p.Finalize()
	}
	return
}
