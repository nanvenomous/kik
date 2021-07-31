package kik

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
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

const ShellToUse = "bash"

func System(command string) error {
	script := "script --return --quiet -c \"" + command + "\" /dev/null"
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", script)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	out := stdout.String()
	e := stderr.String()
	if out != "" {
		fmt.Print(out)
	}
	if e != "" {
		fmt.Print(e)
	}
	return err
}

func printErr(err error) {
	errStr := fmt.Sprintf("%s%s%s%s", red, "[ERROR] ", nc, err)
	fmt.Println(errStr)
}

// FailIf fails the script on an error if error exists
func FailIf(err error) {
	if err != nil {
		printErr(err)
		fmt.Println(string(debug.Stack()))
		os.Exit(1)
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
