package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct{}

var singleInstance *single

func getInstance() *single {
	lock.Lock()
	defer lock.Unlock()

	if singleInstance == nil {
		fmt.Println("Creating instance")
		singleInstance = &single{}
	}

	return singleInstance
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			getInstance()
			fmt.Println("Obtained instance")
			wg.Done()
		}()
	}

	wg.Wait()
}
