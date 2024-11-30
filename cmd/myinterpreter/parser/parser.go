package parser

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/core"
)

type Parser struct {
	Tokens      []core.Token
	expressions []core.Expression
	position    int
}

func (p *Parser) current() core.Token {
	return p.Tokens[p.position]
}

func (p *Parser) previous() core.Token {
	return p.Tokens[p.position-1]
}

func (p *Parser) isAtEnd() bool {
	return p.current().Type == core.EOF
}

func (p *Parser) advance() core.Token {
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

func (p *Parser) expression() (core.Expression, *core.Error) {
	return p.equality()
}

func (p *Parser) equality() (core.Expression, *core.Error) {
	expr, err := p.comparison()
	if err != nil {
		return nil, err
	}

	for p.match(core.EQUAL_EQUAL, core.BANG_EQUAL) {
		operator := p.previous()
		right, err := p.equality()
		if err != nil {
			return nil, err
		}

		expr = core.Binary{Left: expr, Operator: operator, Right: right}
	}

	return expr, nil
}

func (p *Parser) comparison() (core.Expression, *core.Error) {
	expr, err := p.term()
	if err != nil {
		return nil, err
	}

	for p.match(core.GREATER, core.GREATER_EQUAL, core.LESS, core.LESS_EQUAL) {
		operator := p.previous()
		right, err := p.term()
		if err != nil {
			return nil, err
		}

		expr = core.Binary{Left: expr, Operator: operator, Right: right}
	}

	return expr, nil
}

func (p *Parser) term() (core.Expression, *core.Error) {
	expr, err := p.factor()
	if err != nil {
		return nil, err
	}

	for p.match(core.MINUS, core.PLUS) {
		operator := p.previous()
		right, err := p.factor()
		if err != nil {
			return nil, err
		}

		expr = core.Binary{Left: expr, Operator: operator, Right: right}
	}

	return expr, nil
}

func (p *Parser) factor() (core.Expression, *core.Error) {
	expr, err := p.unary()
	if err != nil {
		return expr, err
	}

	for p.match(core.SLASH, core.STAR) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return expr, err
		}

		expr = core.Binary{Left: expr, Operator: operator, Right: right}
	}

	return expr, nil
}

func (p *Parser) unary() (core.Expression, *core.Error) {
	if p.match(core.BANG, core.MINUS) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		return core.Unary{Operator: operator, Right: right}, nil
	}

	return p.primary()
}

func (p *Parser) primary() (core.Expression, *core.Error) {
	if p.match(core.FALSE) {
		return core.Literal{Value: false}, nil
	}

	if p.match(core.TRUE) {
		return core.Literal{Value: true}, nil
	}

	if p.match(core.NIL) {
		return core.Literal{Value: nil}, nil
	}

	if p.match(core.NUMBER, core.STRING) {
		return core.Literal{Value: p.previous().Literal}, nil
	}

	if !p.match(core.LEFT_PAREN) {
		err := fmt.Errorf("Expect ')' after expression.")
		return nil, &core.Error{Line: p.current().Line, Err: err}
	}

	expr, err := p.expression()
	if err != nil {
		return expr, err
	}

	if expr == nil {
		return nil, &core.Error{Line: p.current().Line, Err: fmt.Errorf("Empty group")}
	}

	if !p.match(core.RIGHT_PAREN) {
		err := fmt.Errorf("Expect ')' after expression.")
		return expr, &core.Error{Line: p.current().Line, Err: err}
	}

	return core.Grouping{Expr: expr}, nil
}

func (p *Parser) Parse() ([]core.Expression, *core.Error) {
	for !p.isAtEnd() {
		expr, err := p.expression()
		if err != nil {
			return p.expressions, err
		}

		p.expressions = append(p.expressions, expr)
	}
	return p.expressions, nil
}
