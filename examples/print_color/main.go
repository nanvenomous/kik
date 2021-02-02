package main

import (
	"errors"
	"kik" // for testing
	// "github.com/mrgarelli/kik"
)

func main() {
	kik.Header("A Kik Example Header")
	kik.Success("something went well")
	err := errors.New("this is a warning")
	kik.WarnIf(err)
	kik.Log("logTag", "example log message")
}
