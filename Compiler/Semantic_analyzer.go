package main

import (
	"fmt"
	"os"
	"strings"

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
		newId := Token{id.Class, id.Lexema, TIPO.Lexema}
		symbolTable.Att(newId)
		updateStack("ID", Token{"L", TIPO.Lexema, "L"})
		writeCFile(cFile, rule, newId)
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
	case 13:
		lexema := semanticStack.getStack("ID").Lexema
		id := symbolTable.Get(lexema)
		if id.Type == "NULO" || id.Type == "" {
			color.Red("SEMANTIC ERROR - Variable \"%s\" not declared\nLine: %d, Column: %d", lexema, line, column)
			generateFlag = false
			return
		} else {
			writeCFile(cFile, rule, id)
		}
	case 14:
		arg := semanticStack.getStack("ARG")
		writeCFile(cFile, rule, arg)
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
		if token.Type == "" {
			color.Red("SEMANTIC ERROR - Variable \"%s\" not declared\nLine: %d, Column: %d", id.Lexema, line, column)
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
	fmt.Fprintf(f, "#include <stdio.h>\n\n#define string_size 100\n\nint main() {\n\n")
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
			if token.Type == "literal" {
				fmt.Fprintf(cFile, ", %s[string_size]", token.Lexema)
			} else {
				fmt.Fprintf(cFile, ", %s", token.Lexema)
			}
		case 8:
			if token.Type == "literal" {
				fmt.Fprintf(cFile, "%s[string_size]", token.Lexema)
			} else {
				fmt.Fprintf(cFile, token.Lexema)
			}
		case 9:
			fmt.Fprintf(cFile, "\tint ")
		case 10:
			fmt.Fprintf(cFile, "\tdouble ")
		case 11:
			fmt.Fprintf(cFile, "\tchar ")
		case 13:
			text := ""
			if token.Type == "literal" {
				text = "\tscanf(\"%s\", " + token.Lexema + ");"
			} else if token.Type == "inteiro" {
				text = "\tscanf(\"%d\", &" + token.Lexema + ");"
			} else if token.Type == "real" {
				text = "\tscanf(\"%lf\", &" + token.Lexema + ");"
			} else {
				color.Red("\nERROR: Internal compiler error\n")
				os.Exit(1)
				return
			}
			fmt.Fprintln(cFile, text)

		case 14:
			if token.Type == "INTEGER" || token.Type == "FLOAT" {
				fmt.Fprintf(cFile, "\tprintf(\"%s\");\n", token.Lexema)
			} else if token.Type == "inteiro" {
				fmt.Fprintf(cFile, "\tprintf(\"%%d\", %s);\n", token.Lexema)
			} else if token.Type == "real" {
				fmt.Fprintf(cFile, "\tprintf(\"%%lf\", %s);\n", token.Lexema)
			} else {
				if strings.Contains(token.Lexema, "\"") {
					fmt.Fprintf(cFile, "\tprintf(%s);\n", token.Lexema)
				} else {
					fmt.Fprintf(cFile, "\tprintf(\"%%s\", %s);\n", token.Lexema)
				}
			}
		}
	}
}

func updateStack(update string, token Token) {
	value := semanticStack.Pop()

	for value.Class != update {
		value = semanticStack.Pop()
	}

	if token.Class == "" && token.Lexema == "" && token.Type == "" {
		return
	}

	semanticStack.Push(token)
}
