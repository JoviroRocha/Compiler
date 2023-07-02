package main

type Token struct {
	Class  string
	Lexema string
	Type   string
}

var symbolTable HashTable
var stateTable stateTableType
var stateStack state_Stack
var SLR SLRTable
var prods Productions

func main() {

	symbolTable.Start()
	stateTable = startStateTable()
	stateStack.Start()
	SLR = startSLRTable()
	prods.Start()

	filePtr := openFile()
	defer filePtr.Close()
	Parser(filePtr, 1, 1)
}
