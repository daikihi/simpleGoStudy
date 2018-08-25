package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const (
	ping = "/ping"
	pong = "/pong"
	port = 8090
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

func requestHttpGet(path string) {
	time.Sleep(1 * time.Second)
	http.Get("http://localhost:" + strconv.Itoa(port) + path)
}

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.RequestURI
	fmt.Println("request path = " + path)

	switch path {
	case ping:
		fmt.Println("Ping")
		requestHttpGet(pong)
	case pong:
		fmt.Println("Pong")
		requestHttpGet(ping)
	default:
		fmt.Println("no much path comes......" + path)

	}
	fmt.Fprint(w, "hoge hoge")
}
