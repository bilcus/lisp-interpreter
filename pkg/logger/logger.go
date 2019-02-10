package logger

import "fmt"

var Active bool

func Logf(format string, a ...interface{}) {
	if Active {
		fmt.Printf("gc: "+format+"\n", a...)
	}
}
