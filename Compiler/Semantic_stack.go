package main

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

func (s *semantic_Stack) getStack(value string, jump ...bool) (token Token) {
	index := len(*s) - 2
	for index >= 0 {
		if (*s)[index].Class == value {
			if len(jump) == 0 {
				return (*s)[index]
			}
			jump = nil
		}
		index--
	}
	if (*s)[len(*s)-1].Class == value {
		return (*s)[len(*s)-1]
	}

	return Token{"", "", ""}
}

func (s *semantic_Stack) updateStack(oldValue string, token Token) {
	index := len(*s) - 2
	for index >= 0 {
		if (*s)[index].Class == oldValue {
			(*s)[index] = token
			return
		}
		index--
	}
}
