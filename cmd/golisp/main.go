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
// TODO: instead of just printing the AST, evaluate it
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
				fmt.Println(err)
			} else {
				fmt.Println(golisp.Evaluate(parsed))
			}
		}
	}
}
