package stack

// import (
// 	""
// )

type Stack struct{
	empty bool
	// value int
}

func NewStack() *Stack{
	return &Stack{true};
}

func (s *Stack) IsEmpty() bool {
	return s.empty;
}

 func (s *Stack) AddToStack (value int) {
	s.empty = false;
 }