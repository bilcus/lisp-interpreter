package repl

import (
	"fmt"
	"io"
	"os"

	"lisp-interpreter/pkg/parser"
	"lisp-interpreter/pkg/runtime"
)

func start(r io.Reader) {
	vm := runtime.NewVM()
	p := parser.NewParser(vm, r)

	fmt.Println("Basic Lisp interpreter with Mark and Sweep GC by Ondrej Bilek")
	fmt.Println()

	for {
		if p.IsEOF() {
			return
		}

		fmt.Print("> ")

		object := p.Parse()
		if object == nil {
			fmt.Println("nil input")
			return
		}

		object = object.Evaluate()
		if object == nil {
			fmt.Println("nil input")
			return
		}

		fmt.Println(object)
	}
}

func StartWithFile(name string) {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	start(file)
}

func StartWithStdin() {
	start(os.Stdin)
}
