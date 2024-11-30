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

	exprPrinterVisitor := visitor.PrinterVisitor{}

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
			printError(*err)
			os.Exit(65)
		}

		// visit expressions
		for _, expr := range expressions {
			expr.Accept(exprPrinterVisitor)
		}

	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}
}

func tokenize(filename string) ([]core.Token, []core.Error) {
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	return scanner.ScanFile(fileContents)
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
