# Basic Lisp interpreter with GC

By Ondrej Bilek for MI-RUN

This is an interpreter for basic subset of LISP. For supported syntax see `testInput`. This interpreter has simulated memory with max object count set to 256. It implements basic mark and sweep GC. GC is triggered when OOM or after hitting the GCThreshold.

## Build
Interpreter is written in Go and you will need latest Go installed.
Do not clone the repository into your `GOPATH`.

```bash
git clone https://github.com/bilcus/lisp-interpreter.git
cd lisp-interpreter
make build
```

## Run
There are two modes:

```
Usage:
	lisp-interpreter <command> [arguments]

The commands are:
	repl	start lisp repl
	input	interpret text file

The arguments are:
	-d	Debug GC print
```

## Test
Feel free to use `testInput` and `testGC` to test the implementation. Enabled debug print to see GC runs.

