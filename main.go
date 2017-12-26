package main

import (
	"net/http"
	"github.com/robertbasic/d20/handler"
	"log"
)

func main() {
	http.Handle("/", new(handler.Home))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
