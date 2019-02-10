
build:
	go build ./cmd/lisp-interpreter

run: build
	./lisp-interpreter repl

clean:
	rm lisp-interpreter