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

	default:
		return a.String()
	}
}
