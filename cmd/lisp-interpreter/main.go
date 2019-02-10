package main

import (
	"fmt"
	"os"

	"lisp-interpreter/pkg/logger"
	"lisp-interpreter/pkg/repl"
)

const usage = `Basic Lisp interpreter with Mark and Sweep GC by Ondrej Bilek

Usage:
	lisp-interpreter <command> [arguments]

The commands are:
	repl	start lisp repl
	input	interpret text file

The arguments are:
	-d		Debug GC print

`

const usageInput = `Interpret text file

Usage:
	lisp-interpreter input [path] [arguments]

The arguments are:
	-d		Debug GC print

`

func main() {
	if len(os.Args) == 1 {
		fmt.Print(usage)
		return
	}

	if os.Args[2] == "-d" || len(os.Args) > 4 && os.Args[3] == "-d" {
		logger.Active = true
		fmt.Println("GC logging active")
	}

	switch os.Args[1] {
	case "repl":
		repl.StartWithStdin()
		return
	case "input":
		if len(os.Args) < 3 {
			fmt.Print(usageInput)
			return
		}

		repl.StartWithFile(os.Args[2])
		return
	default:
		fmt.Print(usage)
		return
	}
}
