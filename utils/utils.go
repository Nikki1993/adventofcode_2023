package utils

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func GetFilePath(args []string) string {
	fPath := "example.txt"
	if len(args) > 1 {
		fPath = args[1]
	}

	return fPath
}

func MustOpenFile(p string) (*os.File, func(file *os.File)) {
	f, err := os.Open(p)
	if err != nil {
		panic(fmt.Errorf("error opening a fPath %v", err))
	}

	return f, func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Printf("Error closing the fPath %v", err)
		}
	}
}

func ScanAndSplit(r io.Reader) *bufio.Scanner {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)

	return s
}
