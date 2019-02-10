package runtime

type Object interface {
	Allocate(vm *VM) Object
	Evaluate() Object
	String() string
}
