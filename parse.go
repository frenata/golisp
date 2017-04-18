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

	token := ""
	addToken := func() {
		if token != "" {
			l.list = append(l.list, NewToken(token))
			token = ""
		}
	}

	for i := 1; i < len(str); i++ {
		c := str[i : i+1]
		switch c {
		case " ":
			addToken()
		case ")":
			addToken()
			return l, nil
		default:
			token += c
		case "(":
			match := strings.Index(str[i:], ")")
			nest, err := ParseLisp(str[i : i+match+1])
			if err != nil {
				return nil, err
			}
			l.list = append(l.list, nest)
			i = i + match + 1
		}
	}
	return nil, errors.New("no ')' found")
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
