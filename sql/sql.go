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
Package sql provides miniprofiler wrappers for database/sql.

The API is fully compatible with database/sql, but adds some additional timing
functions. To use, replace the database/sql import with this package. Modify
calls to Exec, Query, and QueryRow to their *Timer variants. Pass the
miniprofiler.Timer reference as first argument.

NOTE: this API is experimental and may change.

Example

This is a small example using this package.

	func Index(t miniprofiler.Timer, w http.ResponseWriter, r *http.Request) {
		db, _ := sql.Open("sqlite3", ":memory:")
		db.ExecTimer(t, "create table x(a, b, c)")
		db.ExecTimer(t, "insert into x (1, 2, 4), (3, 5, 6)")
		db.QueryTimer(t, "select * from x")
		fmt.Fprintf(w, `<html><body>%v</body></html>`, t.Includes())
	}
*/
package sql

import (
	"database/sql"

	"github.com/MiniProfiler/go/miniprofiler"
)

type DB struct {
	*sql.DB
}

func Open(driverName, dataSourceName string) (*DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	d := &DB{db}
	return d, err
}

func (d DB) ExecTimer(t miniprofiler.Timer, query string, args ...interface{}) (result sql.Result, err error) {
	t.StepCustomTiming("sql", "exec", query, func() {
		result, err = d.DB.Exec(query, args...)
	})
	return
}

func (d DB) QueryTimer(t miniprofiler.Timer, query string, args ...interface{}) (rows *sql.Rows, err error) {
	t.StepCustomTiming("sql", "query", query, func() {
		rows, err = d.DB.Query(query, args...)
	})
	return
}

func (d DB) QueryRowTimer(t miniprofiler.Timer, query string, args ...interface{}) (row *sql.Row) {
	t.StepCustomTiming("sql", "query", query, func() {
		row = d.DB.QueryRow(query, args...)
	})
	return
}

var ErrNoRows = sql.ErrNoRows
var ErrTxDone = sql.ErrTxDone
var Register = sql.Register

type NullBool struct{ sql.NullBool }
type NullFloat64 struct{ sql.NullFloat64 }
type NullInt64 struct{ sql.NullInt64 }
type NullString struct{ sql.NullString }
type RawBytes struct{ sql.RawBytes }
type Result struct{ sql.Result }
type Row struct{ sql.Row }
type Rows struct{ sql.Rows }
type Scanner struct{ sql.Scanner }
type Stmt struct{ sql.Stmt }
type Tx struct{ sql.Tx }
