package kik

import (
	"fmt"
	"os"
	"runtime/debug"
)

// FailIf fails the script on an error if error exists
func FailIf(err error) {
	if err != nil {
		os.Stderr.Write(debug.Stack())
		errStr := fmt.Sprintf("%s%s%s%s", red, "[ERROR] ", nc, err)
		errByt := []byte(errStr)
		os.Stderr.Write(errByt)
	}
}

// LogIf logs an error if error exists, does not fail script
func LogIf(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

var (
	nc     string
	red    string
	green  string
	yellow string
	blue   string
	purple string
	cyan   string
	white  string
)

func init() {
	nc = "\033[0m"
	red = "\033[31m"
	green = "\033[32m"
	yellow = "\033[33m"
	blue = "\033[34m"
	purple = "\033[35m"
	cyan = "\033[36m"
	white = "\033[37m"
}

// Red prints red
func Red(msg string) {
	fmt.Println(string(red), msg, string(nc))
}

// Green prints green
func Green(msg string) {
	fmt.Println(string(green), msg, string(nc))
}
