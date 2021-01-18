package kik

import (
	"fmt"
	"log"
)

// FailIf fails the script on an error if error exists
func FailIf(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// LogIf logs an error if error exists, does not fail script
func LogIf(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
