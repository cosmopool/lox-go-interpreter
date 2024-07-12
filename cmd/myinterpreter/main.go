package main

import (
	"fmt"
	"os"
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

	hasError := false
	var position int
	line := 1
	endOfFile := len(fileContents)
	for {
		if position >= endOfFile {
			fmt.Println("EOF  null")
			break
		}

		character := rune(fileContents[position])
		switch character {
		case '\n':
			line++
		case '(':
			fmt.Println("LEFT_PAREN ( null")
		case ')':
			fmt.Println("RIGHT_PAREN ) null")
		case '{':
			fmt.Println("LEFT_BRACE { null")
		case '}':
			fmt.Println("RIGHT_BRACE } null")
		case '*':
			fmt.Println("STAR * null")
		case '.':
			fmt.Println("DOT . null")
		case ',':
			fmt.Println("COMMA , null")
		case '+':
			fmt.Println("PLUS + null")
		case '-':
			fmt.Println("MINUS - null")
		case ';':
			fmt.Println("SEMICOLON ; null")
		case '=':
			if nextCharIsEqual('=', &position, fileContents) {
				fmt.Println("EQUAL_EQUAL == null")
			} else {
				fmt.Println("EQUAL = null")
			}
    case '!':
			if nextCharIsEqual('=', &position, fileContents) {
				fmt.Println("BANG_EQUAL != null")
			} else {
				fmt.Println("BANG ! null")
			}
    case '<':
			if nextCharIsEqual('=', &position, fileContents) {
				fmt.Println("LESS_EQUAL <= null")
			} else {
				fmt.Println("LESS < null")
			}
    case '>':
			if nextCharIsEqual('=', &position, fileContents) {
				fmt.Println("GREATER_EQUAL >= null")
			} else {
				fmt.Println("GREATER > null")
			}
		default:
			fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected character: %s\n", line, string(character))
			hasError = true
		}
		position++
	}

	if hasError {
		os.Exit(65)
	}

}

func nextCharIsEqual(char rune, position *int, content []byte) bool {
  // check if it's within bounds
	if *position >= len(content) - 1 {
		return false
	}

	isEqual := char == rune(content[*position+1])
	if !isEqual {
		return false
	}

	// move scanner to the end of this token
	*position += 1
	return true
}
