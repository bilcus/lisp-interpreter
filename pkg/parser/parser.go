package parser

import (
	"bufio"
	"io"
	"math"
	"strconv"
	"unicode"

	"lisp-interpreter/pkg/runtime"

	"github.com/pkg/errors"
)

const EOF = math.MaxInt32

type Parser struct {
	vm      *runtime.VM
	scanner io.RuneScanner
	isEOF   bool

	currentRune rune
}

func (p *Parser) readRune() (rune, bool) {
	whitespace := false

	for {
		r, _, err := p.scanner.ReadRune()
		if err != nil {
			if err == io.EOF {
				p.isEOF = true
				return EOF, true
			}
			runtime.Error(errors.Wrap(err, "reading rune"))
		}

		if unicode.IsSpace(r) {
			whitespace = true
		} else {
			p.currentRune = r
			return r, whitespace
		}
	}
}

func (p *Parser) unreadRune() {
	if err := p.scanner.UnreadRune(); err != nil {
		runtime.Error(err)
	}
}

func NewParser(vm *runtime.VM, r io.Reader) *Parser {
	return &Parser{
		vm:      vm,
		scanner: bufio.NewReader(r),
		isEOF:   false,
	}
}

func (p *Parser) Parse() runtime.Object {
	return p.parse()
}

func (p *Parser) IsEOF() bool {
	return p.isEOF
}

func (p *Parser) parse() runtime.Object {
	ch, _ := p.readRune()
	if ch == EOF {
		return runtime.NewNilObject().Allocate(p.vm)
	}

	switch ch {
	case '(':
		return p.parseList()
	case ')':
		parseError(errors.New("unexpected ')'"))
	default:
		if unicode.IsDigit(ch) {
			val := p.parseInteger()
			return runtime.NewIntegerObject(val).Allocate(p.vm)
		} else {
			val := p.parseSymbol()
			return runtime.NewSymbolObject(val).Allocate(p.vm)
		}
	}

	return runtime.NewNilObject().Allocate(p.vm)
}

func (p *Parser) parseList() runtime.Object {
	ch, _ := p.readRune()
	if ch == ')' {
		return runtime.NewNilObject().Allocate(p.vm)
	}

	p.unreadRune()

	p.vm.Stack().Push(p.parse())
	p.vm.Stack().Push(p.parseList())

	return runtime.NewConsObject(p.vm.Stack()).Allocate(p.vm)
}

func (p *Parser) parseInteger() int {
	buffer := string(p.currentRune)

	for {
		r, ws := p.readRune()

		if ws == true || !unicode.IsDigit(r) {
			p.unreadRune()

			val, err := strconv.ParseInt(buffer, 10, 64)
			if err != nil {
				parseError(err)
			}

			return int(val)
		}

		buffer += string(r)
	}
}

func (p *Parser) parseSymbol() string {
	buffer := string(p.currentRune)

	for {
		r, ws := p.readRune()

		if ws == true || unicode.IsDigit(r) || r == '(' || r == ')' {
			p.unreadRune()
			return buffer
		}

		buffer += string(r)
	}
}

func parseError(err error) {
	runtime.Error(errors.Wrap(err, "parser"))
}
