package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strconv"

	"github.com/garcia-jc/profiling/fib/mathx"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		number := r.URL.Query().Get("n")
		n, _ := strconv.Atoi(number)
		log.Println(n)
		out := mathx.Fib(uint(n))
		log.Println(out)
		fmt.Fprint(w, out)
	}
	http.DefaultServeMux.HandleFunc("/fib", handler)
	addr := ":6060"
	log.Println("starting server at", addr)
	log.Println(http.ListenAndServe(addr, nil))
}
