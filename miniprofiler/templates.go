package miniprofiler

import (
	"html/template"
	"strings"
)

var includePartialHtmlTmpl = parseInclude("include", includePartialHtml)
var shareHtmlTmpl = parseInclude("share", shareHtml)

func parseInclude(name string, t []byte) *template.Template {
	s := string(t)
	s = strings.Replace(s, "{", "{{.", -1)
	s = strings.Replace(s, "}", "}}", -1)
	return template.Must(template.New(name).Parse(s))
}
