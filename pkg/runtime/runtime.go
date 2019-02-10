package runtime

import (
	"fmt"
	"os"

	"lisp-interpreter/pkg/logger"
)

const StackMaxObjects = 256

type VM struct {
	memory    []Object
	variables map[string]Object

	stack *Stack

	memoryObjects  int
	lastBlockIndex int
	gcThreshold    int
}

func NewVM() *VM {
	vm := &VM{
		memory:         make([]Object, StackMaxObjects),
		variables:      make(map[string]Object),
		stack:          NewStack(),
		memoryObjects:  0,
		lastBlockIndex: -1,
		gcThreshold:    10,
	}

	NewFunctionObject("+", builtinPlus).Allocate(vm)
	NewFunctionObject("-", builtinMinus).Allocate(vm)
	NewFunctionObject("*", builtinTimes).Allocate(vm)
	NewFunctionObject("=", builtinEquals).Allocate(vm)
	NewFunctionObject("<", builtinLessThan).Allocate(vm)
	NewFunctionObject(">", builtinGreaterThan).Allocate(vm)
	NewFunctionObject("car", builtinCar).Allocate(vm)
	NewFunctionObject("cdr", builtinCdr).Allocate(vm)
	NewSyntaxObject("if", builtinIf).Allocate(vm)
	NewSyntaxObject("define", builtinDefine).Allocate(vm)

	return vm
}

func (v *VM) AllocateObject(o Object) {
	if v.memoryObjects >= v.gcThreshold || v.memoryObjects == StackMaxObjects {
		v.gc()
	}

	blockIndex := v.findFreeBlock()

	v.memory[blockIndex] = o
	v.lastBlockIndex = blockIndex
	v.memoryObjects++
}

func (v *VM) FreeObject(blockIndex int) {
	o := v.memory[blockIndex]
	if o == nil {
		fmt.Println("error: double free")
		os.Exit(1)
	}

	for key, varO := range v.variables {
		if o == varO {
			delete(v.variables, key)
		}
	}

	v.memory[blockIndex] = nil
	v.memoryObjects--
}

func (v *VM) Stack() *Stack {
	return v.stack
}

func (v *VM) findFreeBlock() int {
	blockIndex := v.lastBlockIndex

	for {
		blockIndex++
		if blockIndex >= StackMaxObjects {
			blockIndex = 0
		}
		if blockIndex == v.lastBlockIndex {
			fmt.Println("error: out of memory")
			os.Exit(1)
		}
		if v.memory[blockIndex] == nil {
			return blockIndex
		}
	}
}

func (v *VM) gc() {
	before := v.memoryObjects

	v.mark()
	v.sweep()

	v.gcThreshold = v.memoryObjects * 2

	logger.Logf("# of objects %d->%d", before, v.memoryObjects)
}

func (v *VM) mark() {
	for _, v := range v.variables {
		v.Mark()
	}

	for _, o := range v.stack.stack {
		o.Mark()
	}
}

func (v *VM) sweep() {
	for i, o := range v.memory {
		if o == nil {
			continue
		}

		if o.IsMarked() {
			o.UnMark()
		} else {
			v.FreeObject(i)
		}
	}
}

func Error(err error) {
	fmt.Printf("error: %v ", err)
}
