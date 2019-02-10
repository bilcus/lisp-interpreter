package runtime

import (
	"fmt"
	"strconv"
)

// Nil Object
type NilObject struct{}

func NewNilObject() Object {
	return &NilObject{}
}

func (n *NilObject) Allocate(vm *VM) Object {
	vm.AllocateObject(n)
	return n
}

func (n *NilObject) Evaluate() Object {
	return n
}

func (n *NilObject) String() string {
	return "nil"
}

// Integer Object
type IntegerObject struct {
	value int
}

func NewIntegerObject(v int) Object {
	return &IntegerObject{
		value: v,
	}
}

func (i *IntegerObject) Allocate(vm *VM) Object {
	vm.AllocateObject(i)
	return i
}

func (i *IntegerObject) Evaluate() Object {
	return i
}

func (i *IntegerObject) String() string {
	return strconv.Itoa(i.value)
}

// Cons Object

type ConsObject struct {
	car Object
	cdr Object
}

func NewConsObject(s *Stack) Object {
	cdr := s.Pop()
	car := s.Pop()

	return &ConsObject{
		car: car,
		cdr: cdr,
	}
}

func (c *ConsObject) Allocate(vm *VM) Object {
	vm.AllocateObject(c)
	return c
}

func (c *ConsObject) Evaluate() Object {
	return c
}

func (c *ConsObject) String() string {
	return fmt.Sprintf("( %v %v )", c.car, c.cdr)
}
