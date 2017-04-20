package main

import (
	"fmt"
	"testing"
)

func TestParens(t *testing.T) {
	badParens := []string{"(+ 1", ")(", "((((((", "(()", "((()())))"}
	goodParens := []string{"(1 (2 3) (4 5))", "(head ()(()()))"}

	for _, s := range badParens {
		_, err := ParseLisp(s)
		if err == nil {
			t.Fatalf("Parsing %s did not produce an error", s)
		}
	}

	for _, s := range goodParens {
		_, err := ParseLisp(s)
		if err != nil {
			t.Fatalf("Parsing %s produced an error", s)
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
	res, err := ParseLisp("(1 2 3 4)")

	if err != nil {
		t.Fatalf("error while testing, failed to parse %s: %s", res, err)
	}

	if res.IsToken() {
		t.Fatalf("%s is incorrectly recognized as a token", res)
	}
}

func TestEmptyLisp(t *testing.T) {
	res, err := ParseLisp("((((((()))))))")

	if err != nil {
		t.Fatalf("error while testing, failed to parse %s: %s", res, err)
	}

	if res.String() != "()" {
		t.Fatalf("%s is not reduced to \"()\"", res)
	}

	res, err = ParseLisp("")

	if err != nil {
		t.Fatalf("error while testing, failed to parse %s: %s", res, err)
	}

	if res.String() != "()" {
		t.Fatalf("%s is not reduced to \"()\"", res)
	}
}
