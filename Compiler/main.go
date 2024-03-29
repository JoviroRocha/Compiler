package main

import "os"

type Token struct {
	Class  string
	Lexema string
	Type   string
}

var symbolTable HashTable
var stateTable stateTableType
var stateStack state_Stack
var semanticStack semantic_Stack
var SLR SLRTable
var prods Productions
var generateFlag bool
var line int
var column int
var temporary int
var tab int

func main() {

	line, column, temporary, tab = 1, 1, 0, 1

	symbolTable.Start()
	stateTable = startStateTable()
	stateStack.Start()
	SLR = startSLRTable()
	prods.Start()

	filePtr := openFile()
	defer filePtr.Close()

	generateFlag = true
	cFile := openCFile()
	defer cFile.Close()

	Parser(filePtr, cFile)

	if generateFlag == false {
		cFile.Close()
		os.Remove("PROGRAMA.C")
	}
}
