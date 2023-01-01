package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type HttpHandler struct{}

func (HttpHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	data := []byte("Hello")
	response.Write(data)
	io.WriteString(response, "World")
	fmt.Fprintf(response, time.Now().Format(time.UnixDate))
}

func HomePageHandler(response http.ResponseWriter, request *http.Request) {
	data := []byte("Welcome to Homepage")
	response.Write(data)
}

func JsonResponseHandler(resp http.ResponseWriter, req *http.Request) {
	header := resp.Header()

	header.Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)

	jsonSample := map[string]int{"foo": 1, "bar": 2, "baz": 3}

	// a, _ := json.Marshal(jsonSample)
	// fmt.Fprintf(resp, string(a))

	json.NewEncoder(resp).Encode(jsonSample)
}

func handleRequest() {
	fmt.Println("Starting Server")

	mux := http.NewServeMux()

	httpHandler := HttpHandler{}
	mux.Handle("/test", httpHandler)

	mux.HandleFunc("/home", HomePageHandler)

	mux.HandleFunc("/json", JsonResponseHandler)

	rh := http.RedirectHandler("http://google.com", 307)
	mux.Handle("/google", rh)

	log.Fatal(http.ListenAndServe(":8081", mux))
}

func main() {
	handleRequest()
}
