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
Package redis provides miniprofiler wrappers for github.com/garyburd/redigo/redis.

The API is fully compatible with the redis package, but adds some additional
timing functions. To use, replace the redis import with this package. Modify
calls to Do and Send to their *Timer variants. Pass the miniprofiler.Timer
reference as first argument.

Example

This is a small example using this package.

	func Index(t miniprofiler.Timer, w http.ResponseWriter, r *http.Request) {
		conn, _ := redis.Dial("tcp", ":6379")
		defer conn.Close()
		conn.DoTimer(t, "set", "test", "value")
		fmt.Fprintf(w, `<html><body>%v</body></html>`, t.Includes(r))
	}
*/
package redis

import (
	"fmt"
	"strings"

	"github.com/MiniProfiler/go/miniprofiler"
	"github.com/garyburd/redigo/redis"
)

type Conn struct {
	redis.Conn
}

func Dial(network, address string) (*Conn, error) {
	conn, err := redis.Dial(network, address)
	c := &Conn{conn}
	return c, err
}

func (c Conn) DoTimer(t miniprofiler.Timer, commandName string, args ...interface{}) (reply interface{}, err error) {
	t.StepCustomTiming("redis", "do", format(commandName, args...), func() {
		reply, err = c.Conn.Do(commandName, args...)
	})
	return
}

func (c Conn) SendTimer(t miniprofiler.Timer, commandName string, args ...interface{}) (err error) {
	t.StepCustomTiming("redis", "send", format(commandName, args...), func() {
		err = c.Conn.Send(commandName, args...)
	})
	return
}

func format(commandName string, args ...interface{}) string {
	f := strings.Repeat(` "%v"`, len(args))
	return commandName + fmt.Sprintf(f, args...)
}
