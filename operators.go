package golisp

import "strconv"

func plus(list []*lisp) *lisp {
	int1, _ := lispToInt(list[0])
	int2, _ := lispToInt(list[1])

	return NewToken(strconv.Itoa(int1 + int2))
}

func minus(list []*lisp) *lisp {
	int1, _ := lispToInt(list[0])
	int2, _ := lispToInt(list[1])

	return NewToken(strconv.Itoa(int1 - int2))
}

func multiply(list []*lisp) *lisp {
	int1, _ := lispToInt(list[0])
	int2, _ := lispToInt(list[1])

	return NewToken(strconv.Itoa(int1 * int2))
}

func divide(list []*lisp) *lisp {
	int1, _ := lispToInt(list[0])
	int2, _ := lispToInt(list[1])

	return NewToken(strconv.Itoa(int1 / int2))
}

func head(list []*lisp) *lisp {
	arg1 := Evaluate(list[0])
	if len(arg1.list) > 0 {
		return Evaluate(arg1.list[0])
	} else {
		return arg1
	}
}

func lispToInt(a *lisp) (int, error) {
	str := Evaluate(a).String()

	number, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return number, nil
}
