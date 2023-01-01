package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logger struct{}

func (logger) Write(b []byte) (int, error) {
	fmt.Println(string(b))
	return len(b), nil
}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// b := make([]byte, 9999999)
	// resp.Body.Read(b)

	// fmt.Println(string(b))
	l := logger{}
	io.Copy(l, resp.Body)
}
