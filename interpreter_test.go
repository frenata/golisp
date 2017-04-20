package golisp

import "testing"

func TestEvalToken(t *testing.T) {
	a, _ := Parse("24")

	if res, _ := Evaluate(a); res.String() != "24" {
		t.Log(Evaluate(a))
		t.Fatal("24 is not 24")
	}
}

func TestAddition(t *testing.T) {
	a, _ := Parse("(+ 1 2)")

	if res, _ := Evaluate(a); res.String() != "3" {
		t.Fatal("1 + 2 != 3")
	}
}

func TestRecursiveAddition(t *testing.T) {
	a, err := Parse("(+ 1 (+ 2 (+ 3 4)))")

	if err != nil {
		t.Log(a, err)
		t.Fatal(err)
	}

	if res, _ := Evaluate(a); res.String() != "10" {
		t.Fatal("1 + 2 + 3 + 4 != 10")
	}
}

func TestRecursiveSubtract(t *testing.T) {
	a, _ := Parse("(- 1 (- 2 (- 3 4)))")

	if res, _ := Evaluate(a); res.String() != "-2" {
		t.Fatal("(- 1 (- 2 (- 3 4))) != -2")
	}
}

func TestArithmetic(t *testing.T) {
	a, _ := Parse("(/ 56 (+ 1 (* 9 (+ 1 2))))")

	if res, _ := Evaluate(a); res.String() != "2" {
		t.Log(Evaluate(a))
		t.Fatal("(/ 56 (+ 1 (* 9 (+ 1 2)))) != 2")
	}
}

func TestHead(t *testing.T) {
	a, _ := Parse("(head (4 5 6))")

	if res, _ := Evaluate(a); res.String() != "4" {
		t.Fatal("head fails")
	}
}

func TestBadArgs(t *testing.T) {
	badArgsLisps := make([]*lisp, 4)
	badArgsLisps[0], _ = Parse("(+ 1)")
	badArgsLisps[1], _ = Parse("(head 4 5 6)")
	badArgsLisps[2], _ = Parse("(- 4 (* 4))")
	badArgsLisps[3], _ = Parse("(+ for go)")

	for _, bad := range badArgsLisps {
		res, err := Evaluate(bad)
		if err == nil {
			t.Fatalf("%s did not produce an evaluation error, instead: %s", bad, res)
		}
	}
}
