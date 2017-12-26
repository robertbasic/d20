package handler

import (
	"net/http"
	"html/template"
)

var tpls = []string{
	"template/layout.html",
	"template/home.html",
}
var tpl = template.Must(template.ParseFiles(tpls...))

type Home struct {}

func (h *Home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

