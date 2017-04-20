package main

import "testing"

func TestEvalToken(t *testing.T) {
	a, _ := ParseLisp("24")

	if evaluate(a) != "24" {
		t.Fatal("24 is not 24")
	}
}

func TestAddition(t *testing.T) {
	a, _ := ParseLisp("(+ 1 2)")

	if evaluate(a) != "3" {
		t.Fatal("1 + 2 != 3")
	}
}

func TestRecursiveAddition(t *testing.T) {
	a, err := ParseLisp("(+ 1 (+ 2 (+ 3 4)))")

	if err != nil {
		t.Fatal(err)
	}

	if evaluate(a) != "10" {
		t.Fatal("1 + 2 + 3 + 4 != 10")
	}
}

func TestRecursiveSubtract(t *testing.T) {
	a, _ := ParseLisp("(- 1 (- 2 (- 3 4)))")

	if evaluate(a) != "-2" {
		t.Fatal("(- 1 (- 2 (- 3 4))) != -2")
	}
}

func TestHead(t *testing.T) {
	a, _ := ParseLisp("(head (4 5 6))")

	if evaluate(a) != "4" {
		t.Fatal("head fails")
	}
}
