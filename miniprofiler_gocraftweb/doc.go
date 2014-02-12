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
Package miniprofiler_gocraftweb is a simple but effective mini-profiler for gocraft/web.

To use this package, in init.go, import:

	import mpg "github.com/MiniProfiler/go/miniprofiler_gocraftweb"

Embed mpg.MiniProfilerContext in your context:

	type Context struct {
		mpg.MiniProfilerContext
		...
	}

Add the (*Context).MiniProfilerMiddleware middleware:

	router := web.New(Context{}).
		...
		Middleware((*Context).MiniProfileMiddleware).
		...
		Get("/", (*Context).Index)

Add c.MiniProfilerTemplate right before </body>.

Use c.MiniProfilerTimer as a miniprofiler.Timer.

Example

A full example application:

	package main

	import (
		"fmt"
		"net/http"
		"strings"

		"github.com/MiniProfiler/go/miniprofiler"
		mpg "github.com/MiniProfiler/go/miniprofiler_gocraftweb"
		"github.com/gocraft/web"
	)

	type Context struct {
		mpg.MiniProfilerContext
		HelloCount int
	}

	func (c *Context) SetHelloCount(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
		c.HelloCount = 3
		next(rw, req)
	}

	func (c *Context) SayHello(rw web.ResponseWriter, req *web.Request) {
		c.MiniProfilerTimer.Step("test", func(t miniprofiler.Timer) {
			fmt.Fprint(rw,
				"<html><body>",
				strings.Repeat("Hello ", c.HelloCount),
				"World!",
				c.MiniProfilerTemplate,
				"</body></html>")
		})
	}

	func main() {
		router := web.New(Context{}).
						Middleware(web.LoggerMiddleware).
						Middleware(web.ShowErrorsMiddleware).
						Middleware((*Context).MiniProfileMiddleware).
						Middleware((*Context).SetHelloCount).
						Get("/", (*Context).SayHello)
		http.ListenAndServe("localhost:3000", router)
	}

See the miniprofiler package docs about further usage: http://godoc.org/github.com/MiniProfiler/go/miniprofiler.
*/
package miniprofiler_gocraftweb
