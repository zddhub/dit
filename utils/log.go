package utils

import (
	"log"
	"os"
)

var (
	LogE = log.New(os.Stderr, "[E] ", log.Lshortfile)
	LogD = log.New(os.Stdout, "[D] ", log.Lshortfile)
	LogI = log.New(os.Stdout, "[I] ", 0)
	LogT = log.New(os.Stdout, "", 0) // output to terminal
)
