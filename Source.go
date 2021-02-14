package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"sync/atomic"
)

var client = http.Client{}
var counter int32
type Urls struct {
	Urls []string `json:"urls"`
}

func asyncHandler(writer http.ResponseWriter, request *http.Request) {
	counter = 0
	decoder := json.NewDecoder(request.Body)
	urls := Urls{}
	decoder.Decode(&urls)
	for _, url := range urls.Urls {
		go makeRequest(url)
	}
}

func syncHandler(writer http.ResponseWriter, request *http.Request) {
	counter = 0
	decoder := json.NewDecoder(request.Body)
	urls := Urls{}
	decoder.Decode(&urls)
	for _, url := range urls.Urls {
		makeRequest(url)
	}
}

func makeRequest(url string) {
	request, ok := http.NewRequest("GET", url, nil)
	atomic.AddInt32(&counter, 1)
	if ok != nil {
		panic(ok)
	}
	client.Do(request)
	fmt.Println(atomic.LoadInt32(&counter))
}

func runAsync() {
	Router := mux.NewRouter()
	Router.HandleFunc("/asyncHandler", asyncHandler)
	Router.HandleFunc("/syncHandler", syncHandler)
	http.ListenAndServe(":8000", Router)
}

func main() {
	runAsync()
}
