package runtime

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type ObjectType int

const (
	_ = iota // ignore first value by assigning to blank identifier
	TypeNil
	TypeVoid
	TypeInteger
	TypeCons
	TypeFunction
	TypeSyntax
	TypeSymbol
	TypeBool
)

type Object interface {
	Allocate(*VM) Object
	Evaluate() Object
	EvaluateFunction(int) Object
	Cdr() Object
	Car() Object
	IntegerValue() int
	StringValue() string
	BoolValue() bool
	String() string
	Type() ObjectType
	Mark()
	UnMark()
	IsMarked() bool
}

// Nil Object
type NilObject struct {
	marked bool
}

func NewNilObject() Object {
	return &NilObject{
		marked: false,
	}
}

func (n *NilObject) Allocate(vm *VM) Object {
	vm.AllocateObject(n)
	return n
}

func (n *NilObject) Evaluate() Object {
	return n
}

func (n *NilObject) EvaluateFunction(args int) Object {
	Error(errors.New("NilObject does not have EvaluateFunction"))
	return n
}

func (n *NilObject) Car() Object {
	Error(errors.New("NilObject does not have Car"))
	return n
}

func (n *NilObject) Cdr() Object {
	Error(errors.New("NilObject does not have Cdr"))
	return n
}

func (n *NilObject) IntegerValue() int {
	Error(errors.New("NilObject does not have IntegerValue"))
	return 0
}

func (n *NilObject) StringValue() string {
	Error(errors.New("NilObject does not have StringValue"))
	return ""
}

func (n *NilObject) BoolValue() bool {
	Error(errors.New("NilObject does not have BoolValue"))
	return false
}

func (n *NilObject) String() string {
	return ""
}

func (n *NilObject) Type() ObjectType {
	return TypeNil
}

func (n *NilObject) Mark() {
	n.marked = true
}

func (n *NilObject) UnMark() {
	n.marked = false
}

func (n *NilObject) IsMarked() bool {
	return n.marked
}

// Void Object
type VoidObject struct {
	marked bool
}

func NewVoidObject() Object {
	return &VoidObject{
		marked: false,
	}
}

func (v *VoidObject) Allocate(vm *VM) Object {
	vm.AllocateObject(v)
	return v
}

func (v *VoidObject) Evaluate() Object {
	return v
}

func (v *VoidObject) EvaluateFunction(args int) Object {
	Error(errors.New("VoidObject does not have EvaluateFunction"))
	return v
}

func (v *VoidObject) Car() Object {
	Error(errors.New("VoidObject does not have Car"))
	return v
}

func (v *VoidObject) Cdr() Object {
	Error(errors.New("VoidObject does not have Cdr"))
	return v
}

func (v *VoidObject) IntegerValue() int {
	Error(errors.New("VoidObject does not have IntegerValue"))
	return 0
}

func (v *VoidObject) StringValue() string {
	Error(errors.New("VoidObject does not have StringValue"))
	return ""
}

func (v *VoidObject) BoolValue() bool {
	Error(errors.New("VoidObject does not have BoolValue"))
	return false
}

func (v *VoidObject) String() string {
	return ""
}

func (v *VoidObject) Type() ObjectType {
	return TypeVoid
}

func (v *VoidObject) Mark() {
	v.marked = true
}

func (v *VoidObject) UnMark() {
	v.marked = false
}

func (v *VoidObject) IsMarked() bool {
	return v.marked
}

// Integer Object
type IntegerObject struct {
	value  int
	marked bool
}

func NewIntegerObject(v int) Object {
	return &IntegerObject{
		value:  v,
		marked: false,
	}
}

func (i *IntegerObject) Allocate(vm *VM) Object {
	vm.AllocateObject(i)
	return i
}

func (i *IntegerObject) Evaluate() Object {
	return i
}

func (i *IntegerObject) EvaluateFunction(args int) Object {
	Error(errors.New("IntegerObject does not have EvaluateFunction"))
	return nil
}

func (i *IntegerObject) Car() Object {
	Error(errors.New("IntegerObject does not have Car"))
	return i
}

func (i *IntegerObject) Cdr() Object {
	Error(errors.New("IntegerObject does not have Cdr"))
	return i
}

func (i *IntegerObject) IntegerValue() int {
	return i.value
}

func (i *IntegerObject) StringValue() string {
	Error(errors.New("IntegerObject does not have StringValue"))
	return ""
}

func (i *IntegerObject) BoolValue() bool {
	Error(errors.New("IntegerObject does not have BoolValue"))
	return false
}

func (i *IntegerObject) String() string {
	return strconv.Itoa(i.value)
}

func (i *IntegerObject) Type() ObjectType {
	return TypeInteger
}

func (i *IntegerObject) Mark() {
	i.marked = true
}

func (i *IntegerObject) UnMark() {
	i.marked = false
}

func (i *IntegerObject) IsMarked() bool {
	return i.marked
}

// Cons Object

type ConsObject struct {
	car    Object
	cdr    Object
	vm     *VM
	marked bool
}

func NewConsObject(s *Stack) Object {
	cdr := s.Pop()
	car := s.Pop()

	return &ConsObject{
		car:    car,
		cdr:    cdr,
		vm:     nil,
		marked: false,
	}
}

func (c *ConsObject) Allocate(vm *VM) Object {
	vm.AllocateObject(c)
	c.vm = vm
	return c
}

func (c *ConsObject) Evaluate() Object {
	function := c.car.Evaluate()

	if function.Type() != TypeFunction && function.Type() != TypeSyntax {
		return c
	}

	numArgs := 0
	args := c.cdr
	for {
		if args.Type() == TypeNil || args.Type() == TypeVoid {
			break
		}

		o := args.Car()
		if function.Type() == TypeFunction {
			o = o.Evaluate()
		}

		c.vm.Stack().Push(o)
		numArgs++

		args = args.Cdr()
	}

	return function.EvaluateFunction(numArgs)
}

func (c *ConsObject) EvaluateFunction(args int) Object {
	Error(errors.New("ConsObject does not have EvaluateFunction"))
	return nil
}

func (c *ConsObject) Car() Object {
	return c.car
}

func (c *ConsObject) Cdr() Object {
	return c.cdr
}

func (c *ConsObject) IntegerValue() int {
	Error(errors.New("ConsObject does not have IntegerValue"))
	return 0
}

func (c *ConsObject) StringValue() string {
	Error(errors.New("ConsObject does not have StringValue"))
	return ""
}

func (c *ConsObject) BoolValue() bool {
	Error(errors.New("ConsObject does not have BoolValue"))
	return false
}

func (c *ConsObject) String() string {
	car := c.car.String()
	car = strings.Replace(car, "(", "", -1)
	car = strings.Replace(car, ")", "", -1)
	car = strings.TrimSpace(car)

	cdr := c.cdr.String()
	cdr = strings.Replace(cdr, "(", "", -1)
	cdr = strings.Replace(cdr, ")", "", -1)
	cdr = strings.TrimSpace(cdr)

	return fmt.Sprintf("(%s %s)", car, cdr)
}

func (c *ConsObject) Type() ObjectType {
	return TypeCons
}

func (c *ConsObject) Mark() {
	c.marked = true
}

func (c *ConsObject) UnMark() {
	c.marked = false
}

func (c *ConsObject) IsMarked() bool {
	return c.marked
}

// Function Object

type FunctionObject struct {
	name   string
	f      func(int, *VM) Object
	vm     *VM
	marked bool
}

func NewFunctionObject(name string, f func(int, *VM) Object) Object {
	return &FunctionObject{
		name:   name,
		f:      f,
		vm:     nil,
		marked: false,
	}
}

func (f *FunctionObject) Allocate(vm *VM) Object {
	vm.AllocateObject(f)
	f.vm = vm
	f.vm.variables[f.name] = f
	return f
}

func (f *FunctionObject) Evaluate() Object {
	Error(errors.New("FunctionObject does not have Evaluate"))
	return nil
}

func (f *FunctionObject) EvaluateFunction(args int) Object {
	return f.f(args, f.vm)
}

func (f *FunctionObject) Car() Object {
	Error(errors.New("FunctionObject does not have Car"))
	return nil
}

func (f *FunctionObject) Cdr() Object {
	Error(errors.New("FunctionObject does not have Cdr"))
	return nil
}

func (f *FunctionObject) IntegerValue() int {
	Error(errors.New("FunctionObject does not have IntegerValue"))
	return 0
}

func (f *FunctionObject) StringValue() string {
	Error(errors.New("FunctionObject does not have StringValue"))
	return ""
}

func (f *FunctionObject) BoolValue() bool {
	Error(errors.New("FunctionObject does not have BoolValue"))
	return false
}

func (f *FunctionObject) String() string {
	return "func"
}

func (f *FunctionObject) Type() ObjectType {
	return TypeFunction
}

func (f *FunctionObject) Mark() {
	f.marked = true
}

func (f *FunctionObject) UnMark() {
	f.marked = false
}

func (f *FunctionObject) IsMarked() bool {
	return f.marked
}

// Syntax Object

type SyntaxObject struct {
	name   string
	f      func(int, *VM) Object
	vm     *VM
	marked bool
}

func NewSyntaxObject(name string, f func(int, *VM) Object) Object {
	return &SyntaxObject{
		name:   name,
		f:      f,
		vm:     nil,
		marked: false,
	}
}

func (s *SyntaxObject) Allocate(vm *VM) Object {
	vm.AllocateObject(s)
	s.vm = vm
	s.vm.variables[s.name] = s
	return s
}

func (s *SyntaxObject) Evaluate() Object {
	Error(errors.New("SyntaxObject does not have Evaluate"))
	return nil
}

func (s *SyntaxObject) EvaluateFunction(args int) Object {
	return s.f(args, s.vm)
}

func (s *SyntaxObject) Car() Object {
	Error(errors.New("SyntaxObject does not have Car"))
	return nil
}

func (s *SyntaxObject) Cdr() Object {
	Error(errors.New("SyntaxObject does not have Cdr"))
	return nil
}

func (s *SyntaxObject) IntegerValue() int {
	Error(errors.New("SyntaxObject does not have IntegerValue"))
	return 0
}

func (s *SyntaxObject) StringValue() string {
	Error(errors.New("SyntaxObject does not have StringValue"))
	return ""
}

func (s *SyntaxObject) BoolValue() bool {
	Error(errors.New("SyntaxObject does not have BoolValue"))
	return false
}

func (s *SyntaxObject) String() string {
	return "syntax"
}

func (s *SyntaxObject) Type() ObjectType {
	return TypeSyntax
}

func (s *SyntaxObject) Mark() {
	s.marked = true
}

func (s *SyntaxObject) UnMark() {
	s.marked = false
}

func (s *SyntaxObject) IsMarked() bool {
	return s.marked
}

// Symbol Object

type SymbolObject struct {
	name   string
	vm     *VM
	marked bool
}

func NewSymbolObject(name string) Object {
	return &SymbolObject{
		name:   name,
		vm:     nil,
		marked: false,
	}
}

func (s *SymbolObject) Allocate(vm *VM) Object {
	vm.AllocateObject(s)
	s.vm = vm
	return s
}

func (s *SymbolObject) Evaluate() Object {
	o, found := s.vm.variables[s.name]
	if !found {
		Error(errors.Errorf("variable %s not found", s.name))
		return NewVoidObject().Allocate(s.vm)
	}
	return o
}

func (s *SymbolObject) EvaluateFunction(args int) Object {
	Error(errors.New("SymbolObject does not have EvaluateFunction"))
	return nil
}

func (s *SymbolObject) Car() Object {
	Error(errors.New("SymbolObject does not have Car"))
	return nil
}

func (s *SymbolObject) Cdr() Object {
	Error(errors.New("SymbolObject does not have Cdr"))
	return nil
}

func (s *SymbolObject) IntegerValue() int {
	Error(errors.New("SymbolObject does not have IntegerValue"))
	return 0
}

func (s *SymbolObject) StringValue() string {
	return s.name
}

func (s *SymbolObject) BoolValue() bool {
	Error(errors.New("SymbolObject does not have BoolValue"))
	return false
}

func (s *SymbolObject) String() string {
	return "symbol"
}

func (s *SymbolObject) Type() ObjectType {
	return TypeSymbol
}

func (s *SymbolObject) Mark() {
	s.marked = true
}

func (s *SymbolObject) UnMark() {
	s.marked = false
}

func (s *SymbolObject) IsMarked() bool {
	return s.marked
}

// Bool Object

type BoolObject struct {
	value  bool
	marked bool
}

func NewBoolObject(value bool) Object {
	return &BoolObject{
		value:  value,
		marked: false,
	}
}

func (b *BoolObject) Allocate(vm *VM) Object {
	vm.AllocateObject(b)
	return b
}

func (b *BoolObject) Evaluate() Object {
	return b
}

func (b *BoolObject) EvaluateFunction(args int) Object {
	Error(errors.New("BoolObject does not have EvaluateFunction"))
	return nil
}

func (b *BoolObject) Car() Object {
	Error(errors.New("BoolObject does not have Car"))
	return nil
}

func (b *BoolObject) Cdr() Object {
	Error(errors.New("BoolObject does not have Cdr"))
	return nil
}

func (b *BoolObject) IntegerValue() int {
	Error(errors.New("BoolObject does not have IntegerValue"))
	return 0
}

func (b *BoolObject) StringValue() string {
	Error(errors.New("BoolObject does not have StringValue"))
	return ""
}

func (b *BoolObject) BoolValue() bool {
	return b.value
}

func (b *BoolObject) String() string {
	if b.value {
		return "T"
	}
	return "F"
}

func (b *BoolObject) Type() ObjectType {
	return TypeBool
}

func (b *BoolObject) Mark() {
	b.marked = true
}

func (b *BoolObject) UnMark() {
	b.marked = false
}

func (b *BoolObject) IsMarked() bool {
	return b.marked
}
