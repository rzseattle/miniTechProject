package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/test", test)

	http.ListenAndServe(":8090", nil)
}

func test(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}
