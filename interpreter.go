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

// Evaluate recursively evaluates lisps until it finds single value tokens,
// then executes the operators in the lisps
func Evaluate(a *lisp) (*lisp, error) {
	if a.IsToken() || a.IsBlank() {
		return a, nil
	}

	opName := a.list[0].GetToken()
	op := getOperator(opName)
	if op == nil { // if no operator defined, return lisp unevaluated
		return a, nil
	}

	return apply(op, opName, a.list[1:])
}

func getOperator(op string) operator {
	var function operator
	var ok bool

	if function, ok = operators[op]; !ok {
		return nil
	}
	return function
}

func apply(op operator, name string, list []*lisp) (*lisp, error) {
	result, err := op(list)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("operator \"%s\" %s", name, err))
	}
	return result, nil
}
