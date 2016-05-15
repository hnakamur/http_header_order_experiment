package main

import (
	"flag"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Header2", "value2")
	w.Header().Add("X-Header1", "value1")
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	var addr string
	flag.StringVar(&addr, "addr", "127.0.0.1:8080", "listen address")
	flag.Parse()

	http.HandleFunc("/", handler)
	http.ListenAndServe(addr, nil)
}
