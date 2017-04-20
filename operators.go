package golisp

import (
	"errors"
	"fmt"
	"strconv"
)

type operator func(list []*lisp) (*lisp, error)

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
