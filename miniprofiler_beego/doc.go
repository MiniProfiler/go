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
Package miniprofiler_beego is a simple but effective mini-profiler for beego.

To use this package, in init.go, import:

	import mpb "github.com/MiniProfiler/go/miniprofiler_beego"

Add the filters in main.go before beego.Run():

	...
	beego.AddFilter("*", "BeforeRouter", mpb.BeforeRouter)
	beego.AddFilter("*", "AfterExec", mpb.AfterExec)
	...
	beego.Run()

Set the miniprofiler template data in your controller:

	this.Data["miniprofiler"] = this.Ctx.Input.Data["miniprofiler"]

Add {{.miniprofiler}} to your template right before </body>.

For Step and CustomTimer functionality, in your controllers, import:

	"github.com/MiniProfiler/go/miniprofiler"

and add to the top of every Action:

	t := this.Ctx.Input.Data["__miniprofiler"].(miniprofiler.Timer)

Now t is available as a normal miniprofiler.Timer.

See the miniprofiler package docs about further usage: http://godoc.org/github.com/MiniProfiler/go/miniprofiler.
*/
package miniprofiler_beego
