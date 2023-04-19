package main

import (
	"fmt"
	"sync"
	"time"
)

var lock = &sync.Mutex{}

type singleton struct {
}

var singletonInstance *singleton

func NewSingletonInstance() *singleton {
	if singletonInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singletonInstance == nil {
			fmt.Println("create singleton instance")
			singletonInstance = &singleton{}
		} else {
			fmt.Println("singleton instance already exists")
		}
	} else {
		fmt.Println("singleton instance already exists")
	}
	return singletonInstance
}

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
