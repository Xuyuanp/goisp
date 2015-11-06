package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Xuyuanp/goisp/lexer"
)

func prompt() {
	fmt.Print("lisp>> ")
}

func repl() {
	prompt()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		tokens, err := lexer.Tokenizer(line)
		if err != nil {
			fmt.Printf("tokenizer %s failed: %s\n", line, err)
		} else if tokens != nil {
			for it := tokens.Front(); it != nil; it = it.Next() {
				fmt.Printf("%+v\n", it.Value)
			}
		}
		prompt()
	}
}

func main() {
	args := os.Args
	if len(args) > 1 {
		line := args[1]
		tokens, err := lexer.Tokenizer(line)
		if err != nil {
			fmt.Printf("tokenizer %s failed: %s\n", line, err)
		} else if tokens != nil {
			for it := tokens.Front(); it != nil; it = it.Next() {
				fmt.Printf("%+v\n", it.Value)
			}
		}
	} else {
		repl()
	}
}
