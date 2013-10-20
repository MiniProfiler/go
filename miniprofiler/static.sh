#!/bin/sh

cd ../ui
bin2go -a -p miniprofiler -s ../miniprofiler/static.go *.html *.css *.js *.tmpl
