package main

import (
	"fmt"
	"os"
)

type TokenType = string

const (
	LEFT_PAREN    TokenType = "LEFT_PAREN"
	RIGHT_PAREN   TokenType = "RIGHT_PAREN"
	LEFT_BRACE    TokenType = "LEFT_BRACE"
	RIGHT_BRACE   TokenType = "RIGHT_BRACE"
	COMMA         TokenType = "COMMA"
	DOT           TokenType = "DOT"
	MINUS         TokenType = "MINUS"
	PLUS          TokenType = "PLUS"
	SEMICOLON     TokenType = "SEMICOLON"
	SLASH         TokenType = "SLASH"
	STAR          TokenType = "STAR"
	BANG          TokenType = "BANG"
	BANG_EQUAL    TokenType = "BANG_EQUAL"
	EQUAL         TokenType = "EQUAL"
	EQUAL_EQUAL   TokenType = "EQUAL_EQUAL"
	GREATER       TokenType = "GREATER"
	GREATER_EQUAL TokenType = "GREATER_EQUAL"
	LESS          TokenType = "LESS"
	LESS_EQUAL    TokenType = "LESS_EQUAL"
	EOF           TokenType = "EOF"
	STRING        TokenType = "STRING"
	COMMENT                 = "COMMENT"
	NONE                    = "NONE"
)

type Token struct {
	Type    TokenType
	Literal string
	Name    any
}

type Error struct {
	line int
	err  error
}

var tokens []Token
var errors []Error

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

	var position int
	line := 1
	endOfFile := len(fileContents)
	for position < endOfFile {
		character := rune(fileContents[position])

		switch character {
		case '(':
			tokens = append(tokens, Token{LEFT_PAREN, "(", nil})
		case ')':
			tokens = append(tokens, Token{RIGHT_PAREN, ")", nil})
		case '{':
			tokens = append(tokens, Token{LEFT_BRACE, "{", nil})
		case '}':
			tokens = append(tokens, Token{RIGHT_BRACE, "}", nil})
		case '*':
			tokens = append(tokens, Token{STAR, "*", nil})
		case '.':
			tokens = append(tokens, Token{DOT, ".", nil})
		case ',':
			tokens = append(tokens, Token{COMMA, ",", nil})
		case '+':
			tokens = append(tokens, Token{PLUS, "+", nil})
		case '-':
			tokens = append(tokens, Token{MINUS, "-", nil})
		case ';':
			tokens = append(tokens, Token{SEMICOLON, ";", nil})
		case '=':
			if moveToNextRuneIfEqualsTo('=', &position, fileContents) {
				tokens = append(tokens, Token{EQUAL_EQUAL, "==", nil})
			} else {
				tokens = append(tokens, Token{EQUAL, "=", nil})
			}
		case '!':
			if moveToNextRuneIfEqualsTo('=', &position, fileContents) {
				tokens = append(tokens, Token{BANG_EQUAL, "!=", nil})
			} else {
				tokens = append(tokens, Token{BANG, "!", nil})
			}
		case '<':
			if moveToNextRuneIfEqualsTo('=', &position, fileContents) {
				tokens = append(tokens, Token{LESS_EQUAL, "<=", nil})
			} else {
				tokens = append(tokens, Token{LESS, "<", nil})
			}
		case '>':
			if moveToNextRuneIfEqualsTo('=', &position, fileContents) {
				tokens = append(tokens, Token{GREATER_EQUAL, ">=", nil})
			} else {
				tokens = append(tokens, Token{GREATER, ">", nil})
			}
		case '/':
			if moveToNextRuneIfEqualsTo('/', &position, fileContents) {
				var char rune
				for char != '\n' {
					position++
					if position == endOfFile {
						break
					}

					char = rune(fileContents[position])
				}
        line++
			} else {
				tokens = append(tokens, Token{SLASH, "/", nil})
			}
		case '#':
			reportError(line, fmt.Errorf("Unexpected character: %s", string(character)))
		case '$':
			reportError(line, fmt.Errorf("Unexpected character: %s", string(character)))
		case '@':
			reportError(line, fmt.Errorf("Unexpected character: %s", string(character)))
		case '%':
			reportError(line, fmt.Errorf("Unexpected character: %s", string(character)))
		case '\t', ' ':
			// ignore whitespaces
		case '\n':
			line++
		case '"':
			startPosition := position
			position++

			var char rune
			for char != '"' {
				position++
				if position >= endOfFile {
					break
				}

				char = rune(fileContents[position])
			}

			if position >= endOfFile {
				reportError(line, fmt.Errorf("Unterminated string."))
				continue
			}

			literal := string(fileContents[startPosition+1 : position])
			tokens = append(tokens, Token{STRING, `"` + literal + `"`, literal})
		default:
		}

		position++
	}

	tokens = append(tokens, Token{EOF, "", nil})

	// print all tokens
	for _, token := range tokens {
		var name string
		if token.Name == nil {
			name = "null"
		} else {
			name = fmt.Sprintf("%v", token.Name)
		}
		fmt.Fprintf(os.Stdout, "%v %s %s\n", token.Type, token.Literal, name)
	}

	// check for errors and print them all
	if len(errors) > 0 {
		for _, err := range errors {
			fmt.Fprintf(os.Stderr, "[line %d] Error: %v\n", err.line, err.err)
		}
		os.Exit(65)
	}

}

func reportError(line int, err error) {
	errors = append(errors, Error{line, err})
}

// Returns a boolean if next [position] rune is equal to [targetChar].
// If true, will move [position] by adding +1.
func moveToNextRuneIfEqualsTo(targetRune rune, position *int, content []byte) bool {
	// check if it's within bounds
	if *position >= len(content)-1 {
		return false
	}

	isEqual := targetRune == rune(content[*position+1])
	if !isEqual {
		return false
	}

	// move scanner to the end of this token
	*position += 1
	return true
}
