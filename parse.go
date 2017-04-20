package golisp

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

// Parse recursively parses a string, returning the lisp represented.
// If the string does not represent a valid lisp, an error will be returned
func Parse(str string) (*lisp, error) {
	//fmt.Println(str)
	str = strings.TrimSpace(str)
	l := &lisp{}
	l.list = make([]*lisp, 0)

	if !validateLisp(str) {
		return nil, errors.New("unbalanced parens")
	}

	token := ""
	// add the current token to the list and reset the token
	addToken := func() {
		if token != "" {
			l.list = append(l.list, NewToken(token))
			token = ""
		}
	}

	i := 0
	if strings.HasPrefix(str, "(") {
		i = 1
	}
	for ; i < len(str); i++ {
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
			close := findMatchedClose(str[i+1:])
			// below should never occur:
			if close == -1 {
				panic("no ')' found even though lisp was validated: " + str)
			}

			//fmt.Println("nested parse:", str[i:i+close+2])
			nest, err := Parse(str[i : i+close+2]) // Why +2?
			if err != nil {
				return nil, err
			}

			l.list = append(l.list, nest)
			i = i + close + 1 // skip parsing the nested lisp again
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

// finds the matching close parens
func findMatchedClose(str string) int {
	stack := ""
	for i, c := range str {
		if c == '(' {
			stack += "("
		} else if c == ')' && strings.HasSuffix(stack, "(") {
			stack = stack[:len(stack)-1]
		} else if c == ')' {
			return i
		}
	}
	return -1
}
