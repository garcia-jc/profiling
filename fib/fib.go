package main

import "log"

func main() {
	log.Println(fib(30))
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
