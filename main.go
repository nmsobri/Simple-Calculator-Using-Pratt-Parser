package main

import (
	"ExprParser/lexer"
	"ExprParser/parser"
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome To Simple Calculator")

	for {
		fmt.Print(">> ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()

		lexer := lexer.New(line)
		parser := parser.New(lexer)

		result := parser.ParseProgram()
		fmt.Printf("%s = %d\n", line, result)
	}
}
