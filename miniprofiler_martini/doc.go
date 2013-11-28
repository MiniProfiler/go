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
Package miniprofiler_martini is a simple but effective mini-profiler for martini.

To use this package, import:

	import mmp "github.com/MiniProfiler/go/miniprofiler_martini"

Add the middleware:

	m.Use(mmp.Profiler())

Now a mmp.Timer object is available via the injector. Call p.Includes() like
normal and add it to your template.

Example

	func main() {
		m := martini.Classic()
		m.Use(mmp.Profiler())
		m.Get("/", func(p mmp.Timer, r *http.Request) string {
			return fmt.Sprintf("<html><body>%v</body></html>", p.Includes())
		})
		m.Run()
	}

See the miniprofiler package docs about further usage: http://godoc.org/github.com/MiniProfiler/go/miniprofiler.
*/
package miniprofiler_martini
