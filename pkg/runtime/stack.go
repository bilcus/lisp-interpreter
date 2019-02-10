package runtime

import (
	"github.com/pkg/errors"
)

type Stack struct {
	stack []Object
}

func NewStack() *Stack {
	return &Stack{
		stack: make([]Object, 0),
	}
}

func (s *Stack) Push(o Object) *Stack {
	s.stack = append(s.stack, o)
	return s
}

func (s *Stack) Pop() Object {
	l := len(s.stack)
	if l <= 0 {
		Error(errors.New("stack underflow"))
	}

	toRet := s.stack[l-1]
	s.stack = s.stack[:l-1]
	return toRet
}

func (s *Stack) PopTimes(n int) Object {
	var o Object
	for i := 0; i < n; i++ {
		o = s.Pop()
	}
	return o
}
