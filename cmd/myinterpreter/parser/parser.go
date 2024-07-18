package parser

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/scanner"
)

type Parser struct {
	Tokens      []scanner.Token
	expressions []Expression
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

func (p *Parser) expression() (Expression, error) {
	return p.equality()
}

func (p *Parser) equality() (Expression, error) {
	expr, err := p.comparison()
	if err != nil {
		return nil, err
	}

	for p.match(scanner.EQUAL_EQUAL, scanner.BANG_EQUAL) {
		operator := p.previous()
		right, err := p.equality()
		if err != nil {
			return nil, err
		}

		expr = Binary{Left: expr, Operator: operator, Right: right}
	}

	return expr, nil
}

func (p *Parser) comparison() (Expression, error) {
	expr, err := p.term()
	if err != nil {
		return nil, err
	}

	for p.match(scanner.GREATER, scanner.GREATER_EQUAL, scanner.LESS, scanner.LESS_EQUAL) {
		operator := p.previous()
		right, err := p.term()
		if err != nil {
			return nil, err
		}

		expr = Binary{Left: expr, Operator: operator, Right: right}
	}

	return expr, nil
}

func (p *Parser) term() (Expression, error) {
	expr, err := p.factor()
	if err != nil {
		return nil, err
	}

	for p.match(scanner.MINUS, scanner.PLUS) {
		operator := p.previous()
		right, err := p.factor()
		if err != nil {
			return nil, err
		}

		expr = Binary{Left: expr, Operator: operator, Right: right}
	}

	return expr, nil
}

func (p *Parser) factor() (Expression, error) {
	expr, err := p.unary()
	if err != nil {
		return expr, err
	}

	for p.match(scanner.SLASH, scanner.STAR) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return expr, err
		}

		expr = Binary{Left: expr, Operator: operator, Right: right}
	}

	return expr, nil
}

func (p *Parser) unary() (Expression, error) {
	if p.match(scanner.BANG, scanner.MINUS) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		return Unary{Operator: operator, Right: right}, nil
	}

	return p.primary()
}

func (p *Parser) primary() (Expression, error) {
	if p.match(scanner.FALSE) {
		return Literal{Value: false}, nil
	}

	if p.match(scanner.TRUE) {
		return Literal{Value: true}, nil
	}

	if p.match(scanner.NIL) {
		return Literal{Value: nil}, nil
	}

	if p.match(scanner.NUMBER, scanner.STRING) {
		return Literal{Value: p.previous().Literal}, nil
	}

	if !p.match(scanner.LEFT_PAREN) {
		return nil, nil
	}

	expr, err := p.expression()
	if err != nil {
		return expr, err
	}

	if expr == nil {
		return nil, fmt.Errorf("Empty group")
	}

	if !p.match(scanner.RIGHT_PAREN) {
		return expr, fmt.Errorf("Expect ')' after expression.")
	}

	return Grouping{Expr: expr}, nil
}

func (p *Parser) Parse() ([]Expression, error) {
	for !p.isAtEnd() {
		expr, err := p.expression()
		if err != nil {
			return p.expressions, err
		}

		p.expressions = append(p.expressions, expr)
	}
	return p.expressions, nil
}
