package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func GetTestFilePath(relpath string) string {
	_, b, _, _ := runtime.Caller(0)
	cwd := filepath.Dir(b)
	pth := filepath.Join(cwd, "../testdata", relpath)
	return pth
}

func GetTestFile(relpath string) []byte {
	pth := GetTestFilePath(relpath)
	buffer, err := os.ReadFile(pth)
	if err != nil {
		fmt.Fprintf(os.Stderr, "no such test file: %s: %v", pth, err)
		return []byte{}
	}
	return buffer
}
