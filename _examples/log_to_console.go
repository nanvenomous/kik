package main

import (
	"errors"
	// "kik" // for testing
	"github.com/mrgarelli/kik"
)

func main() {
	kik.Header("A Kik Example Header")
	kik.Success("something went well")
	kik.Log("LogTag", "example log message")
	err := errors.New("this is a warning")
	kik.WarnIf(err)
	err = errors.New("this is an error")
	kik.FailIf(err, 1)
}
