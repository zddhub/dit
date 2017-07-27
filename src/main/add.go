package main

import (
	. "dit"
	"flag"
)

func main() {
	flag.Parse()

	for _, file := range flag.Args() {
		var blob Blob
		blob.Hash(file)
	}
}
