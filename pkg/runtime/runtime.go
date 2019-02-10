package runtime

import (
	"errors"
	"fmt"
	"os"
)

const StackMaxObjects = 256

type VM struct {
	memory []Object
	stack  *Stack

	memoryObjects  int
	lastBlockIndex int
}

func NewVM() *VM {
	return &VM{
		memory:         make([]Object, StackMaxObjects),
		stack:          NewStack(),
		lastBlockIndex: -1,
	}
}

func (v *VM) findFreeBlock() (int, error) {
	blockIndex := v.lastBlockIndex

	for {
		blockIndex++
		if blockIndex >= StackMaxObjects {
			blockIndex = 0
		}
		if blockIndex == v.lastBlockIndex {
			return 0, errors.New("out of memory")
		}
		if v.memory[blockIndex] == nil {
			return blockIndex, nil
		}
	}
}

func (v *VM) AllocateObject(o Object) {
	blockIndex, err := v.findFreeBlock()
	if err != nil {
		Error(err)
	}

	v.memory[blockIndex] = o
	v.lastBlockIndex = blockIndex
	v.memoryObjects++
}

func (v *VM) Stack() *Stack {
	return v.stack
}

func Error(err error) {
	fmt.Printf("error: %v\n", err)
	os.Exit(1)
}
