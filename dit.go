package dit

import (
	"log"
	"os"
)

var (
	LogE = log.New(os.Stdout, "[E] ", log.Lshortfile)
	LogD = log.New(os.Stdout, "[D] ", log.Lshortfile)
	LogI = log.New(os.Stdout, "[I] ", 0)
)

type dit struct {
}
