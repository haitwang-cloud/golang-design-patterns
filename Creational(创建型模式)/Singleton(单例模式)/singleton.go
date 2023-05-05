package main

import (
	"fmt"
	"sync"
)

type singleton struct {
}

/*
var lock = &sync.Mutex{}

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
*/

var once sync.Once

var singleInstance *singleton

func NewSingletonInstance() *singleton {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now in Once.")
				singleInstance = &singleton{}
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}
