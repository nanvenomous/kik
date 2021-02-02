package main

import (
	"errors"

	// "kik" // for testing

	"github.com/mrgarelli/kik"
)

func main() {
	demonstrateFailIf()
}

func demonstrateFailIf() {
	err := errors.New("example error")
	kik.FailIf(err)
}
