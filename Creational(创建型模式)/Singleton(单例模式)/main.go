package main

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println("init start")
	NewSingletonInstance()
	fmt.Println("init end")
}

func main() {
	for i := 0; i < 10; i++ {
		go NewSingletonInstance()
	}
	time.Sleep(time.Second * 1)
}
