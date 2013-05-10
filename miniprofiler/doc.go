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

/*
Package miniprofiler is a simple but effective mini-profiler for websites.

To use this package, change your HTTP handler functions to use this signature:

    func(*miniprofiler.Profile, http.ResponseWriter, *http.Request)

Register them in the usual way, wrapping them with NewHandler.

Set miniprofiler.Enable to a function that returns true if profiling is enabled:

    miniprofiler.Enable = func(r *http.Request) bool {
        // Filter on the request, perhaps like:
        // return isUserAuthenticated(r)

        // For now, always enable:
        return true
    }

Set miniprofiler.Store and miniprofiler.Get to functions that can get and store
Profile data, perhaps in memory, redis, or a database. Get key is Profile.Id.

Send output of p.Includes(r) to your HTML (it is empty if Enable returns
false).

Step

The Step function can be used to profile more specific parts of your code. It
should be called with the name of the step and a closure:

    p.Step("something", func() {
        // do some work
    })

AddCustomTiming

AddCustomTiming can be used to record any kind of call (redis, RPC, etc.)

    p.AddCustomTiming(
        "redis",       // call type
        "get",         // execute type
        1.0,           // start milliseconds
        5.2,           // duration milliseconds
        "get key_name" // command string
    )

Example

This is a small example using this package.

    package main

    import "fmt"
    import "github.com/mjibson/MiniProfiler/go/miniprofiler"
    import "net/http"

    func Index(p *miniprofiler.Profile, w http.ResponseWriter, r *http.Request) {
        p.Step("something", func() {
            p.AddCustomTiming("RPC", "get", 1.0, 5.2, "get key_name")
        })
        fmt.Fprintf(w, "<html><body>%v</body></html>", p.Includes(r))
    }

    func main() {
        profiles := make(map[string]*miniprofiler.Profile)
        miniprofiler.Enable = func(r *http.Request) bool { return true }
        miniprofiler.Store = func(r *http.Request, p *miniprofiler.Profile) {
            profiles[string(p.Id)] = p
        }
        miniprofiler.Get = func(r *http.Request, id string) *miniprofiler.Profile {
            return profiles[id]
        }
        http.Handle("/", miniprofiler.NewHandler(Index))
        http.ListenAndServe(":8080", nil)
    }

Configuration

Refer to the variables section of the documentation: http://godoc.org/github.com/mjibson/MiniProfiler/go/miniprofiler#_variables.

Other implementations and resources: http://miniprofiler.com.
*/
package miniprofiler
