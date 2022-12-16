package main

import (
	"fmt"
	"os"
)

var PrintDebug bool = true

type ConsolePrinter struct{}

func (x ConsolePrinter) Error(err error, msg string, rest ...interface{}) {
	fmt.Fprintf(os.Stderr, "[ERROR] %s:\n\t%v\n",
		fmt.Sprintf(msg, rest...), err)
}

func (x ConsolePrinter) Debug(msg string, rest ...interface{}) {
	if !PrintDebug {
		return
	}
	fmt.Fprintf(os.Stdout, "[DEBUG]:\n\t%s\n",
		fmt.Sprintf(msg, rest...))
}

func (x ConsolePrinter) Out(msg string, rest ...interface{}) {
	fmt.Fprintf(os.Stdout, fmt.Sprintf(msg, rest...))
}
