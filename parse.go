package main

import (
	"errors"
	"strings"
)

type lisp struct {
	list  []*lisp
	token string
}

func NewToken(token string) *lisp {
	return &lisp{nil, token}
}

func ParseLisp(str string) (*lisp, error) {
	l := &lisp{}
	l.list = make([]*lisp, 0)
	str = strings.TrimSpace(str)

	if !validateLisp(str) {
		return nil, errors.New("unbalanced parens")
	}

	token := ""
	addToken := func() {
		if token != "" {
			l.list = append(l.list, NewToken(token))
			token = ""
		}
	}

	for i := 1; i < len(str); {
		c := str[i : i+1]
		switch c {
		case " ":
			addToken()
			i++
		case ")":
			addToken()
			return l, nil
		default:
			token += c
			i++
		case "(":
			match := strings.Index(str[i:], ")")
			if match == -1 {
				return nil, errors.New("no ')' found")
			}

			nest, err := ParseLisp(str[i : len(str)-1])
			if err != nil {
				return nil, err
			}

			l.list = append(l.list, nest)
			i = i + match
		}
	}
	return l, nil
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

func (a lisp) IsToken() bool {
	ls := a
	for len(ls.list) == 1 {
		ls = *ls.list[0]
	}
	return ls.token != ""
}

func validateLisp(str string) bool {
	stack := ""

	for _, s := range str {
		if s == '(' {
			stack += "("
		} else if s == ')' {
			if strings.HasSuffix(stack, "(") {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return stack == ""
}
