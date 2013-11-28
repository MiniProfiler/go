package miniprofiler_martini

import (
	"net/http"
	"strings"

	"github.com/MiniProfiler/go/miniprofiler"
	"github.com/codegangsta/martini"
)

type Timer interface {
	miniprofiler.Timer
}

func Profiler() martini.Handler {
	return func(c martini.Context, w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, miniprofiler.PATH) {
			miniprofiler.MiniProfilerHandler(w, r)
			return
		}
		p := miniprofiler.NewProfile(w, r, r.URL.Path)
		c.MapTo(p, (*Timer)(nil))
		c.Next()
		p.Finalize()
	}
}
