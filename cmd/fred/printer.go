package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
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

func (x *FilePrinter) Done() { x.fp.Close() }

func NewFilePrinter(filename string) (*FilePrinter, error) {
	fp, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	return &FilePrinter{filename: filename, StreamPrinter: StreamPrinter{fp: fp}}, nil
}

type TempFilePrinter struct{ FilePrinter }

func NewTempFilePrinter() (*TempFilePrinter, error) {
	rand.Seed(time.Now().UnixNano())
	size := 16
	letters := "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuWwXxYyZz"
	var filename strings.Builder
	filename.Grow(size)
	for i := 0; i < size; i++ {
		filename.WriteByte(letters[rand.Intn(len(letters))])
	}
	filename.WriteString(".md")
	fp, err := NewFilePrinter(filepath.Join(os.TempDir(), filename.String()))
	if err != nil {
		return nil, err
	}
	return &TempFilePrinter{FilePrinter: *fp}, nil
}

func (x *TempFilePrinter) Done() {
	defer os.RemoveAll(x.filename)
	x.fp.Close()
}

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
