package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func Semantic(cFile *os.File, rule int) {
	voidToken := Token{"", "", ""}
	switch rule {
	case 5:
		writeCFile(cFile, rule, voidToken)
	case 6:
		writeCFile(cFile, rule, voidToken)
	case 7, 8:
		id := semanticStack.getStack("ID")
		TIPO := semanticStack.getStack("TIPO")
		symbolTable.Att(Token{id.Class, id.Lexema, TIPO.Lexema})
		updateStack("ID", Token{"L", TIPO.Lexema, "L"})
		writeCFile(cFile, rule, id)
	case 9:
		token := Token{"TIPO", "inteiro", "TIPO"}
		semanticStack.Push(token)
		writeCFile(cFile, rule, voidToken)
	case 10:
		token := Token{"TIPO", "real", "TIPO"}
		semanticStack.Push(token)
		writeCFile(cFile, rule, voidToken)
	case 11:
		token := Token{"TIPO", "literal", "TIPO"}
		semanticStack.Push(token)
		writeCFile(cFile, rule, voidToken)
	case 14:
		//oldToken := semanticStack.Pop()
		//writeCFile(cFile, rule, oldToken)
	case 15:
		lit := semanticStack.getStack("LIT")
		token := Token{"ARG", lit.Lexema, lit.Type}
		updateStack("LIT", token)
	case 16:
		num := semanticStack.getStack("NUM")
		token := Token{"ARG", num.Lexema, num.Type}
		updateStack("NUM", token)
	case 17:
		id := semanticStack.getStack("ID")
		token := symbolTable.Get(id.Lexema)
		if token.Class == "" {
			color.Red("SEMANTIC ERROR - Variable \"%s\" not declared")
			generateFlag = false
			return
		}
		updateStack("ID", Token{"ARG", id.Lexema, id.Type})
	}
}

func openCFile() (f *os.File) {

	f, err := os.Create("PROGRAMA.C")
	if err != nil {
		color.Red("\nERROR: The file PROGRAMA.C cannot be created by the system\n %s", err)
		os.Exit(10)
	}
	return
}

func writeCFile(cFile *os.File, rule int, token Token) {
	if generateFlag == true {
		switch rule {
		case 5:
			fmt.Fprintf(cFile, "\n\n\n")
		case 6:
			fmt.Fprintf(cFile, ";\n")
		case 7:
			fmt.Fprintf(cFile, ", %s", token.Lexema)
		case 8:
			fmt.Fprintf(cFile, token.Lexema)
		case 9:
			fmt.Fprintf(cFile, "int ")
		case 10:
			fmt.Fprintf(cFile, "float ")
		case 11:
			fmt.Fprintf(cFile, "char *")
		case 14:
			fmt.Fprintf(cFile, "printf(%s);\n", token.Lexema)

		}
	}
}

func updateStack(update string, token Token) {
	value := semanticStack.Pop()

	for value.Class != update {
		value = semanticStack.Pop()
	}

	semanticStack.Push(token)
}
