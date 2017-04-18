package main

import "strings"

type lisp struct {
	list  []lisp
	token string
}

func NewToken(token string) lisp {
	return lisp{nil, token}
}

func ParseLisp(str string) lisp {
	l := lisp{}

	l.list = make([]lisp, 0)

	token := ""
	for i := 1; i < len(str); i++ {
		c := str[i : i+1]
		switch c {
		case " ":
			l.list = append(l.list, NewToken(token))
			token = ""
		case "(":
			match := strings.Index(str[i:], ")")
			l.list = append(l.list, ParseLisp(str[i:i+match+1]))
			i = i + match + 1
		case ")":
			l.list = append(l.list, NewToken(token))
			return l
		default:
			token += c
		}
	}

	return l
}

func (a lisp) String() string {
	if a.list == nil {
		return a.token
	} else {
		str := ""
		for i, s := range a.list {
			space := " "
			if i == len(a.list)-1 {
				space = ""
			}
			str += s.String() + space
		}

		if len(a.list) > 1 {
			str = "(" + str + ")"
		}
		return str
	}
}
