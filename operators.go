package golisp

import "strconv"

func plus(list []*lisp) string {
	//arg1 := Evaluate(list[0])
	arg2 := Evaluate(list[1])

	//int1, _ := strconv.Atoi(arg1)
	int2, _ := strconv.Atoi(arg2)

	int1, err := lispToInt(list[0])
	_ = err

	return strconv.Itoa(int1 + int2)
}

func minus(list []*lisp) string {
	arg1 := Evaluate(list[0])
	arg2 := Evaluate(list[1])

	int1, _ := strconv.Atoi(arg1)
	int2, _ := strconv.Atoi(arg2)

	return strconv.Itoa(int1 - int2)
}

func multiply(list []*lisp) string {
	arg1 := Evaluate(list[0])
	arg2 := Evaluate(list[1])

	int1, _ := strconv.Atoi(arg1)
	int2, _ := strconv.Atoi(arg2)

	return strconv.Itoa(int1 * int2)
}

func divide(list []*lisp) string {
	arg1 := Evaluate(list[0])
	arg2 := Evaluate(list[1])

	int1, _ := strconv.Atoi(arg1)
	int2, _ := strconv.Atoi(arg2)

	return strconv.Itoa(int1 / int2)
}

func head(list []*lisp) string {
	arg1, _ := Parse(Evaluate(list[0]))
	if len(arg1.list) > 0 {
		return Evaluate(arg1.list[0])
	} else {
		return arg1.String()
	}
}

func lispToInt(a *lisp) (int, error) {
	str := Evaluate(a)

	number, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return number, nil
}
