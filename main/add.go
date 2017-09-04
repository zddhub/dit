package main

import (
	"flag"
	. "github.com/zddhub/dit/dit"
)

func main() {
	flag.Parse()
	repo := NewRepository()
	repo.AddFiles(flag.Args()...)
}
