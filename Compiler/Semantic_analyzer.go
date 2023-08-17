package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func Semantic(cFile *os.File, rule int) {
	switch rule {
	case 5:
		writeCFile(cFile, rule)
	case 6:
		writeCFile(cFile, rule)
	case 7:
		id := semanticStack.getStack("ID")
		TIPO := semanticStack.getStack("TIPO")
		newId := Token{id.Class, id.Lexema, TIPO.Lexema}
		symbolTable.Att(newId)
		semanticStack.updateStack("L", Token{"", "", ""})
		semanticStack.updateStack("VIR", Token{"", "", ""})
		semanticStack.updateStack("ID", Token{"L", TIPO.Lexema, "L"})
		writeCFile(cFile, rule, newId)
	case 8:
		id := semanticStack.getStack("ID")
		TIPO := semanticStack.getStack("TIPO")
		newId := Token{id.Class, id.Lexema, TIPO.Lexema}
		symbolTable.Att(newId)
		semanticStack.updateStack("ID", Token{"L", TIPO.Lexema, "L"})
		writeCFile(cFile, rule, newId)
	case 9:
		token := Token{"TIPO", "inteiro", "TIPO"}
		semanticStack.updateStack("inteiro", token)
		writeCFile(cFile, rule)
	case 10:
		token := Token{"TIPO", "real", "TIPO"}
		semanticStack.updateStack("real", token)
		writeCFile(cFile, rule)
	case 11:
		token := Token{"TIPO", "literal", "TIPO"}
		semanticStack.updateStack("literal", token)
		writeCFile(cFile, rule)
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
		semanticStack.updateStack("leia", Token{"", "", ""})
		semanticStack.updateStack("ID", Token{"", "", ""})
		semanticStack.updateStack("PT_V", Token{"", "", ""})
	case 14:
		arg := semanticStack.getStack("ARG")
		writeCFile(cFile, rule, arg)
		semanticStack.updateStack("escreva", Token{"", "", ""})
		semanticStack.updateStack("ARG", Token{"", "", ""})
		semanticStack.updateStack("PT_V", Token{"", "", ""})
	case 15:
		lit := semanticStack.getStack("LIT")
		token := Token{"ARG", lit.Lexema, lit.Type}
		semanticStack.updateStack("LIT", token)
	case 16:
		num := semanticStack.getStack("NUM")
		token := Token{"ARG", num.Lexema, num.Type}
		semanticStack.updateStack("NUM", token)
	case 17:
		id := semanticStack.getStack("ID")
		token := symbolTable.Get(id.Lexema)
		if token.Type == "" {
			color.Red("SEMANTIC ERROR - Variable \"%s\" not declared\nLine: %d, Column: %d", id.Lexema, line, column)
			generateFlag = false
			return
		}
		semanticStack.updateStack("ID", Token{"ARG", id.Lexema, id.Type})
	case 19:
		ld := semanticStack.getStack("LD")
		id := semanticStack.getStack("ID")
		token := symbolTable.Get(id.Lexema)

		if token.Type == "" {
			color.Red("SEMANTIC ERROR - Variable \"%s\" not declared\nLine: %d, Column: %d", id.Lexema, line, column)
			generateFlag = false
			return
		}
		if ld.Type != token.Type {
			color.Red("SEMANTIC ERROR - Different types for attribution\nLine: %d, Column: %d", line, column)
			generateFlag = false
			return
		}
		writeCFile(cFile, rule, token, ld)
		semanticStack.updateStack("ID", Token{"", "", ""})
		semanticStack.updateStack("RCB", Token{"", "", ""})
		semanticStack.updateStack("LD", Token{"", "", ""})
		semanticStack.updateStack("PT_V", Token{"", "", ""})
	case 20:
		oprd_1 := semanticStack.getStack("OPRD")
		oprd_2 := semanticStack.getStack("OPRD", true)
		opm := semanticStack.getStack("OPM")
		if oprd_1.Type != oprd_2.Type {
			color.Red("SEMANTIC ERROR - Operands with incompatible types\nLine: %d, Column: %d", line, column)
			generateFlag = false
			return
		}
		writeCFile(cFile, rule, oprd_1, opm, oprd_2)
		semanticStack.updateStack("OPRD", Token{"", "", ""})
		semanticStack.updateStack("OPM", Token{"", "", ""})
		temporary_variable := "T" + strconv.Itoa(temporary)
		semanticStack.updateStack("OPRD", Token{"LD", temporary_variable, oprd_1.Type})
		temporary++
	case 21:
		oprd := semanticStack.getStack("OPRD")
		semanticStack.updateStack("OPRD", Token{"LD", oprd.Lexema, oprd.Type})
	case 22:
		id := semanticStack.getStack("ID")
		token := symbolTable.Get(id.Lexema)
		if token.Type == "" {
			color.Red("SEMANTIC ERROR - Variable \"%s\" not declared\nLine: %d, Column: %d", id.Lexema, line, column)
			generateFlag = false
			return
		}
		if token.Type == "literal" {
			color.Red("SEMANTIC ERROR -Cannot use a literal as an operand\nLine: %d, Column: %d", line, column)
			generateFlag = false
			return
		}
		semanticStack.updateStack("ID", Token{"OPRD", token.Lexema, token.Type})
	case 23:
		num := semanticStack.getStack("NUM")
		var token Token
		if num.Type == "FLOAT" {
			token = Token{"OPRD", num.Lexema, "real"}
		} else if num.Type == "INTEGER" {
			token = Token{"OPRD", num.Lexema, "inteiro"}
		}
		semanticStack.updateStack("NUM", token)
	case 25:
		writeCFile(cFile, rule)
	case 26:
		exp_r := semanticStack.getStack("EXP_R")
		writeCFile(cFile, rule, exp_r)
		semanticStack.updateStack("EXP_R", Token{"", "", ""})
	case 27:
		oprd_1 := semanticStack.getStack("OPRD")
		oprd_2 := semanticStack.getStack("OPRD", true)
		opr := semanticStack.getStack("OPR")
		if oprd_1.Type != oprd_2.Type {
			color.Red("SEMANTIC ERROR - Operands with incompatible types\nLine: %d, Column: %d", line, column)
			generateFlag = false
			return
		}
		writeCFile(cFile, rule, oprd_1, opr, oprd_2)
		semanticStack.updateStack("OPRD", Token{"", "", ""})
		semanticStack.updateStack("OPR", Token{"", "", ""})
		temporary_variable := "T" + strconv.Itoa(temporary)
		class := oprd_2.Lexema + opr.Lexema + oprd_1.Lexema
		semanticStack.updateStack("OPRD", Token{"EXP_R", temporary_variable, class})
		temporary++
	case 34:
		exp_r := semanticStack.getStack("EXP_R")
		writeCFile(cFile, rule, exp_r)
		semanticStack.updateStack("EXP_R", Token{"", "", ""})
	case 38:
		writeCFile(cFile, rule)
	case 39:
		writeCFile(cFile, rule)
	}
}

func openCFile() (f *os.File) {

	f, err := os.Create("PROGRAMA.C")
	if err != nil {
		color.Red("\nERROR: The file PROGRAMA.C cannot be created by the system\n %s", err)
		os.Exit(10)
	}
	fmt.Fprintf(f, "#include <stdio.h>\n#include <string.h>\n\ntypedef char literal[256];\n\nint main(void) {\n\n")
	return
}

func writeCFile(cFile *os.File, rule int, token ...Token) {
	if generateFlag == true {
		switch rule {
		case 5:
			fmt.Fprintf(cFile, "\n\n")
		case 6:
			fmt.Fprintf(cFile, ";\n")
		case 7:
			fmt.Fprintf(cFile, ", %s", token[0].Lexema)
		case 8:
			fmt.Fprintf(cFile, token[0].Lexema)
		case 9:
			for x := 0; x < tab; x++ {
				fmt.Fprintf(cFile, "\t")
			}
			fmt.Fprintf(cFile, "int ")
		case 10:
			for x := 0; x < tab; x++ {
				fmt.Fprintf(cFile, "\t")
			}
			fmt.Fprintf(cFile, "double ")
		case 11:
			for x := 0; x < tab; x++ {
				fmt.Fprintf(cFile, "\t")
			}
			fmt.Fprintf(cFile, "literal ")
		case 13:
			for x := 0; x < tab; x++ {
				fmt.Fprintf(cFile, "\t")
			}
			text := ""
			if token[0].Type == "literal" {
				text = "scanf(\"%s\", " + token[0].Lexema + ");"
			} else if token[0].Type == "inteiro" {
				text = "scanf(\"%d\", &" + token[0].Lexema + ");"
			} else if token[0].Type == "real" {
				text = "scanf(\"%lf\", &" + token[0].Lexema + ");"
			} else {
				color.Red("\nERROR: Internal compiler error\n")
				os.Exit(1)
				return
			}
			fmt.Fprintln(cFile, text)
		case 14:
			for x := 0; x < tab; x++ {
				fmt.Fprintf(cFile, "\t")
			}
			if token[0].Type == "INTEGER" || token[0].Type == "FLOAT" {
				fmt.Fprintf(cFile, "printf(\"%s\");\n", token[0].Lexema)
			} else if token[0].Type == "inteiro" {
				fmt.Fprintf(cFile, "printf(\"%%d\", %s);\n", token[0].Lexema)
			} else if token[0].Type == "real" {
				fmt.Fprintf(cFile, "printf(\"%%lf\", %s);\n", token[0].Lexema)
			} else {
				if strings.Contains(token[0].Lexema, "\"") {
					fmt.Fprintf(cFile, "printf(%s);\n", token[0].Lexema)
				} else {
					fmt.Fprintf(cFile, "printf(\"%%s\", %s);\n", token[0].Lexema)
				}
			}
		case 19:
			for x := 0; x < tab; x++ {
				fmt.Fprintf(cFile, "\t")
			}
			fmt.Fprintf(cFile, "%s = %s;\n", token[0].Lexema, token[1].Lexema)
		case 20:
			for x := 0; x < tab; x++ {
				fmt.Fprintf(cFile, "\t")
			}
			if token[0].Type == "inteiro" {
				fmt.Fprintf(cFile, "int T%d;\n", temporary)
			} else {
				fmt.Fprintf(cFile, "double T%d;\n", temporary)
			}
			for x := 0; x < tab; x++ {
				fmt.Fprintf(cFile, "\t")
			}
			fmt.Fprintf(cFile, "T%d = %s %s %s;\n", temporary, token[2].Lexema, token[1].Lexema, token[0].Lexema)
		case 25:
			tab--
			for x := 0; x < tab; x++ {
				fmt.Fprintf(cFile, "\t")
			}
			fmt.Fprintf(cFile, "}\n")
		case 26:
			for x := 0; x < tab; x++ {
				fmt.Fprintf(cFile, "\t")
			}
			fmt.Fprintf(cFile, "if(%s){\n", token[0].Lexema)
			tab++
		case 27:
			for x := 0; x < tab; x++ {
				fmt.Fprintf(cFile, "\t")
			}
			if token[0].Type == "inteiro" {
				fmt.Fprintf(cFile, "int T%d;\n", temporary)
			} else {
				fmt.Fprintf(cFile, "double T%d;\n", temporary)
			}
			for x := 0; x < tab; x++ {
				fmt.Fprintf(cFile, "\t")
			}
			if token[1].Lexema == "=" {
				fmt.Fprintf(cFile, "T%d = %s %s %s;\n", temporary, token[2].Lexema, "==", token[0].Lexema)
			} else if token[1].Lexema == "<>" {
				fmt.Fprintf(cFile, "T%d = %s %s %s;\n", temporary, token[2].Lexema, "!=", token[0].Lexema)
			} else {
				fmt.Fprintf(cFile, "T%d = %s %s %s;\n", temporary, token[2].Lexema, token[1].Lexema, token[0].Lexema)
			}
		case 34:
			for x := 0; x < tab; x++ {
				fmt.Fprintf(cFile, "\t")
			}
			fmt.Fprintf(cFile, "while(%s){\n", token[0].Type)
			tab++
		case 38:
			tab--
			for x := 0; x < tab; x++ {
				fmt.Fprintf(cFile, "\t")
			}
			fmt.Fprintf(cFile, "}\n")
		case 39:
			fmt.Fprintf(cFile, "\treturn 0;\n}\n")
		}
	}
}
