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

package miniprofiler_beego

import (
	"strings"

	"github.com/MiniProfiler/go/miniprofiler"
	"github.com/astaxie/beego/context"
)

func BeforeRouter(c *context.Context) {
	ok := miniprofiler.Enable(c.Request)
	if ok && strings.HasPrefix(c.Request.URL.Path, miniprofiler.PATH) {
		miniprofiler.MiniProfilerHandler(c.ResponseWriter, c.Request)
		return
	}
	p := miniprofiler.NewProfile(c.ResponseWriter, c.Request, c.Request.URL.Path)
	c.Input.Data["__miniprofiler"] = p
	if ok {
		c.Input.Data["miniprofiler"] = p.Includes()
	}
}

func AfterExec(c *context.Context) {
	d, present := c.Input.Data["__miniprofiler"]
	if !present {
		return
	}
	p, ok := d.(*miniprofiler.Profile)
	if ok {
		p.Finalize()
	}
}
