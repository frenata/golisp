package golisp

import (
	"errors"
	"fmt"
)

type operator func(list []*lisp) (*lisp, error)

var operators map[string]operator

func init() {
	operators = map[string]operator{
		"+":    plus,
		"-":    minus,
		"head": head,
		"*":    multiply,
		"/":    divide,
	}
}

func Evaluate(a *lisp) (*lisp, error) {
	if a.IsToken() || a.IsBlank() {
		return a, nil
	}

	op := a.list[0].GetToken()
	list := a.list[1:]

	var function operator
	var ok bool

	if function, ok = operators[op]; !ok {
		return a, nil
	}

	result, err := function(list)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("operator \"%s\" %s", op, err))
	}
	return result, nil

}
