package main

import (
	"fmt"
	"os"
)

type Printer interface {
	Error(error, string, ...interface{})
	Debug(string, ...interface{})
}

type ConsolePrinter struct{}

func (x ConsolePrinter) Error(err error, msg string, rest ...interface{}) {
	fmt.Fprintf(os.Stderr, "[ERROR] %s:\n\t%v\n",
		fmt.Sprintf(msg, rest...), err)
}

func (x ConsolePrinter) Debug(msg string, rest ...interface{}) {
	fmt.Fprintf(os.Stdout, "[DEBUG]:\n\t%s\n",
		fmt.Sprintf(msg, rest...))
}
