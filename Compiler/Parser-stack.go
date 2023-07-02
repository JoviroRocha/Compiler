package main

type state_Stack []int

func (s *state_Stack) Push(state int) {
	*s = append(*s, state)
}

func (s *state_Stack) Pop(cardinality int) {
	for i := 0; i < cardinality; i++ {
		index := len(*s) - 1
		*s = (*s)[:index]
	}
}

func (s *state_Stack) Get() (state int) {
	index := len(*s) - 1
	return (*s)[index]
}

func (s *state_Stack) Start() {
	s.Push(0)
}
