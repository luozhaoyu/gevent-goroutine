package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func simple() {
	//fmt.Println("simple")
	time.Sleep(100 * time.Second)
}

func main() {
	children_number, _ := strconv.Atoi(os.Args[1])
	fmt.Println(os.Args[1:], children_number)
	for i := 0; i < children_number; i++ {
		go simple()
	}
	fmt.Println("main thread sleep 100s...")
	time.Sleep(100 * time.Second)
}
