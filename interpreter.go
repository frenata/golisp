package golisp

import (
	"errors"
	"fmt"
)

var operators map[string]operator

func init() {
	operators = map[string]operator{
		"+":    plus,
		"-":    minus,
		"head": head,
		"*":    multiply,
		"/":    divide,
		"inc":  increment,
		"dec":  decrement,
		"map":  mapfunc,
	}
}

type Interpreter struct {
	err error
}

func (i *Interpreter) Err() error {
	return i.err
}

// Evaluate recursively evaluates lisps until it finds single value tokens,
// then executes the operators in the lisps
func (i *Interpreter) Evaluate(a *lisp) *lisp {
	if i.err != nil {
		return nil
	}
	if a.IsToken() || a.IsBlank() {
		return a
	}

	opName := a.list[0].GetToken()
	op := i.getOperator(opName)
	if op == nil { // if no operator defined, return lisp unevaluated
		return a
	}

	return i.apply(op, opName, a.list[1:])
}

func (i *Interpreter) getOperator(op string) operator {
	if i.err != nil {
		return nil
	}
	var function operator
	var ok bool

	if function, ok = operators[op]; !ok {
		return nil
	}
	return function
}

func (i *Interpreter) apply(op operator, name string, list []*lisp) *lisp {
	if i.err != nil {
		return nil
	}
	result, err := op(list)
	if err != nil {
		i.err = errors.New(fmt.Sprintf("operator \"%s\" %s", name, err))
		return nil
	}
	return result
}
