package main

import (
	"net/http"
	"github.com/robertbasic/d20/handler"
	"log"
)

func main() {
	http.Handle("/", new(handler.Home))
	http.Handle("/static/", http.FileServer(http.Dir("template/")))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
