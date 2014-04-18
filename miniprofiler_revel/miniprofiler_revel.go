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

package miniprofiler_revel

import (
	"strings"

	"github.com/MiniProfiler/go/miniprofiler"
	"github.com/revel/revel"
)

func Filter(c *revel.Controller, fc []revel.Filter) {
	ok := miniprofiler.Enable(c.Request.Request)
	if ok && strings.HasPrefix(c.Request.Request.URL.Path, miniprofiler.PATH) {
		miniprofiler.MiniProfilerHandler(c.Response.Out, c.Request.Request)
		return
	}
	p := miniprofiler.NewProfile(c.Response.Out, c.Request.Request, c.Action)
	c.Args["miniprofiler"] = p
	if ok {
		c.RenderArgs["miniprofiler"] = p.Includes()
	}
	fc[0](c, fc[1:])
	if ok {
		p.SetName(c.Action)
		p.Finalize()
	}
}
