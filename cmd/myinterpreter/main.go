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
		expressions, err := parser.ParseExpressions(tokens)
		if err != nil {
			printErrorAndExit(err)
			os.Exit(65)
		}

		// visit expressions
		printer := visitor.PrinterVisitor{}
		for _, expr := range expressions {
			printer.PrintExpression(expr)
		}

	case "evaluate":
		tokens := tokenize(filename, false)
		expressions, err := parser.ParseExpressions(tokens)
		if err != nil {
			printErrorAndExit(err)
			os.Exit(65)
		}

		// visit expressions
		evaluator := visitor.Evaluator{}
		for _, expr := range expressions {
			value, err := evaluator.Evaluate(expr)
			printErrorAndExit(&err)

			if value == nil {
				fmt.Println("nil")
			} else {
				fmt.Println(value)
			}
		}

	case "run":
		tokens := tokenize(filename, false)
		statements, err := parser.Parse(tokens)
		printErrorAndExit(err)

		// visit expressions
		interpreter := visitor.Interpreter{}
		for _, expr := range statements {
			_, err := interpreter.Interpret(expr)
			printErrorAndExit(&err)
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

	printErrorsAndExit(errors)

	return tokens
}

func printTokens(tokens []core.Token) {
	for _, token := range tokens {
		fmt.Print(token)
	}
}

func printErrorsAndExit(errors []core.Error) {
	if len(errors) == 0 {
		return
	}

	for _, err := range errors {
		fmt.Fprintf(os.Stderr, "[line %d] Error: %v\n", err.Line, err.Err)
	}

	os.Exit(errors[0].ExitCode)
}

func printErrorAndExit(error *core.Error) {
	if error == nil {
		return
	}
	if error.Err == nil {
		return
	}

	fmt.Fprintf(os.Stderr, "[line %d] Error: %v\n", error.Line, error.Err)
	os.Exit(error.ExitCode)
}
