package main

import "strconv"

func evaluate(a *lisp) string {
	if a.IsToken() {
		return a.GetToken()
	}

	switch a.list[0].GetToken() {
	case "+":
		arg1 := evaluate(a.list[1])
		arg2 := evaluate(a.list[2])

		int1, _ := strconv.Atoi(arg1)
		int2, _ := strconv.Atoi(arg2)

		return strconv.Itoa(int1 + int2)

	case "-":
		arg1 := evaluate(a.list[1])
		arg2 := evaluate(a.list[2])

		int1, _ := strconv.Atoi(arg1)
		int2, _ := strconv.Atoi(arg2)

		return strconv.Itoa(int1 - int2)

	case "head":
		arg1, _ := ParseLisp(evaluate(a.list[1]))
		return evaluate(arg1.list[0])

	default:
		return a.String()
	}
}
