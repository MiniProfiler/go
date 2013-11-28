package main

import (
	"fmt"
	"net/http"

	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"github.com/MiniProfiler/go/miniprofiler"
	"github.com/MiniProfiler/go/redis"
	"github.com/MiniProfiler/go/sql"
)

func Index(t miniprofiler.Timer, w http.ResponseWriter, r *http.Request) {
	db, _ := sql.Open("sqlite3", ":memory:")
	db.ExecTimer(t, "create table x(a, b, c)")
	db.ExecTimer(t, "insert into x (1, 2, 4), (3, 5, 6)")
	db.QueryTimer(t, "select * from x")
	t.Step("redissss", func(t miniprofiler.Timer) {
		conn, _ := redis.Dial("tcp", ":6379")
		defer conn.Close()
		conn.DoTimer(t, "set", "tes t", "value")
		conn.SendTimer(t, "get", "test t")
	})
	fmt.Fprintf(w, `<html><body>%v</body></html>`, t.Includes())
}

func main() {
	miniprofiler.TrivialMilliseconds = 0
	http.Handle("/", miniprofiler.NewHandler(Index))
	fmt.Println("serving")
	http.ListenAndServe(":8080", nil)
}
