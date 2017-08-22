package main

import (
	"flag"
	. "github.com/zddhub/dit/dit"
)

func main() {
	flag.Parse()

	for _, file := range flag.Args() {
		AddFileToObjects(file)
	}
}
