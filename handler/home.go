package handler

import (
	"net/http"
	"html/template"
)

var tpl = template.Must(template.ParseFiles("template/home.html"))

type Home struct {}

func (h *Home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

