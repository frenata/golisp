package golisp

type operator func(list []*lisp) string

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

func Evaluate(a *lisp) string {
	if a.IsToken() {
		return a.GetToken()
	} else if a.IsBlank() {
		return a.String()
	}

	op := a.list[0].GetToken()
	list := a.list[1:]

	var result operator
	var ok bool

	if result, ok = operators[op]; !ok {
		return a.String()
	}

	//fmt.Println(result(list))
	return result(list)
}
