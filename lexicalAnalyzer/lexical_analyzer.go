package main

import (
	"github.com/fatih/color"
)

type Token struct {
	Class  string
	Lexema string
	Type   string
}

var symbolTable HashTable
var stateTable stateTableType

func main() {

	symbolTable.Start()
	stateTable = startStateTable()

	filePtr := openFile()
	line := 1
	column := 1
	token := Token{"", "", ""}
	for token.Class != "EOF" {
		token = SCANNER(filePtr, &line, &column)
		if token.Class != "ERROR" {
			color.Cyan("Class: %s \tLexema: %s \tType: %s\n", token.Class, token.Lexema, token.Type)
		}
	}
	filePtr.Close()
}
