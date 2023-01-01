package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://youtube.com",
		"http://amazon.com",
		"http://gmail.com",
	}

	// c := make(chan string)

	// for _, l := range links {
	// 	go checkLink(l, c)
	// }

	// for l := range c {
	// 	go func(link string) {
	// 		time.Sleep(5 * time.Second)
	// 		checkLink(link, c)
	// 	}(l)
	// }

	wg := &sync.WaitGroup{}

	for _, l := range links {
		wg.Add(1)
		go func(link string) {
			checkLink(link, nil)
			wg.Done()
		}(l)
	}

	wg.Wait()

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "seems to be down!")
	} else {
		fmt.Println(link, "is up!")
	}

	if c != nil {
		c <- link
	}
}
