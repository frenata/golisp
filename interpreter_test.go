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
	a, _ := ParseLisp("(+ 1 (+ 2 (+ 3 4)))")

	if evaluate(a) != "10" {
		t.Fatal("1 + 2 + 3 + 4 != 10")
	}
}
