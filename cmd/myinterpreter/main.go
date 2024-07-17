package main

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/scanner"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	tokens, errors := scanner.ScanFile(fileContents)

	printTokens(tokens)
	printErrors(errors)
	if len(errors) > 0 {
		os.Exit(65)
	}
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
