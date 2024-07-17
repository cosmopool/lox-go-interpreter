package parser

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/scanner"
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/utils"
)

type Expression interface {
	Expression()
}

type Binary struct {
	Left     Expression
	Operator scanner.Token
	Right    Expression
}

func (b Binary) String() string {
	return fmt.Sprintf("(%s (%s %s))", b.Operator.Lexeme, b.Left, b.Right)
}

func (b Binary) Expression() {}

type Grouping struct {
	Expr Expression
}

func (g Grouping) Expression() {}
func (g Grouping) String() string {
	return fmt.Sprintf("(group %s)", g.Expr)
}

type Literal struct {
	Value any
}

func (l Literal) Expression() {}

func (l Literal) String() string {
	return utils.VariableToString(l.Value, false)
}

type Unary struct {
	Operator scanner.Token
	Right    Expression
}

func (u Unary) Expression() {}
func (u Unary) String() string {
	return fmt.Sprintf("(%s %v)", u.Operator.Lexeme, u.Right)
}

type Error struct {
	Line int
	Err  error
}
