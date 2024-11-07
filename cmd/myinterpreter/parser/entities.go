package parser

import (
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/scanner"
)

type Expression[T any] interface {
	Accept(visitor ExpressionVisitor[T]) (T, error)
}

type Binary[T any] struct {
	Left     Expression[T]
	Operator scanner.Token
	Right    Expression[T]
}

func (b Binary[T]) Accept(visitor ExpressionVisitor[T]) (T, error) {
	return visitor.VisitBinaryExpr(b)
}

type Grouping[T any] struct {
	Expr Expression[T]
}

func (g Grouping[T]) Accept(visitor ExpressionVisitor[T]) (T, error) {
	return visitor.VisitGroupExpr(g)
}

type Literal[T any] struct {
	Value any
}

func (l Literal[T]) Accept(visitor ExpressionVisitor[T]) (T, error) {
	return visitor.VisitLiteralExpr(l)
}

type Unary[T any] struct {
	Operator scanner.Token
	Right    Expression[T]
}

func (u Unary[T]) Accept(visitor ExpressionVisitor[T]) (T, error) {
	return visitor.VisitUnaryExpr(u)
}

type Error struct {
	Line int
	Err  error
}
