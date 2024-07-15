package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
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
	NUMBER        TokenType = "NUMBER"
)

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal any
}

type Error struct {
	line int
	err  error
}

var tokens []Token
var errors []Error
var position int
var fileContents []byte
var endOfFile int

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

	var err error
	filename := os.Args[2]
	fileContents, err = os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	line := 1
	endOfFile = len(fileContents)
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
			if nextRuneEquals('=') {
				advanceCursor()
				tokens = append(tokens, Token{EQUAL_EQUAL, "==", nil})
			} else {
				tokens = append(tokens, Token{EQUAL, "=", nil})
			}
		case '!':
			if nextRuneEquals('=') {
				advanceCursor()
				tokens = append(tokens, Token{BANG_EQUAL, "!=", nil})
			} else {
				tokens = append(tokens, Token{BANG, "!", nil})
			}
		case '<':
			if nextRuneEquals('=') {
				advanceCursor()
				tokens = append(tokens, Token{LESS_EQUAL, "<=", nil})
			} else {
				tokens = append(tokens, Token{LESS, "<", nil})
			}
		case '>':
			if nextRuneEquals('=') {
				advanceCursor()
				tokens = append(tokens, Token{GREATER_EQUAL, ">=", nil})
			} else {
				tokens = append(tokens, Token{GREATER, ">", nil})
			}
		case '\t', ' ':
			// ignore whitespaces
		case '\n':
			line++
		case '/':
			if nextRuneEquals('/') {
				advanceCursor()
				for !currentRuneEquals('\n') {
					advanceCursor()
				}
				line++
			} else {
				tokens = append(tokens, Token{SLASH, "/", nil})
			}
		case '"':
			startPosition := position
			advanceCursor()

			for !currentRuneEquals('"') {
				advanceCursor()
			}

			if position >= endOfFile {
				reportError(line, fmt.Errorf("Unterminated string."))
				continue
			}

			literal := string(fileContents[startPosition+1 : position])
			lexeme := `"` + literal + `"`
			tokens = append(tokens, Token{STRING, lexeme, literal})
		default:
			if unicode.IsDigit(character) {
				tokenizeNumber()
				// the tokenizeNumber already advances the cursor
				// that's why we must go to the next iteration manually
				continue
			} else {
				reportError(line, fmt.Errorf("Unexpected character: %s", string(character)))
			}
		}

		advanceCursor()
	}

	tokens = append(tokens, Token{EOF, "", nil})

	// print all tokens
	for _, token := range tokens {
		// var name string
		// if token.Literal == nil {
		// 	name = "null"
		// } else {
			// name = fmt.Sprintf("%v", token.Literal)
		// }
		fmt.Fprintf(os.Stdout, "%v %s %v\n", token.Type, token.Lexeme, token.Literal)
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

func advanceCursor() {
	position++
}

func currentRune() rune {
	if position >= len(fileContents) {
		return -1
	}

	return rune(fileContents[position])
}

func nextRune() rune {
	nextPosition := position + 1
	if nextPosition >= len(fileContents) {
		return -1
	}

	return rune(fileContents[nextPosition])
}

func currentRuneEquals(target rune) bool {
	return currentRune() == target
}

func nextRuneEquals(target rune) bool {
	return nextRune() == target
}

func tokenizeNumber() {
	startPosition := position

	for unicode.IsDigit(currentRune()) {
		advanceCursor()
	}

	if currentRune() == '.' && unicode.IsDigit(nextRune()) {
		advanceCursor()

		for unicode.IsDigit(currentRune()) {
			advanceCursor()
		}
	}

	lexeme := string(fileContents[startPosition:position])
	literal, err := strconv.ParseFloat(lexeme, 64)
	if err != nil {
		panic(err)
	}
	token := Token{NUMBER, lexeme, literal}
	tokens = append(tokens, token)
}
