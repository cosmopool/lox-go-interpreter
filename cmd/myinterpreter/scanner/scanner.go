package scanner

import (
	"fmt"
	"strconv"
	"unicode"
)

type tokenType = string

const (
	LEFT_PAREN    tokenType = "LEFT_PAREN"
	RIGHT_PAREN   tokenType = "RIGHT_PAREN"
	LEFT_BRACE    tokenType = "LEFT_BRACE"
	RIGHT_BRACE   tokenType = "RIGHT_BRACE"
	COMMA         tokenType = "COMMA"
	DOT           tokenType = "DOT"
	MINUS         tokenType = "MINUS"
	PLUS          tokenType = "PLUS"
	SEMICOLON     tokenType = "SEMICOLON"
	SLASH         tokenType = "SLASH"
	STAR          tokenType = "STAR"
	BANG          tokenType = "BANG"
	BANG_EQUAL    tokenType = "BANG_EQUAL"
	EQUAL         tokenType = "EQUAL"
	EQUAL_EQUAL   tokenType = "EQUAL_EQUAL"
	GREATER       tokenType = "GREATER"
	GREATER_EQUAL tokenType = "GREATER_EQUAL"
	LESS          tokenType = "LESS"
	LESS_EQUAL    tokenType = "LESS_EQUAL"
	EOF           tokenType = "EOF"
	STRING        tokenType = "STRING"
	NUMBER        tokenType = "NUMBER"
	IDENTIFIER    tokenType = "IDENTIFIER"
)

func keywords() map[string]tokenType {
	return map[string]tokenType{
		"and":    "AND",
		"class":  "CLASS",
		"else":   "ELSE",
		"false":  "FALSE",
		"for":    "FOR",
		"fun":    "FUN",
		"if":     "IF",
		"nil":    "NIL",
		"or":     "OR",
		"print":  "PRINT",
		"return": "RETURN",
		"super":  "SUPER",
		"this":   "THIS",
		"true":   "TRUE",
		"var":    "VAR",
		"while":  "WHILE",
	}
}

type Token struct {
	Type    tokenType
	Lexeme  string
	Literal any
}

type Error struct {
	Line int
	Err  error
}

var tokens []Token
var errors []Error
var position int
var contents []byte
var endOfFile int

func ScanFile(fileContents []byte) ([]Token, []Error) {
	contents = fileContents
	endOfFile = len(contents)
	line := 1

	for position < endOfFile {
		character := rune(contents[position])

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
					if position >= endOfFile {
						break
					}
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
				if position >= endOfFile {
					break
				}
			}

			if position >= endOfFile {
				reportError(line, fmt.Errorf("Unterminated string."))
				break
			}

			literal := string(contents[startPosition+1 : position])
			lexeme := `"` + literal + `"`
			tokens = append(tokens, Token{STRING, lexeme, literal})
		default:
			if unicode.IsDigit(character) {
				tokenizeNumber()
				// the tokenizeNumber already advances the cursor
				// that's why we must go to the next iteration manually
				continue
			} else if unicode.IsLetter(character) || character == '_' {
				tokenizeIdentifier()
				continue
			} else {
				reportError(line, fmt.Errorf("Unexpected character: %s", string(character)))
			}
		}

		advanceCursor()
	}

	tokens = append(tokens, Token{EOF, "", nil})

	return tokens, errors
}

func reportError(line int, err error) {
	errors = append(errors, Error{line, err})
}

func advanceCursor() {
	position++
}

func currentRune() rune {
	if position >= len(contents) {
		return -1
	}

	return rune(contents[position])
}

func nextRune() rune {
	nextPosition := position + 1
	if nextPosition >= len(contents) {
		return -1
	}

	return rune(contents[nextPosition])
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

	lexeme := string(contents[startPosition:position])
	literal, err := strconv.ParseFloat(lexeme, 64)
	if err != nil {
		panic(err)
	}
	token := Token{NUMBER, lexeme, literal}
	tokens = append(tokens, token)
}

func isAlphaNumeric(char rune) bool {
	return unicode.IsLetter(char) || unicode.IsDigit(char) || char == '_'
}

func tokenizeIdentifier() {
	startPos := position
	for isAlphaNumeric(currentRune()) {
		advanceCursor()
	}

	lexeme := string(contents[startPos:position])

	var tokenType string
	keyword := keywords()[lexeme]
	if keyword == "" {
		tokenType = IDENTIFIER
	} else {
		tokenType = keyword
	}

	tokens = append(tokens, Token{tokenType, lexeme, nil})
}
