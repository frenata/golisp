package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/frenata/golisp"
)

// REPL like client
// The user can enter a lisp, it will be parsed into a AST
// Now: the AST will be printed, in lisp style
func main() {
	cli := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("golisp> ")

		input, err := cli.ReadString('\n')
		if err != nil {
			panic(err)
		}

		switch strings.TrimSpace(input) {
		case "q", "quit", "x", "exit":
			fmt.Println("Have a nice day!")
			return
		default:
			parsed, err := golisp.Parse(input)
			if err != nil {
				fmt.Println("parse error:", err)
				continue
			}
			value, err := golisp.Evaluate(parsed)
			if err != nil {
				fmt.Println("evaluate error:", err)
				continue
			}
			fmt.Println(value)
		}
	}
}
