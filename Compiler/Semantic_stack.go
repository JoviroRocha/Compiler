package main

import (
	"os"

	"github.com/fatih/color"
)

type semantic_Stack []Token

func (s *semantic_Stack) Push(token Token) {
	*s = append(*s, token)
}

func (s *semantic_Stack) Pop() (token Token) {
	index := len(*s) - 1
	token = (*s)[index]
	*s = (*s)[:index]
	return
}

func (s *semantic_Stack) Get() (token Token) {
	index := len(*s) - 1
	return (*s)[index]
}

func (s *semantic_Stack) getStack(value string) (token Token) {
	index := len(*s) - 1
	for index >= 0 {
		if (*s)[index].Class == value {
			return (*s)[index]
		}
		index--
	}
	color.Red("\nERROR: Internal compiler error\n")
	os.Exit(10)
	return
}
