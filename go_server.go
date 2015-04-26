package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	//	"time"
)

import _ "net/http/pprof"

var mutex = &sync.Mutex{}
var enable_lock bool = true

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You are visiting: %s", r.URL.Path)
	v := r.URL.Query()
	if v.Get("print") == "1" {
		fmt.Printf("You are visiting: %s\n", r.URL.Path)
	}
}

func write(data []byte) {
	if enable_lock {
		mutex.Lock()
	}
	error := ioutil.WriteFile("go.tmp", data, 0644)
	if error != nil {
		panic(error)
	}
	if enable_lock {
		mutex.Unlock()
	}
}

func read() (data []byte) {
	data, error := ioutil.ReadFile("go.tmp")
	if error != nil {
		panic(error)
	}
	return data
}

func contention_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You are visiting contention handler: %s", r.URL)
	v := r.URL.Query()
	if v.Get("print") == "1" {
		fmt.Printf("You are visiting: %s\n", r.URL)
	}
	if v.Get("lock") == "0" {
		enable_lock = false
	}
	data := []byte{1}
	write(data)
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	http.HandleFunc("/", handler)
	http.HandleFunc("/contention", contention_handler)
	fmt.Printf("Listening at 8080, debug at 6060\n")
	http.ListenAndServe(":8080", nil)
}
