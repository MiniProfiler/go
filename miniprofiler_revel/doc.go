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
Package miniprofiler_revel is a simple but effective mini-profiler for revel.

To use this package, in init.go, import:

	import mpr "github.com/MiniProfiler/go/miniprofiler_revel"

and add mpr.Filter above revel.RouterFilter:

	revel.Filters = []revel.Filter{
 		revel.PanicFilter,
		mpr.Filter,
 		revel.RouterFilter,
 		...

Add {{.miniprofiler}} to your footer template right before </body>.

For Step and CustomTimer functionality, in your controllers, import:

	"github.com/MiniProfiler/go/miniprofiler"

and add to the top of every Action:

	t := c.Args["miniprofiler"].(miniprofiler.Timer)

Now t is available as a normal miniprofiler.Timer.

See the miniprofiler package docs about further usage: http://godoc.org/github.com/MiniProfiler/go/miniprofiler.
*/
package miniprofiler_revel
