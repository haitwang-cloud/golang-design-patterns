package main

import (
	"fmt"
	"sync"
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
