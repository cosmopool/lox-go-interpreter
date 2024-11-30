package scanner

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/core"
)

var tokens []core.Token
var errors []core.Error
var position int
var contents []byte
var endOfFile int
var line int

func ScanFile(fileContents []byte) ([]core.Token, []core.Error) {
	contents = fileContents
	endOfFile = len(contents)
	line = 1

	for position < endOfFile {
		character := rune(contents[position])

		switch character {
		case '(':
			tokens = append(tokens, core.Token{Type: core.LEFT_PAREN, Lexeme: "(", Literal: nil, Line: line})
		case ')':
			tokens = append(tokens, core.Token{Type: core.RIGHT_PAREN, Lexeme: ")", Literal: nil, Line: line})
		case '{':
			tokens = append(tokens, core.Token{Type: core.LEFT_BRACE, Lexeme: "{", Literal: nil, Line: line})
		case '}':
			tokens = append(tokens, core.Token{Type: core.RIGHT_BRACE, Lexeme: "}", Literal: nil, Line: line})
		case '*':
			tokens = append(tokens, core.Token{Type: core.STAR, Lexeme: "*", Literal: nil, Line: line})
		case '.':
			tokens = append(tokens, core.Token{Type: core.DOT, Lexeme: ".", Literal: nil, Line: line})
		case ',':
			tokens = append(tokens, core.Token{Type: core.COMMA, Lexeme: ",", Literal: nil, Line: line})
		case '+':
			tokens = append(tokens, core.Token{Type: core.PLUS, Lexeme: "+", Literal: nil, Line: line})
		case '-':
			tokens = append(tokens, core.Token{Type: core.MINUS, Lexeme: "-", Literal: nil, Line: line})
		case ';':
			tokens = append(tokens, core.Token{Type: core.SEMICOLON, Lexeme: ";", Literal: nil, Line: line})
		case '=':
			if nextRuneEquals('=') {
				advanceCursor()
				tokens = append(tokens, core.Token{Type: core.EQUAL_EQUAL, Lexeme: "==", Literal: nil, Line: line})
			} else {
				tokens = append(tokens, core.Token{Type: core.EQUAL, Lexeme: "=", Literal: nil, Line: line})
			}
		case '!':
			if nextRuneEquals('=') {
				advanceCursor()
				tokens = append(tokens, core.Token{Type: core.BANG_EQUAL, Lexeme: "!=", Literal: nil, Line: line})
			} else {
				tokens = append(tokens, core.Token{Type: core.BANG, Lexeme: "!", Literal: nil, Line: line})
			}
		case '<':
			if nextRuneEquals('=') {
				advanceCursor()
				tokens = append(tokens, core.Token{Type: core.LESS_EQUAL, Lexeme: "<=", Literal: nil, Line: line})
			} else {
				tokens = append(tokens, core.Token{Type: core.LESS, Lexeme: "<", Literal: nil, Line: line})
			}
		case '>':
			if nextRuneEquals('=') {
				advanceCursor()
				tokens = append(tokens, core.Token{Type: core.GREATER_EQUAL, Lexeme: ">=", Literal: nil, Line: line})
			} else {
				tokens = append(tokens, core.Token{Type: core.GREATER, Lexeme: ">", Literal: nil, Line: line})
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
				tokens = append(tokens, core.Token{Type: core.SLASH, Lexeme: "/", Literal: nil, Line: line})
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
			tokens = append(tokens, core.Token{Type: core.STRING, Lexeme: lexeme, Literal: literal, Line: line})
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

	tokens = append(tokens, core.Token{Type: core.EOF, Lexeme: "", Literal: nil, Line: line})

	return tokens, errors
}

func reportError(line int, err error) {
	errors = append(errors, core.Error{Line: line, Err: err})
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
	token := core.Token{Type: core.NUMBER, Lexeme: lexeme, Literal: literal, Line: line}
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
	keyword := core.Keywords()[lexeme]
	if keyword == "" {
		tokenType = core.IDENTIFIER
	} else {
		tokenType = keyword
	}

	tokens = append(tokens, core.Token{Type: tokenType, Lexeme: lexeme, Literal: nil, Line: line})
}
