package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func Parser(filePtr *os.File, line int, column int) {

	state := 0
	token := SCANNER(filePtr, &line, &column)

	for true {

		if token.Class == "ERROR" {
			token = SCANNER(filePtr, &line, &column)
			continue
		}

		state = stateStack.Get()

		action, exists := SLR[state][token.Class]

		if exists == false {
			action = errorHandler(state, line, column, token)
		}

		if newState, err := strconv.Atoi(action); err != nil {
			if action == "ACC" && token.Class == "EOF" {
				fmt.Println("P' -> P")
				return
			}
			Reduce(action)
			continue

		} else {
			stateStack.Push(newState)
			token = SCANNER(filePtr, &line, &column)
		}
	}

}

func Reduce(action string) {

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

	trash(exists)
}

func errorHandler(state, line, column int, token Token) (action string) {

	if state == 0 {
		errorPrinter("inicio", token, line, column)
		return errorCorrector(state, line, column, token, "inicio")
	} else if state == 1 {
		errorPrinter("EOF", token, line, column)
		return errorCorrector(state, line, column, token, "EOF")
	} else if state == 2 {
		errorPrinter("varinicio", token, line, column)
		return errorCorrector(state, line, column, token, "varinicio")
	} else if state == 4 || state == 19 {
		errorPrinter("varfim\" or \"inteiro\" or \"real\" or \"literal", token, line, column)
	} else if state == 11 || state == 21 || state == 67 {
		errorPrinter("ID", token, line, column)
		return errorCorrector(state, line, column, token, "ID")
	} else if state == 12 {
		errorPrinter("LIT\" or \"NUM\" or \"ID", token, line, column)
	} else if state == 13 {
		errorPrinter("RCB", token, line, column)
		return errorCorrector(state, line, column, token, "RCB")
	} else if state == 16 || state == 17 {
		errorPrinter("AB_P", token, line, column)
		return errorCorrector(state, line, column, token, "AB_P")
	} else if state == 20 || state == 29 || state == 30 || state == 49 || state == 53 {
		errorPrinter("PT_V", token, line, column)
		return errorCorrector(state, line, column, token, "PT_V")
	} else if state == 34 || state == 45 || state == 46 || state == 69 || state == 71 {
		errorPrinter("ID\" or \"NUM", token, line, column)
	} else if state == 54 {
		errorPrinter("OPM", token, line, column)
		return errorCorrector(state, line, column, token, "OPM")
	} else if state == 63 || state == 65 {
		errorPrinter("FC_P", token, line, column)
		return errorCorrector(state, line, column, token, "FC_P")
	} else if state == 64 {
		errorPrinter("OPR", token, line, column)
		return errorCorrector(state, line, column, token, "OPR")
	} else if state == 70 {
		errorPrinter("entao", token, line, column)
		return errorCorrector(state, line, column, token, "entao")
	} else if state == 3 || state == 6 || state == 7 || state == 8 || state == 9 {
		errorPrinter("fim\" or \"leia\" or \"escreva\" or \"ID\" or \"se\" or \"repita", token, line, column)
	} else if state == 14 || state == 36 || state == 37 || state == 38 {
		errorPrinter("fimse\" or \"leia\" or \"escreva\" or \"ID\" or \"se", token, line, column)
	} else if state == 15 || state == 41 || state == 42 || state == 43 {
		errorPrinter("fimrepita\" or \"leia\" or \"escreva\" or \"ID\" or \"se", token, line, column)
	}

	panic("should never happen")
}

func errorCorrector(state, line, column int, oldToken Token, fixToken string) (newAction string) {

	action, _ := SLR[state][fixToken]

	if newState, err := strconv.Atoi(action); err != nil {
		Reduce(action)
	} else {
		stateStack.Push(newState)
	}

	newState := stateStack.Get()
	newAction, exist := SLR[newState][oldToken.Class]

	if exist == false {
		newAction = errorHandler(newState, line, column, oldToken)
	}

	return
}

func errorPanic() {

}
func errorPrinter(expected string, token Token, line, column int) {

	color.Red("SYNTACTIC ERROR - Expected \"%s\", found \"%s\"\nLine:%d\tColumn:%d", expected, token.Class, line, column-1)

}

func trash(b bool) {}
