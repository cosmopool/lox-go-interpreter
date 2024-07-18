package scanner

import (
	"fmt"
	"strconv"
	"unicode"
)

var tokens []Token
var errors []Error
var position int
var contents []byte
var endOfFile int
var line int

func ScanFile(fileContents []byte) ([]Token, []Error) {
	contents = fileContents
	endOfFile = len(contents)
	line = 1

	for position < endOfFile {
		character := rune(contents[position])

		switch character {
		case '(':
			tokens = append(tokens, Token{LEFT_PAREN, "(", nil, line})
		case ')':
			tokens = append(tokens, Token{RIGHT_PAREN, ")", nil, line})
		case '{':
			tokens = append(tokens, Token{LEFT_BRACE, "{", nil, line})
		case '}':
			tokens = append(tokens, Token{RIGHT_BRACE, "}", nil, line})
		case '*':
			tokens = append(tokens, Token{STAR, "*", nil, line})
		case '.':
			tokens = append(tokens, Token{DOT, ".", nil, line})
		case ',':
			tokens = append(tokens, Token{COMMA, ",", nil, line})
		case '+':
			tokens = append(tokens, Token{PLUS, "+", nil, line})
		case '-':
			tokens = append(tokens, Token{MINUS, "-", nil, line})
		case ';':
			tokens = append(tokens, Token{SEMICOLON, ";", nil, line})
		case '=':
			if nextRuneEquals('=') {
				advanceCursor()
				tokens = append(tokens, Token{EQUAL_EQUAL, "==", nil, line})
			} else {
				tokens = append(tokens, Token{EQUAL, "=", nil, line})
			}
		case '!':
			if nextRuneEquals('=') {
				advanceCursor()
				tokens = append(tokens, Token{BANG_EQUAL, "!=", nil, line})
			} else {
				tokens = append(tokens, Token{BANG, "!", nil, line})
			}
		case '<':
			if nextRuneEquals('=') {
				advanceCursor()
				tokens = append(tokens, Token{LESS_EQUAL, "<=", nil, line})
			} else {
				tokens = append(tokens, Token{LESS, "<", nil, line})
			}
		case '>':
			if nextRuneEquals('=') {
				advanceCursor()
				tokens = append(tokens, Token{GREATER_EQUAL, ">=", nil, line})
			} else {
				tokens = append(tokens, Token{GREATER, ">", nil, line})
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
				tokens = append(tokens, Token{SLASH, "/", nil, line})
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
			tokens = append(tokens, Token{STRING, lexeme, literal, line})
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

	tokens = append(tokens, Token{EOF, "", nil, line})

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
	token := Token{NUMBER, lexeme, literal, line}
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
	keyword := Keywords()[lexeme]
	if keyword == "" {
		tokenType = IDENTIFIER
	} else {
		tokenType = keyword
	}

	tokens = append(tokens, Token{tokenType, lexeme, nil, line})
}
