package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func Parser(filePtr, cFile *os.File) {

	state := 0
	token := SCANNER(filePtr)

	for true {

		if token.Class == "ERROR" {
			token = SCANNER(filePtr)
			continue
		}

		state = stateStack.Get()

		action, exists := SLR[state][token.Class]

		if exists == false {
			action = errorHandler(state, token, filePtr, cFile)
		}

		if newState, err := strconv.Atoi(action); err != nil {

			if action == "ACC" {
				fmt.Println("P' -> P")
				return
			}
			Reduce(action, cFile)
			continue

		} else {
			stateStack.Push(newState)
			token = SCANNER(filePtr)
		}
	}

}

func Reduce(action string, cFile *os.File) {

	var length = len([]rune(action))
	value, _ := strconv.Atoi(action[1:length])
	rule := prods.Get(value)

	stateStack.Pop(rule.size)
	state := stateStack.Get()

	nonTerminal := rule.production[:strings.IndexByte(rule.production, ' ')]
	action, exists := SLR[state][nonTerminal]
	newState, _ := strconv.Atoi(action)
	stateStack.Push(newState)

	fmt.Println(rule.production)

	Semantic(cFile, value)

	trash(exists)
}

func errorHandler(state int, token Token, filePtr, cFile *os.File) (action string) {

	if state == 0 {
		errorPrinter("inicio", token, line, column)
		return errorCorrector(state, token, "inicio", filePtr, cFile)
	} else if state == 1 {
		errorPrinter("EOF", token, line, column)
		follow := [6]string{"EOF", "", "", "", "", ""}
		return errorPanic(follow, filePtr, state)
	} else if state == 2 {
		errorPrinter("varinicio", token, line, column)
		return errorCorrector(state, token, "varinicio", filePtr, cFile)
	} else if state == 4 || state == 19 {
		errorPrinter("varfim\" or \"inteiro\" or \"real\" or \"literal", token, line, column)
		return errorCorrector(state, token, "varfim", filePtr, cFile)
	} else if state == 11 || state == 21 || state == 67 {
		errorPrinter("ID", token, line, column)
		return errorCorrector(state, token, "ID", filePtr, cFile)
	} else if state == 12 {
		errorPrinter("LIT\" or \"NUM\" or \"ID", token, line, column)
		follow := [6]string{"LIT", "NUM", "ID", "", "", ""}
		return errorPanic(follow, filePtr, state)
	} else if state == 13 {
		errorPrinter("RCB", token, line, column)
		return errorCorrector(state, token, "RCB", filePtr, cFile)
	} else if state == 16 || state == 17 {
		errorPrinter("AB_P", token, line, column)
		return errorCorrector(state, token, "AB_P", filePtr, cFile)
	} else if state == 20 || state == 29 || state == 30 || state == 49 || state == 53 {
		errorPrinter("PT_V", token, line, column)
		return errorCorrector(state, token, "PT_V", filePtr, cFile)
	} else if state == 34 || state == 45 || state == 46 || state == 69 || state == 71 {
		errorPrinter("ID\" or \"NUM", token, line, column)
		follow := [6]string{"ID", "NUM", "", "", "", ""}
		return errorPanic(follow, filePtr, state)
	} else if state == 54 {
		errorPrinter("OPM", token, line, column)
		return errorCorrector(state, token, "OPM", filePtr, cFile)
	} else if state == 63 || state == 65 {
		errorPrinter("FC_P", token, line, column)
		return errorCorrector(state, token, "FC_P", filePtr, cFile)
	} else if state == 64 {
		errorPrinter("OPR", token, line, column)
		return errorCorrector(state, token, "OPR", filePtr, cFile)
	} else if state == 70 {
		errorPrinter("entao", token, line, column)
		return errorCorrector(state, token, "entao", filePtr, cFile)
	} else if state == 3 || state == 6 || state == 7 || state == 8 || state == 9 {
		errorPrinter("fim\" or \"leia\" or \"escreva\" or \"ID\" or \"se\" or \"repita", token, line, column)
		if token.Class == "EOF" {
			return errorCorrector(state, token, "fim", filePtr, cFile)
		} else {
			follow := [6]string{"fim", "leia", "escreva", "ID", "se", "repita"}
			return errorPanic(follow, filePtr, state)
		}

	} else if state == 14 || state == 36 || state == 37 || state == 38 {
		errorPrinter("fimse\" or \"leia\" or \"escreva\" or \"ID\" or \"se", token, line, column)
		return errorCorrector(state, token, "fimse", filePtr, cFile)
	} else if state == 15 || state == 41 || state == 42 || state == 43 {
		errorPrinter("fimrepita\" or \"leia\" or \"escreva\" or \"ID\" or \"se", token, line, column)
		return errorCorrector(state, token, "fimrepita", filePtr, cFile)
	}

	panic("should never happen")
}

func errorCorrector(state int, oldToken Token, fixToken string, filePtr, cFile *os.File) (newAction string) {

	action, _ := SLR[state][fixToken]

	if newState, err := strconv.Atoi(action); err != nil {
		Reduce(action, cFile)
	} else {
		stateStack.Push(newState)
	}

	newState := stateStack.Get()
	newAction, exist := SLR[newState][oldToken.Class]

	if exist == false {
		newAction = errorHandler(newState, oldToken, filePtr, cFile)
	}

	return
}

func errorPanic(follow [6]string, filePtr *os.File, state int) (action string) {

	for true {
		token := SCANNER(filePtr)

		for i := 0; i < 6; i++ {
			if follow[i] == token.Class {
				return SLR[state][token.Class]
			}
		}
		if token.Class == "EOF" {
			color.Red("SYNTACTIC ERROR - Unexpected EOF")
			os.Exit(0)
		}
	}
	panic("should never happen")
}

func errorPrinter(expected string, token Token, line, column int) {

	color.Red("SYNTACTIC ERROR - Expected \"%s\", found \"%s\"\nLine:%d\tColumn:%d", expected, token.Class, line, column-1)

}

func trash(b bool) {}
