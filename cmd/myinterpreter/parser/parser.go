package parser

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/core"
)

var tokens = []core.Token{}
var statements = []core.Statement{}
var expressions = []core.Expression{}
var position = 0

func current() core.Token {
	return tokens[position]
}

func previous() core.Token {
	return tokens[position-1]
}

func isAtEnd() bool {
	return current().Type == core.EOF
}

func advance() core.Token {
	if isAtEnd() {
		return current()
	}

	position++
	return previous()
}

func match(tokenTypes ...string) bool {
	for _, tokenType := range tokenTypes {
		if current().Type == tokenType {
			advance()
			return true
		}
	}
	return false
}

func isNextTokenSemicolon() *core.Error {
	if match(core.SEMICOLON) {
		return nil
	}

	err := fmt.Errorf("Expect ';' after expression.")
	return &core.Error{Line: current().Line, Err: err, ExitCode: 65}
}

func statement() (core.Statement, *core.Error) {
	if match(core.PRINT) {
		return printStatement()
	}
	return expressionStatement()
}

func printStatement() (core.Statement, *core.Error) {
	value, err := expression()
	if err != nil {
		return nil, err
	}

	err = isNextTokenSemicolon()
	if err != nil {
		return nil, err
	}

	return core.PrintStmt{Expr: value}, nil
}

func expressionStatement() (core.Statement, *core.Error) {
	expr, err := expression()
	if err != nil {
		return nil, err
	}

	err = isNextTokenSemicolon()
	if err != nil {
		return nil, err
	}

	return core.ExpressionStmt{Expr: expr}, nil
}

func expression() (core.Expression, *core.Error) {
	return equality()
}

func equality() (core.Expression, *core.Error) {
	expr, err := comparison()
	if err != nil {
		return nil, err
	}

	for match(core.EQUAL_EQUAL, core.BANG_EQUAL) {
		operator := previous()
		right, err := equality()
		if err != nil {
			return nil, err
		}

		expr = core.Binary{Left: expr, Operator: operator, Right: right}
	}

	return expr, nil
}

func comparison() (core.Expression, *core.Error) {
	expr, err := term()
	if err != nil {
		return nil, err
	}

	for match(core.GREATER, core.GREATER_EQUAL, core.LESS, core.LESS_EQUAL) {
		operator := previous()
		right, err := term()
		if err != nil {
			return nil, err
		}

		expr = core.Binary{Left: expr, Operator: operator, Right: right}
	}

	return expr, nil
}

func term() (core.Expression, *core.Error) {
	expr, err := factor()
	if err != nil {
		return nil, err
	}

	for match(core.MINUS, core.PLUS) {
		operator := previous()
		right, err := factor()
		if err != nil {
			return nil, err
		}

		expr = core.Binary{Left: expr, Operator: operator, Right: right}
	}

	return expr, nil
}

func factor() (core.Expression, *core.Error) {
	expr, err := unary()
	if err != nil {
		return expr, err
	}

	for match(core.SLASH, core.STAR) {
		operator := previous()
		right, err := unary()
		if err != nil {
			return expr, err
		}

		expr = core.Binary{Left: expr, Operator: operator, Right: right}
	}

	return expr, nil
}

func unary() (core.Expression, *core.Error) {
	if match(core.BANG, core.MINUS) {
		operator := previous()
		right, err := unary()
		if err != nil {
			return nil, err
		}
		return core.Unary{Operator: operator, Right: right}, nil
	}

	return primary()
}

func primary() (core.Expression, *core.Error) {
	if match(core.FALSE) {
		return core.Literal{Value: false}, nil
	}

	if match(core.TRUE) {
		return core.Literal{Value: true}, nil
	}

	if match(core.NIL) {
		return core.Literal{Value: nil}, nil
	}

	if match(core.NUMBER, core.STRING) {
		return core.Literal{Value: previous().Literal}, nil
	}

	if !match(core.LEFT_PAREN) {
		err := fmt.Errorf("Expect ')' after expression.")
		return nil, &core.Error{Line: current().Line, Err: err, ExitCode: 65}
	}

	expr, err := expression()
	if err != nil {
		return expr, err
	}

	if expr == nil {
		return nil, &core.Error{Line: current().Line, Err: fmt.Errorf("Empty group"), ExitCode: 65}
	}

	if !match(core.RIGHT_PAREN) {
		err := fmt.Errorf("Expect ')' after expression.")
		return expr, &core.Error{Line: current().Line, Err: err, ExitCode: 65}
	}

	return core.Grouping{Expr: expr}, nil
}

func Parse(scannedTokens []core.Token) ([]core.Statement, *core.Error) {
	tokens = scannedTokens
	for !isAtEnd() {
		stmt, err := statement()
		if err != nil {
			return statements, err
		}

		statements = append(statements, stmt)
	}
	return statements, nil
}

func ParseExpressions(scannedTokens []core.Token) ([]core.Expression, *core.Error) {
	tokens = scannedTokens
	for !isAtEnd() {
		expr, err := expression()
		if err != nil {
			return expressions, err
		}

		expressions = append(expressions, expr)
	}
	return expressions, nil
}
