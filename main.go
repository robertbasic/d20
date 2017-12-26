package main

import (
	"log"
	"net/http"

	"github.com/robertbasic/d20/handler"
)

func main() {
	http.Handle("/", new(handler.Home))
	http.Handle("/roll", new(handler.Roll))
	http.Handle("/static/", http.FileServer(http.Dir("template/")))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
