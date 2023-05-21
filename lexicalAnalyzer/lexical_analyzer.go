package main

import "fmt"

type Token struct {
	Classe string
	Lexema string
	Tipo   string
}

func main() {

	symbolTable := HashTable{}
	symbolTable.Start()

	filePtr := openFile()
	line := 1
	column := 1
	token := Token{"", "", ""}
	for token.Classe != "eof" {
		token = SCANNER(filePtr, &line, &column)
		fmt.Printf(" from line: %d and column %d\n", line, column)
	}
	filePtr.Close()
}
