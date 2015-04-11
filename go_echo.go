package main

import (
	"fmt"
	"log"
	"net/http"
)

import _ "net/http/pprof"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You are visiting: %s", r.URL.Path[1:])
	fmt.Printf("You are visiting: %s\n", r.URL.Path[1:])
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
