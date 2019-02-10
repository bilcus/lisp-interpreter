package logger

import "fmt"

var Active bool

func Log(s string) {
	if Active {
		fmt.Printf("gc: %s\n", s)
	}
}
