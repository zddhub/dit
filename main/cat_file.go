package main

import (
	"flag"
	"fmt"
	. "github.com/zddhub/dit/dit"
)

func main() {
	flag.Parse()
	repo := NewRepository()
	_, buffer, _ := repo.CatFile(flag.Args()[0])
	fmt.Printf("%s", buffer)
}
