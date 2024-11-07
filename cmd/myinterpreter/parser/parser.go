package parser

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/scanner"
)

type Parser struct {
	Tokens      []scanner.Token
	expressions []Expression[any]
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

func (p *Parser) expression() (Expression[any], *scanner.Error) {
	return p.equality()
}

func (p *Parser) equality() (Expression[any], *scanner.Error) {
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

		expr = Binary[any]{Left: expr, Operator: operator, Right: right}
	}

	return expr, nil
}

func (p *Parser) comparison() (Expression[any], *scanner.Error) {
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

		expr = Binary[any]{Left: expr, Operator: operator, Right: right}
	}

	return expr, nil
}

func (p *Parser) term() (Expression[any], *scanner.Error) {
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

		expr = Binary[any]{Left: expr, Operator: operator, Right: right}
	}

	return expr, nil
}

func (p *Parser) factor() (Expression[any], *scanner.Error) {
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

		expr = Binary[any]{Left: expr, Operator: operator, Right: right}
	}

	return expr, nil
}

func (p *Parser) unary() (Expression[any], *scanner.Error) {
	if p.match(scanner.BANG, scanner.MINUS) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		return Unary[any]{Operator: operator, Right: right}, nil
	}

	return p.primary()
}

func (p *Parser) primary() (Expression[any], *scanner.Error) {
	if p.match(scanner.FALSE) {
		return Literal[any]{Value: false}, nil
	}

	if p.match(scanner.TRUE) {
		return Literal[any]{Value: true}, nil
	}

	if p.match(scanner.NIL) {
		return Literal[any]{Value: nil}, nil
	}

	if p.match(scanner.NUMBER, scanner.STRING) {
		return Literal[any]{Value: p.previous().Literal}, nil
	}

	if !p.match(scanner.LEFT_PAREN) {
		err := fmt.Errorf("Expect ')' after expression.")
		return nil, &scanner.Error{Line: p.current().Line, Err: err}
	}

	expr, err := p.expression()
	if err != nil {
		return expr, err
	}

	if expr == nil {
		return nil, &scanner.Error{Line: p.current().Line, Err: fmt.Errorf("Empty group")}
	}

	if !p.match(scanner.RIGHT_PAREN) {
		err := fmt.Errorf("Expect ')' after expression.")
		return expr, &scanner.Error{Line: p.current().Line, Err: err}
	}

	return Grouping[any]{Expr: expr}, nil
}

func (p *Parser) Parse() ([]Expression[any], *scanner.Error) {
	for !p.isAtEnd() {
		expr, err := p.expression()
		if err != nil {
			return p.expressions, err
		}

		p.expressions = append(p.expressions, expr)
	}
	return p.expressions, nil
}
