package golisp

import "testing"

func TestEvalToken(t *testing.T) {
	a, _ := Parse("24")

	if Evaluate(a) != "24" {
		t.Fatal("24 is not 24")
	}
}

func TestAddition(t *testing.T) {
	a, _ := Parse("(+ 1 2)")

	if Evaluate(a) != "3" {
		t.Fatal("1 + 2 != 3")
	}
}

func TestRecursiveAddition(t *testing.T) {
	a, err := Parse("(+ 1 (+ 2 (+ 3 4)))")

	if err != nil {
		t.Fatal(err)
	}

	if Evaluate(a) != "10" {
		t.Fatal("1 + 2 + 3 + 4 != 10")
	}
}

func TestRecursiveSubtract(t *testing.T) {
	a, _ := Parse("(- 1 (- 2 (- 3 4)))")

	if Evaluate(a) != "-2" {
		t.Fatal("(- 1 (- 2 (- 3 4))) != -2")
	}
}

func TestArithmetic(t *testing.T) {
	a, _ := Parse("(/ 56 (+ 1 (* 9 (+ 1 2))))")

	if Evaluate(a) != "2" {
		t.Log(Evaluate(a))
		t.Fatal("(/ 56 (+ 1 (* 9 (+ 1 2)))) != 2")
	}
}

func TestHead(t *testing.T) {
	a, _ := Parse("(head (4 5 6))")

	if Evaluate(a) != "4" {
		t.Fatal("head fails")
	}
}
