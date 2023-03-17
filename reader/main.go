package main

import (
	"os"

	"github.com/garcia-jc/profiling/reader/functional"
)

func main() {
	functional.Work(
		functional.ReadItems("large.json"), os.Stdout,
	)
}
