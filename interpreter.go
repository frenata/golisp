package golisp

type operator func(list []*lisp) *lisp

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

func Evaluate(a *lisp) *lisp {
	if a.IsToken() {
		return a
	} else if a.IsBlank() {
		return a
	}

	op := a.list[0].GetToken()
	list := a.list[1:]

	var result operator
	var ok bool

	if result, ok = operators[op]; !ok {
		return a
	}

	//fmt.Println(result(list))
	return result(list)
}
