package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
)

func openFile() (f *os.File) {

	filePtr := flag.String("file", "", "file path")
	flag.StringVar(filePtr, "f", "", "file path")
	flag.Parse()
	f, err := os.Open(*filePtr)
	if err != nil {
		color.Red("\nERROR: The file \"%s\" cannot be found by the system\n", *filePtr)
		os.Exit(10)
	}
	return
}
func SCANNER(filePtr *os.File, line *int, column *int) (token Token) {
	b := make([]byte, 1)
	_, err := filePtr.Read(b)
	if err == io.EOF {
		token = Token{"eof", "eof", "eof"}
	} else if b[0] == '\n' {
		*line++
		*column = 1
	} else {
		*column++
	}
	fmt.Print(b)
	// Retornar uma posição da carretilha
	_, err = filePtr.Seek(-1, 1)
	return token
}
