package miniprofiler

import (
	"html/template"
	"strings"
)

var includesTmpl = parseInclude("include", include_partial_html)
var shareHtml = parseInclude("share", share_html)

func parseInclude(name string, t []byte) *template.Template {
	s := string(t)
	s = strings.Replace(s, "{", "{{.", -1)
	s = strings.Replace(s, "}", "}}", -1)
	return template.Must(template.New(name).Parse(s))
}
