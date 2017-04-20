package golisp

import (
	"errors"
	"fmt"
	"strconv"
)

type operator func(list []*lisp) (*lisp, error)

// binary operators
func mapfunc(list []*lisp) (*lisp, error) {
	if len(list) != 2 {
		return nil, errors.New("expects 2 arguments")
	}

	opName := list[0].GetToken()
	op := getOperator(opName)
	if op == nil {
		return nil, errors.New("expects a valid operator as the 1st argument")
	}

	oldlist, err := Evaluate(list[1])
	if err != nil {
		return nil, err
	}
	if oldlist.IsToken() { // allows for something like (map inc 4)
		return apply(op, opName, []*lisp{oldlist})
	}

	newlist := ""
	for _, a := range oldlist.list {
		result, err := apply(op, opName, []*lisp{a})
		if err != nil {
			return nil, err
		}
		newlist += result.String() + " "
	}

	return Parse(newlist)
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
	if len(list) != 1 {
		return nil, errors.New("expects 1 argument")
	}
	arg1, err := Evaluate(list[0])
	if err != nil {
		return nil, err
	}
	if len(arg1.list) > 0 {
		return Evaluate(arg1.list[0])
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
	list, err := Evaluate(a)
	if err != nil {
		return 0, err
	}

	number, err := strconv.Atoi(list.String())
	if err != nil {
		return 0, errors.New(fmt.Sprintf("expects numeric arguments, was given argument \"%s\"", a))
	}

	return number, nil
}
