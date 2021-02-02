package main

import (
	"errors"

	"kik" // for testing
	// "github.com/mrgarelli/kik"
)

func main() {
	demonstrateFailIf()
}

func demonstrateFailIf() {
	kik.Success("something went well")
	err := errors.New("this is a warning")
	kik.WarnIf(err) // will only warn if an error is found (execution continues)
	err = errors.New("example error")
	kik.FailIf(err)
	kik.Log("logTag", "example log message") // will not reach this message
}
