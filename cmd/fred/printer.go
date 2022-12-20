package main

import (
	"fmt"
	"os"
)

var PrintDebug bool = true

type StreamPrinter struct {
	fp *os.File
}

func (x *StreamPrinter) Error(err error, msg string, rest ...interface{}) {
	fmt.Fprintf(x.fp, "[ERROR] %s:\n\t%v\n",
		fmt.Sprintf(msg, rest...), err)
}

func (x *StreamPrinter) Debug(msg string, rest ...interface{}) {
	if !PrintDebug {
		return
	}
	fmt.Fprintf(x.fp, "[DEBUG]:\n\t%s\n",
		fmt.Sprintf(msg, rest...))
}

func (x *StreamPrinter) Out(msg string, rest ...interface{}) {
	fmt.Fprintf(x.fp, fmt.Sprintf(msg, rest...))
}

type FilePrinter struct {
	StreamPrinter
	filename string
}

func NewFilePrinter(filename string) (*FilePrinter, error) {
	fp, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	return &FilePrinter{filename: filename, StreamPrinter: StreamPrinter{fp: fp}}, nil
}

func (x *FilePrinter) Done() { x.fp.Close() }

type ConsolePrinter struct {
	out *StreamPrinter
	err *StreamPrinter
}

func NewConsolePrinter() *ConsolePrinter {
	return &ConsolePrinter{
		out: &StreamPrinter{fp: os.Stdout},
		err: &StreamPrinter{fp: os.Stderr},
	}
}
func (x ConsolePrinter) Error(err error, msg string, rest ...interface{}) {
	x.err.Error(err, msg, rest...)
}
func (x ConsolePrinter) Debug(msg string, rest ...interface{}) { x.out.Debug(msg, rest...) }
func (x ConsolePrinter) Out(msg string, rest ...interface{})   { x.out.Out(msg, rest...) }
func (x ConsolePrinter) Done()                                 { return }
