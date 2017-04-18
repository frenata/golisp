package main

import (
	"fmt"
	"testing"
)

func TestParens(t *testing.T) {
	badParens := []string{"(+ 1", ")(", "((((((", "(()", "((()())))"}

	for _, s := range badParens {
		res, err := ParseLisp(s)
		t.Log(res, err)
		if err == nil {
			t.Fatalf("Parsing %s did not produce an error", s)
		}
	}
}

func TestInternalWhitespace(t *testing.T) {
	input := "(   5   6     9      )"
	expected := "(5 6 9)"
	actual, _ := ParseLisp(input)

	if fmt.Sprint(actual) != expected {
		t.Fatal("whitespace is not properly stripped from %s when parsing, %s should be %s", input, actual, expected)
	}
}

func TestExternalWhitespace(t *testing.T) {
	input := "   (   5   6     9      )    "
	expected := "(5 6 9)"
	actual, _ := ParseLisp(input)

	if fmt.Sprint(actual) != expected {
		t.Fatalf("whitespace is not properly stripped from %s when parsing, %s should be %s", input, actual, expected)
	}
}

func TestIsToken(t *testing.T) {
	res, err := ParseLisp("((((((25))))))")

	if err != nil {
		t.Fatalf("error while testing, failed to parse %s: %s", res, err)
	}

	if !res.IsToken() {
		t.Fatalf("%s is not recognized as a token", res)
	}
}

func TestIsNotToken(t *testing.T) {
	res, err := ParseLisp("((((((25))))))")

	if err != nil {
		t.Fatalf("error while testing, failed to parse %s: %s", res, err)
	}

	if !res.IsToken() {
		t.Fatalf("%s is not recognized as a token", res)
	}
}
