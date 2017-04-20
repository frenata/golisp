package golisp

import "strconv"

func Evaluate(a *lisp) string {
	if a.IsToken() {
		return a.GetToken()
	} else if a.IsBlank() {
		return a.String()
	}

	switch a.list[0].GetToken() {
	case "+":
		arg1 := Evaluate(a.list[1])
		arg2 := Evaluate(a.list[2])

		int1, _ := strconv.Atoi(arg1)
		int2, _ := strconv.Atoi(arg2)

		return strconv.Itoa(int1 + int2)

	case "-":
		arg1 := Evaluate(a.list[1])
		arg2 := Evaluate(a.list[2])

		int1, _ := strconv.Atoi(arg1)
		int2, _ := strconv.Atoi(arg2)

		return strconv.Itoa(int1 - int2)

	case "head":
		arg1, _ := Parse(Evaluate(a.list[1]))
		if len(arg1.list) > 0 {
			return Evaluate(arg1.list[0])
		} else {
			return arg1.String()
		}

	default:
		return a.String()
	}
}
