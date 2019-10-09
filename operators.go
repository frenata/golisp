package golisp

import (
	"errors"
	"fmt"
	"strconv"
)

type operator func(list []*lisp) (*lisp, error)

// binary operators
func mapfunc(list []*lisp) (*lisp, error) {
	interpreter := Interpreter{}
	if len(list) != 2 {
		return nil, errors.New("expects 2 arguments")
	}

	opName := list[0].GetToken()
	op := interpreter.getOperator(opName)
	if op == nil {
		return nil, errors.New("expects a valid operator as the 1st argument")
	}

	oldlist := interpreter.Evaluate(list[1])
	if oldlist.IsToken() { // allows for something like (map inc 4)
		return interpreter.apply(op, opName, []*lisp{oldlist}), interpreter.err
	}

	newlist := ""
	for _, a := range oldlist.list {
		result := interpreter.apply(op, opName, []*lisp{a})
		newlist += result.String() + " "
	}

	parsed, err := Parse(newlist)
	if interpreter.err != nil {
		err = interpreter.err
	}
	return parsed, err
}

func plus(list []*lisp) (*lisp, error) {
	if len(list) != 2 {
		return nil, errors.New("expects 2 arguments")
	}
	int1, err := lispToInt(list[0])
	if err != nil {
		return nil, err
	}
	int2, err := lispToInt(list[1])
	if err != nil {
		return nil, err
	}

	return NewToken(strconv.Itoa(int1 + int2)), nil
}

func minus(list []*lisp) (*lisp, error) {
	if len(list) != 2 {
		return nil, errors.New("expects 2 arguments")
	}
	int1, err := lispToInt(list[0])
	if err != nil {
		return nil, err
	}
	int2, err := lispToInt(list[1])
	if err != nil {
		return nil, err
	}

	return NewToken(strconv.Itoa(int1 - int2)), nil
}

func multiply(list []*lisp) (*lisp, error) {
	if len(list) != 2 {
		return nil, errors.New("expects 2 arguments")
	}
	int1, err := lispToInt(list[0])
	if err != nil {
		return nil, err
	}
	int2, err := lispToInt(list[1])
	if err != nil {
		return nil, err
	}

	return NewToken(strconv.Itoa(int1 * int2)), nil
}

func divide(list []*lisp) (*lisp, error) {
	if len(list) != 2 {
		return nil, errors.New("expects 2 arguments")
	}
	int1, err := lispToInt(list[0])
	if err != nil {
		return nil, err
	}
	int2, err := lispToInt(list[1])
	if err != nil {
		return nil, err
	}

	return NewToken(strconv.Itoa(int1 / int2)), nil
}

// unary operators
func head(list []*lisp) (*lisp, error) {
	interpreter := Interpreter{}
	if len(list) != 1 {
		return nil, errors.New("expects 1 argument")
	}
	arg1 := interpreter.Evaluate(list[0])
	if len(arg1.list) > 0 {
		return interpreter.Evaluate(arg1.list[0]), interpreter.err
	} else {
		return arg1, nil
	}
}

func increment(list []*lisp) (*lisp, error) {
	if len(list) != 1 {
		return nil, errors.New("expects 1 argument")
	}
	int1, err := lispToInt(list[0])
	if err != nil {
		return nil, err
	}
	return NewToken(strconv.Itoa(int1 + 1)), nil
}

func decrement(list []*lisp) (*lisp, error) {
	if len(list) != 1 {
		return nil, errors.New("expects 1 argument")
	}
	int1, err := lispToInt(list[0])
	if err != nil {
		return nil, err
	}
	return NewToken(strconv.Itoa(int1 - 1)), nil
}

func lispToInt(a *lisp) (int, error) {
	interpreter := Interpreter{}
	list := interpreter.Evaluate(a)
	if interpreter.err != nil {
		return 0, interpreter.err
	}

	number, err := strconv.Atoi(list.String())
	if err != nil {
		return 0, errors.New(fmt.Sprintf("expects numeric arguments, was given argument \"%s\"", a))
	}

	return number, nil
}
