package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//	"strconv"
	"sync"
	"time"
)

import _ "net/http/pprof"

var mutex = &sync.Mutex{}
var enable_lock bool = true

// make a buffered channel, so that it would accept send even though receiver is not present
var statistic = make(chan int, 1)

func handler(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	if v.Get("print") == "1" {
		fmt.Printf("You are visiting: %s\n", r.URL.Path)
	}
	fmt.Fprintf(w, "You are visiting: %s", r.URL.Path)
}

func total_handler(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	d := <-statistic
	d++
	if v.Get("sleep") != "0" {
		sleep_duration, _ := time.ParseDuration(v.Get("sleep"))
		time.Sleep(sleep_duration)
	}
	if v.Get("print") == "1" {
		fmt.Printf("You are visiting: %s Total: %d\n", r.URL.Path, d)
	}
	fmt.Fprintf(w, "You are visiting: %s Total: %d", r.URL.Path, d)
	statistic <- d
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
	if v.Get("sleep") != "0" {
		sleep_duration, _ := time.ParseDuration(v.Get("sleep"))
		time.Sleep(sleep_duration)
		return
	}
	data := []byte{1}
	write(data)
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	go func() {
		statistic <- 0
	}()
	http.HandleFunc("/", handler)
	http.HandleFunc("/total", total_handler)
	http.HandleFunc("/contention", contention_handler)
	fmt.Printf("Listening at 8080, debug at 6060\n")
	http.ListenAndServe(":8080", nil)
}
