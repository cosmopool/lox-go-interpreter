package main

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/parser"
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/scanner"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]
	filename := os.Args[2]

	switch command {
	case "tokenize":
		tokens, errors := tokenize(filename)
		printTokens(tokens)

		if len(errors) > 0 {
			printErrors(errors)
			os.Exit(65)
		}
	case "parse":
		tokens, tokenErrors := tokenize(filename)
		if len(tokenErrors) > 0 {
			printErrors(tokenErrors)
			os.Exit(65)
		}
		parser := parser.Parser{Tokens: tokens}
		expressions, err := parser.Parse()
		if err != nil {
			fmt.Fprintf(os.Stderr, "[line %d] Error: %v\n", err.Line, err.Err)
			os.Exit(65)
		}
		printExpressions(expressions)

	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}
}

func tokenize(filename string) ([]scanner.Token, []scanner.Error) {
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	return scanner.ScanFile(fileContents)
}

func printTokens(tokens []scanner.Token) {
	for _, token := range tokens {
		fmt.Print(token)
	}
}

func printErrors(errors []scanner.Error) {
	for _, err := range errors {
		fmt.Fprintf(os.Stderr, "[line %d] Error: %v\n", err.Line, err.Err)
	}
}

func printExpressions(expressions []parser.Expression[any]) {
	for _, expr := range expressions {
		fmt.Println(expr)
	}
}
