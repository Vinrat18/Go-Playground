package main

import (
	"fmt"
	"sync"
)

func mainMap() {
	m := &sync.Map{}
	wg := &sync.WaitGroup{}

	// Put elements
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(index int) {
			m.Store(index, string(index))
			wg.Done()
		}(i)
	}

	// Get element 1
	value, contains := m.Load(1)
	if contains {
		fmt.Printf("%s\n", value.(string))
	}

	// Returns the existing value if present, otherwise stores it
	value, loaded := m.LoadOrStore(3, "three")
	if !loaded {
		fmt.Printf("%s\n", value.(string))
	}

	// Delete element 3
	m.Delete(3)

	// Iterate over all the elements
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("%d: %s\n", key.(int), value.(string))
		return true
	})
}

func main() {
	once := &sync.Once{}
	wg := &sync.WaitGroup{}
	for i := 0; i < 4; i++ {
		i := i
		wg.Add(1)
		go func() {
			once.Do(func() {
				fmt.Printf("first %d\n", i)
			})
			wg.Done()
		}()
	}

	wg.Wait()
}
