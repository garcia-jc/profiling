package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strconv"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		number := r.URL.Query().Get("n")
		n, _ := strconv.Atoi(number)
		log.Println(n)
		out := fib(uint(n))
		log.Println(out)
		fmt.Fprint(w, out)
	}
	http.DefaultServeMux.HandleFunc("/fib", handler)
	addr := ":6060"
	log.Println("starting server at", addr)
	log.Println(http.ListenAndServe(addr, nil))
}

func fib(n uint) uint {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fib(n-1) + fib(n-2)
	}
}
