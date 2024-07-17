package parser

import (
	"reflect"
	"testing"

	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/scanner"
)

// isAtEnd() should return FALSE if NOT reached EOF token
func TestIsAtEndReturnsFalseIfNotEOFToken(t *testing.T) {
	tokens := []scanner.Token{
		{Type: scanner.NUMBER, Lexeme: "2", Literal: 2},
	}

	parser := Parser{Tokens: tokens}
	if parser.isAtEnd() == true {
		t.Fatal("isAtEnd() should return FALSE if NOT reached EOF token")
	}
}

// isAtEnd() should return TRUE if reached EOF token
func TestIsAtEndReturnsTrueReachedEOFToken(t *testing.T) {
	tokens := []scanner.Token{
		{Type: scanner.EOF, Lexeme: "EOF", Literal: nil},
	}

	parser := Parser{Tokens: tokens}
	if parser.isAtEnd() == false {
		t.Fatal("isAtEnd() should return if reached EOF token")
	}
}

// advance() should increase position if not reached the end tokens[]
func TestAdvanceShouldIncreasePositionWhenNotOnEnd(t *testing.T) {
	tokens := []scanner.Token{
		{Type: scanner.NUMBER, Lexeme: "2", Literal: 2},
		{Type: scanner.EOF, Lexeme: "EOF", Literal: nil},
	}

	parser := Parser{Tokens: tokens}
	if parser.position != 0 {
		t.Fatalf("position shoul be 0, but found: %v", parser.position)
	}

	token := parser.advance()
	if parser.position != 1 {
		t.Fatalf("position shoul be 1, but found: %v", parser.position)
	}
	if !reflect.DeepEqual(token, tokens[0]) {
		t.Fatalf("expecting first advance to return: %v, but found: %v", tokens[0], token)
	}

	token = parser.advance()
	if !reflect.DeepEqual(token, tokens[1]) {
		t.Fatalf("expecting second advance to return: %v, but found: %v", tokens[1], token)
	}
	if parser.position != 1 {
		t.Fatalf("position shoul be 1, but found: %v", parser.position)
	}
}

func TestCallingTermShouldIncreasePosition(t *testing.T) {
	tokens := []scanner.Token{
		{Type: scanner.NUMBER, Lexeme: "2", Literal: 2},
		{Type: scanner.EQUAL, Lexeme: "=", Literal: nil},
		{Type: scanner.NUMBER, Lexeme: "3", Literal: 3},
	}

	parser := Parser{Tokens: tokens}
	parser.term()
	if parser.position == 0 {
		t.Fatal("term() should increase position")
	}
}
