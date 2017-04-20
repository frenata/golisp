package main

import (
	"errors"
	"strings"
)

// A lisp should either be a string or a list of other lisps
// ie, one of these values should be empty
type lisp struct {
	list  []*lisp
	token string
}

// NewToken encapsulates a string into a lisp.
func NewToken(token string) *lisp {
	return &lisp{nil, token}
}

// ParseLisp recursively parses a string, returning the lisp represented.
// If the string does not represent a valid lisp, an error will be returned
func ParseLisp(str string) (*lisp, error) {
	l := &lisp{}
	l.list = make([]*lisp, 0)

	if !validateLisp(str) {
		return nil, errors.New("unbalanced parens")
	}

	str = strings.TrimSpace(str)
	//if strings.HasPrefix(str, "(") && strings.HasSuffix(str, ")") {
	//	str = str[1 : len(str)-1]
	//}
	if strings.HasPrefix(str, "(") && strings.HasSuffix(str, ")") {
		str = str[1 : len(str)-1]
	}

	token := ""
	// add the current token to the list and reset the token
	addToken := func() {
		if token != "" {
			l.list = append(l.list, NewToken(token))
			token = ""
		}
	}

	for i := 0; i < len(str); i++ {
		c := str[i : i+1]
		switch c {
		case ")":
			addToken()
			return l, nil
		case " ":
			addToken()
		default:
			token += c
		case "(":
			// TODO: need to write a function that will find the real closing parens
			close := strings.Index(str[i:], ")")
			// below should never occur:
			if close == -1 {
				panic("no ')' found even though lisp was validated: " + str)
			}

			// remove the last character, which should be the closing ')'
			nest, err := ParseLisp(str[i:])
			if err != nil {
				return nil, err
			}

			l.list = append(l.list, nest)
			i = i + close // skip parsing the nested lisp again
		}
	}
	addToken()
	return l, nil
}

// String prints a lisp according to normal standards
// space separated tokens are surrounded by parens
// Nested lisps take the place of a token and recursively call this function.
func (a lisp) String() string {
	if a.IsToken() {
		return a.GetToken()
	} else {
		str := ""
		for i, s := range a.list {
			space := " "
			if i == len(a.list)-1 {
				space = ""
			}
			str += s.String() + space
		}
		// if lisp is only parens, empty it
		if strings.Trim(str, "()") == "" {
			str = ""
		}
		return "(" + str + ")"
	}
}

// IsToken indicates whether the lisp has a token value
func (a lisp) IsToken() bool {
	return a.GetToken() != ""
}

// GetToken retrieves the token value of a lisp. It ignores excessive
// parens. Thus, (((25))) will return 25.
func (a lisp) GetToken() string {
	ls := a
	for len(ls.list) == 1 {
		ls = *ls.list[0]
	}
	return ls.token
}

// is the lisp just parens?
func (a *lisp) IsBlank() bool {
	return a.String() == "()"
}

// ensures that the parenthesis are properly balanced
func validateLisp(str string) bool {
	stack := ""
	//fmt.Println(str)
	for _, s := range str {
		if s == '(' {
			stack += "("
			//fmt.Println(stack)
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
