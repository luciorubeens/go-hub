package main

import (
	"fmt"
	"log"
	"net/http"
)

type ApiHandler struct{}

func (ApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API!")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/api", ApiHandler{})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		fmt.Fprintf(w, "Welcome")
	})
	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
