package golisp

import (
	"errors"
	"testing"
)

func TestEvalToken(t *testing.T) {
	interpreter := Interpreter{}
	a, _ := Parse("24")

	if res := interpreter.Evaluate(a); res.String() != "24" {
		t.Log(res)
		t.Fatal("24 is not 24")
	}
}

func TestAddition(t *testing.T) {
	interpreter := Interpreter{}
	a, _ := Parse("(+ 1 2)")

	if res := interpreter.Evaluate(a); res.String() != "3" {
		t.Fatal("1 + 2 != 3")
	}
}

func TestRecursiveAddition(t *testing.T) {
	interpreter := Interpreter{}
	a, err := Parse("(+ 1 (+ 2 (+ 3 4)))")

	if err != nil {
		t.Log(a, err)
		t.Fatal(err)
	}

	if res := interpreter.Evaluate(a); res.String() != "10" {
		t.Fatal("1 + 2 + 3 + 4 != 10")
	}
}

func TestRecursiveSubtract(t *testing.T) {
	interpreter := Interpreter{}
	a, _ := Parse("(- 1 (- 2 (- 3 4)))")

	if res := interpreter.Evaluate(a); res.String() != "-2" {
		t.Fatal("(- 1 (- 2 (- 3 4))) != -2")
	}
}

func TestArithmetic(t *testing.T) {
	interpreter := Interpreter{}
	a, _ := Parse("(/ 56 (+ 1 (* 9 (+ 1 2))))")

	if res := interpreter.Evaluate(a); res.String() != "2" {
		t.Log(res)
		t.Fatal("(/ 56 (+ 1 (* 9 (+ 1 2)))) != 2")
	}
}

func TestHead(t *testing.T) {
	interpreter := Interpreter{}
	a, _ := Parse("(head (4 5 6))")

	if res := interpreter.Evaluate(a); res.String() != "4" {
		t.Fatal("head fails")
	}

	a, _ = Parse("(head ((1 2) (3 4)))")

	if res := interpreter.Evaluate(a); res.String() != "(1 2)" {
		t.Log(a, res)
		t.Fatal("head fails")
	}
}

func TestIncDec(t *testing.T) {
	interpreter := Interpreter{}
	a, _ := Parse("(inc (dec (inc (dec 999))))")

	if res := interpreter.Evaluate(a); res.String() != "999" {
		t.Fatal("(inc (dec (inc (dec 999)))) != 999")
	}
}

func TestBadArgs(t *testing.T) {
	interpreter := Interpreter{}
	badArgsLisps := make([]*lisp, 4)
	badArgsLisps[0], _ = Parse("(+ 1)")
	badArgsLisps[1], _ = Parse("(head 4 5 6)")
	badArgsLisps[2], _ = Parse("(- 4 (* 4))")
	badArgsLisps[3], _ = Parse("(+ for go)")

	for _, bad := range badArgsLisps {
		res := interpreter.Evaluate(bad)
		if bad := interpreter.Err(); bad == nil {
			t.Fatalf("%s did not produce an evaluation error, instead: %s", bad, res)
		}
	}
}

func TestSimpleMap(t *testing.T) {
	interpreter := Interpreter{}
	input, err := Parse("(map inc (1 2 3))")
	t.Log(input, err)
	expected := "(2 3 4)"

	if actual := interpreter.Evaluate(input); actual == nil || actual.String() != expected {
		t.Fatalf("%s did not evalute to %s, instead %s", input, expected, actual)
	}

	input, err = Parse("(map head ((1 2) (3 4)))")
	t.Log(input, err)
	expected = "(1 3)"

	if actual := interpreter.Evaluate(input); actual == nil || actual.String() != expected {
		t.Fatalf("%s did not evalute to %s, instead %s", input, expected, actual)
	}
}

func TestNestedMap(t *testing.T) {
	interpreter := Interpreter{}
	input, err := Parse("(map inc (map head ((99 55 66) (3 4 5) (1 1 1 1 1))))")
	t.Log(input, err)
	expected := "(100 4 2)"

	if actual := interpreter.Evaluate(input); actual == nil || actual.String() != expected {
		t.Fatalf("%s did not evalute to %s, instead %s", input, expected, actual)
	}
}

func TestMapError(t *testing.T) {
	interpreter := Interpreter{}
	input, err := Parse("(map 4 (4 5))")
	t.Log(input, err)

	if _ = interpreter.Evaluate(input); interpreter.Err() == nil {
		t.Fatalf("%s failed to produce an error", input)
	}
}

func TestGetOperatorSuccess(t *testing.T) {
	interpreter := Interpreter{}
	input := "+"
	op := interpreter.getOperator(input)
	if op == nil {
		t.Fatalf("%s failed to be recognized as an operator", input)
	}
}

func TestGetOperatorFailure(t *testing.T) {
	interpreter := Interpreter{}
	input := "o"
	op := interpreter.getOperator(input)
	if op != nil {
		t.Fatalf("%s improperly recognized as an operator", input)
	}
}

func TestInterpreterEarlyReturn(t *testing.T) {
	interpreter := Interpreter{err: errors.New("foo")}
	input := "+"
	op := interpreter.getOperator(input)
	if op != nil {
		t.Fatal("getOperator didn't short-circuit when there was an error already")
	}
}
