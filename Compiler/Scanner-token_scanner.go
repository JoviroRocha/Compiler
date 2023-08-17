package main

import (
	"bytes"
	"flag"
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

func checkTable(state int, b []byte) (newState int, end int) {
	var char string
	if b[0] == 0 {
		char = "EOF"
	} else if (state != 16 && state != 18) && ((b[0] >= 65 && b[0] <= 90) || (b[0] >= 97 && b[0] <= 122)) {
		char = "letra"
	} else if b[0] >= 48 && b[0] <= 57 {
		char = "numero"
	} else {
		char = string(b)
	}
	newState, exist := stateTable[state][char]
	if !exist {
		return state, 1
	}
	return newState, 0
}

func createToken(b, lexema string, state int) (token Token) {
	if state == 0 {
		token = Token{"ERROR", "NULO", "NULO"}
		color.Red("LEXICAL ERROR - unexpected character \"%s\"\nLine:%d\tColumn:%d\t", b, line, column)
		generateFlag = false
		column++
	}
	if state == 1 {
		token = Token{"ERROR", "NULO", "NULO"}
		color.Red("LEXICAL ERROR - unterminated literal constant \"%s\"\nLine:%d\tColumn:%d\t", lexema, line, column)
		generateFlag = false
	} else if state == 2 {
		token = Token{"LIT", lexema, "literal"}
	} else if state == 3 {
		if !symbolTable.Search(lexema) {
			symbolTable.Put(token)
			token = Token{"ID", lexema, "NULO"}
		} else {
			token = symbolTable.Get(lexema)
		}
	} else if state == 4 {
		token = Token{"ERROR", "NULO", "NULO"}
		color.Red("LEXICAL ERROR - unterminated comment \"%s\"\nLine:%d\tColumn:%d\t", lexema, line, column)
		generateFlag = false
	} else if state == 5 {
		token = Token{"IGNORE", "NULO", "NULO"}
		if lexema == "\n" {
			line++
			column = 0
		}
	} else if state == 6 {
		token = Token{"EOF", "EOF", "NULO"}
	} else if state == 7 || state == 8 || state == 9 {
		token = Token{"OPR", lexema, "NULO"}
	} else if state == 10 {
		token = Token{"RCB", lexema, "NULO"}
	} else if state == 11 {
		token = Token{"OPM", lexema, "NULO"}
	} else if state == 12 {
		token = Token{"AB_P", lexema, "NULO"}
	} else if state == 13 {
		token = Token{"FC_P", lexema, "NULO"}
	} else if state == 14 {
		token = Token{"PT_V", lexema, "NULO"}
	} else if state == 15 {
		token = Token{"VIR", lexema, "NULO"}
	} else if state == 16 {
		token = Token{"NUM", lexema, "INTEGER"}
	} else if state == 17 || state == 19 || state == 20 {
		token = Token{"IGNORE", "NULO", "NULO"}
		color.Red("LEXICAL ERROR - malformed number \"%s\"\nLine:%d\tColumn:%d\t", lexema, line, column)
		generateFlag = false
	} else if state == 18 || state == 21 || state == 22 {
		token = Token{"NUM", lexema, "FLOAT"}
	}
	column += len(lexema)
	return
}

func SCANNER(filePtr *os.File) (token Token) {
	state := 0
	end := 0
	var lexema bytes.Buffer

	for true {
		b := make([]byte, 1)
		_, err := filePtr.Read(b)
		state, end = checkTable(state, b)

		if end == 1 {
			if err == nil && state != 0 {
				filePtr.Seek(-1, 1)
			}
			token = createToken(string(b), lexema.String(), state)
			lexema.Reset()
			if token.Class == "IGNORE" {
				end, state = 0, 0
			} else {
				semanticStack.Push(token)
				return
			}
		} else {
			lexema.Write(b)
		}
	}
	return
}
