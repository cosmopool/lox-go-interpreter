package main

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/core"
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/parser"
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/scanner"
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/visitor"
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
		tokenize(filename, true)

	case "parse":
		tokens := tokenize(filename, false)
		parser := parser.Parser{Tokens: tokens}
		expressions, err := parser.Parse()
		if err != nil {
			printError(*err)
			os.Exit(65)
		}

		// visit expressions
		visitor := visitor.PrinterVisitor{}
		for _, expr := range expressions {
			expr.Accept(visitor)
		}

	case "evaluate":
		tokens := tokenize(filename, false)
		parser := parser.Parser{Tokens: tokens}
		expressions, err := parser.Parse()
		if err != nil {
			printError(*err)
			os.Exit(65)
		}

		// visit expressions
		visitor := visitor.EvaluatorVisitor{}
		for _, expr := range expressions {
			value, err := expr.Accept(visitor)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v", err)
				os.Exit(70)
			}

			if value == nil {
				fmt.Println("nil")
			} else {
				fmt.Println(value)
			}
		}

	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}
}

func tokenize(filename string, shouldPrintTokens bool) []core.Token {
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	tokens, errors := scanner.ScanFile(fileContents)
	if shouldPrintTokens {
		printTokens(tokens)
	}

	if len(errors) > 0 {
		printErrors(errors)
		os.Exit(65)
	}

	return tokens
}

func printTokens(tokens []core.Token) {
	for _, token := range tokens {
		fmt.Print(token)
	}
}

func printErrors(errors []core.Error) {
	for _, err := range errors {
		printError(err)
	}
}

func printError(error core.Error) {
	fmt.Fprintf(os.Stderr, "[line %d] Error: %v\n", error.Line, error.Err)
}
