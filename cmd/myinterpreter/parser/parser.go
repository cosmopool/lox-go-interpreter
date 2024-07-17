package parser

import (
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/scanner"
)

type Parser struct {
	Tokens      []scanner.Token
	expressions []Expression
	errors      []Error
	position    int
}

func (p *Parser) current() scanner.Token {
	return p.Tokens[p.position]
}

func (p *Parser) previous() scanner.Token {
	return p.Tokens[p.position-1]
}

func (p *Parser) isAtEnd() bool {
	return p.current().Type == scanner.EOF
}

func (p *Parser) advance() scanner.Token {
	if p.isAtEnd() {
		return p.current()
	}

	p.position++
	return p.previous()
}

func (p *Parser) match(tokenTypes ...string) bool {
	for _, tokenType := range tokenTypes {
		if p.current().Type == tokenType {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) term() Expression {
	expr := p.primary()

	for p.match(scanner.MINUS, scanner.PLUS) {
		operator := p.previous()
		right := p.primary()
		expr = Binary{Left: expr, Operator: operator, Right: right}
	}

	return expr
}

func (p *Parser) primary() Expression {
	if p.match(scanner.FALSE) {
		return Literal{Value: false}
	}

	if p.match(scanner.TRUE) {
		return Literal{Value: true}
	}

	if p.match(scanner.NIL) {
		return Literal{Value: nil}
	}

	return Literal{Value: p.advance().Literal}
}

func (p *Parser) Parse() ([]Expression, []Error) {
	for !p.isAtEnd() {
		p.expressions = append(p.expressions, p.term())
	}
	return p.expressions, p.errors
}
