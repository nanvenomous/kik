package kik

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"
)

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

func printErr(err error) {
	errStr := fmt.Sprintf("%s%s%s%s", red, "[ERROR] ", nc, err)
	fmt.Fprintln(os.Stderr, errStr)
}

// FailDebugIf fails the script if the err is not nil
// prints the debug stack trace
// exits with the provided exit code
func FailDebugIf(err error, code int) {
	if err != nil {
		printErr(err)
		fmt.Fprintln(os.Stderr, string(debug.Stack()))
		os.Exit(code)
	}
}

// FailIf fails the script if the err is not nil
// exits with the provided exit code
func FailIf(err error, code int) {
	if err != nil {
		printErr(err)
		os.Exit(code)
	}
}

// WarnIf prints an error message with [WARNING] but does not fail execution
func WarnIf(err error) {
	if err != nil {
		warnMessage := fmt.Sprintf("%s%s%s%s", yellow, "[WARNING] ", nc, err)
		fmt.Println(warnMessage)
	}
}

// Log prints the log message after the [tag] in cyan
func Log(tag string, msg string) {
	logMessage := fmt.Sprintf("%s[%s]%s %s", cyan, tag, nc, msg)
	fmt.Println(logMessage)
}

// Success prints a log prefaced by [SUCCESS] in green
func Success(msg string) {
	succMessage := fmt.Sprintf("%s%s%s%s", green, "[SUCCESS] ", nc, msg)
	fmt.Println(succMessage)
}

// Header prints a message in color with a break of underscores
func Header(msg string) {
	color := blue
	headerMessage := fmt.Sprintf("%s%s%s", color, msg, nc)
	headerBreak := strings.Repeat("_", len(msg))
	coloredHeaderBreak := fmt.Sprintf("%s%s%s", color, headerBreak, nc)
	fmt.Println(coloredHeaderBreak)
	fmt.Println(headerMessage)
}
