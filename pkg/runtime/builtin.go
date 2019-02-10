package runtime

import (
	"github.com/pkg/errors"
)

func builtinPlus(numArgs int, vm *VM) Object {
	sum := 0
	for i := 0; i < numArgs; i++ {
		sum += vm.Stack().Pop().IntegerValue()
	}

	return NewIntegerObject(sum).Allocate(vm)
}

func builtinMinus(numArgs int, vm *VM) Object {
	if numArgs == 0 {
		Error(errors.New("no arguments to minus operator"))
		return NewVoidObject().Allocate(vm)
	}

	if numArgs == 1 {
		result := vm.Stack().Pop().IntegerValue()
		return NewIntegerObject(-result).Allocate(vm)
	}

	sum := 0
	for i := 0; i < numArgs-1; i++ {
		sum += vm.Stack().Pop().IntegerValue()
	}
	last := vm.Stack().Pop().IntegerValue()

	return NewIntegerObject(last - sum).Allocate(vm)
}

func builtinTimes(numArgs int, vm *VM) Object {
	sum := 1
	for i := 0; i < numArgs; i++ {
		sum *= vm.Stack().Pop().IntegerValue()
	}

	return NewIntegerObject(sum).Allocate(vm)
}

func builtinEquals(numArgs int, vm *VM) Object {
	if numArgs != 2 {
		Error(errors.New("equals operator expects 2 arguments"))
		vm.Stack().PopTimes(numArgs)
		return NewVoidObject().Allocate(vm)
	}

	val1 := vm.Stack().Pop().IntegerValue()
	val2 := vm.Stack().Pop().IntegerValue()

	return NewBoolObject(val1 == val2).Allocate(vm)
}

func builtinLessThan(numArgs int, vm *VM) Object {
	if numArgs != 2 {
		Error(errors.New("less than operator expects 2 arguments"))
		vm.Stack().PopTimes(numArgs)
		return NewVoidObject().Allocate(vm)
	}

	val2 := vm.Stack().Pop().IntegerValue()
	val1 := vm.Stack().Pop().IntegerValue()

	return NewBoolObject(val1 < val2).Allocate(vm)
}

func builtinGreaterThan(numArgs int, vm *VM) Object {
	if numArgs != 2 {
		Error(errors.New("greater than operator expects 2 arguments"))
		vm.Stack().PopTimes(numArgs)
		return NewVoidObject().Allocate(vm)
	}

	val2 := vm.Stack().Pop().IntegerValue()
	val1 := vm.Stack().Pop().IntegerValue()

	return NewBoolObject(val1 > val2).Allocate(vm)
}

func builtinCar(numArgs int, vm *VM) Object {
	if numArgs != 1 {
		Error(errors.New("car operator expects 1 argument"))
		vm.Stack().PopTimes(numArgs)
		return NewVoidObject().Allocate(vm)
	}

	return vm.Stack().Pop().Car()
}

func builtinCdr(numArgs int, vm *VM) Object {
	if numArgs != 1 {
		Error(errors.New("cdr operator expects 1 argument"))
		vm.Stack().PopTimes(numArgs)
		return NewVoidObject().Allocate(vm)
	}

	return vm.Stack().Pop().Cdr()
}

func builtinIf(numArgs int, vm *VM) Object {
	if numArgs != 3 {
		Error(errors.New("if operator expects 3 argument"))
		vm.Stack().PopTimes(numArgs)
		return NewVoidObject().Allocate(vm)
	}

	falseExpr := vm.Stack().Pop()
	trueExpr := vm.Stack().Pop()
	condExpr := vm.Stack().Pop()

	if condExpr.Evaluate().BoolValue() {
		return trueExpr.Evaluate()
	}
	return falseExpr.Evaluate()
}

func builtinDefine(numArgs int, vm *VM) Object {
	if numArgs != 2 {
		Error(errors.New("define operator expects 2 argument"))
		vm.Stack().PopTimes(numArgs)
		return NewVoidObject().Allocate(vm)
	}

	expr := vm.Stack().Pop()
	varName := vm.Stack().Pop().StringValue()

	vm.variables[varName] = expr

	return NewVoidObject().Allocate(vm)
}
