package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
			parsed := ParseLisp(input)
			fmt.Println(parsed)
		}
	}

	fmt.Println(ParseLisp("(+ 1 2)"))
}
