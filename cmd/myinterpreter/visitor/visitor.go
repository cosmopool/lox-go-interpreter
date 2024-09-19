package utils

import (
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/parser"
)

type ExpressionVisitor interface {
	visitBinaryExpr(expr parser.Binary)
	visitGroupExpr(expr parser.Grouping)
	visitLiteralExpr(expr parser.Literal)
	visitUnaryExpr(expr parser.Unary)
}
