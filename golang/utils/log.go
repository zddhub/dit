package utils

import (
	"log"
	"os"
)

var (
	DitDebug = os.Getenv("DIT_DEBUG")
	LogE     = log.New(os.Stderr, "[E] ", log.Lshortfile)
	LogD     = log.New(os.Stdout, "[D] ", log.Lshortfile)
	LogI     = log.New(os.Stdout, "[I] ", 0)
	LogT     = log.New(os.Stdout, "", 0) // output to terminal
)

func init() {
	DitDebug = os.Getenv("DIT_DEBUG")
	if DitDebug != "true" {
		nullFile, _ := os.OpenFile("/dev/null", os.O_RDWR, 0644)
		defer nullFile.Close()

		LogE = log.New(nullFile, "[E] ", log.Lshortfile)
		LogD = log.New(nullFile, "[D] ", log.Lshortfile)
		LogI = log.New(nullFile, "[I] ", 0)
		LogT = log.New(os.Stdout, "", 0) // output to terminal
	}
}
