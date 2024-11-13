package visitor

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/parser"
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/utils"
)

type PrinterVisitor[T string] struct{}

func (p PrinterVisitor[string]) Print(expr parser.Expression[string]) (string, error) {
	return expr.Accept(p)
}

func (p PrinterVisitor[string]) VisitBinaryExpr(expr parser.Binary[string]) (string, error) {
	left, err := expr.Left.Accept(p)
	if err != nil {
		return "(error)", err
	}

	right, err := expr.Right.Accept(p)
	if err != nil {
		return "(error)", err
	}

	return string(fmt.Sprintf("(%s %s %s)", expr.Operator.Lexeme, left, right)), nil
}

func (p PrinterVisitor[string]) VisitGroupExpr(expr parser.Grouping[string]) (string, error) {
	subExpr, err := expr.Expr.Accept(p)
	if err != nil {
		return "(error)", err
	}

	return string(fmt.Sprintf("(group %s)", subExpr)), nil
}

func (p PrinterVisitor[string]) VisitLiteralExpr(expr parser.Literal[string]) (string, error) {
	val, err := expr.Accept(p)
	if err != nil {
		return "(error)", err
	}
	return string(val), nil
}

func (p PrinterVisitor[string]) VisitUnaryExpr(expr parser.Unary[string]) (string, error) {
	right, err := expr.Accept(p)
	if err != nil {
		return "(error)", err
	}

	return string(fmt.Sprintf("(%s %v)", expr.Operator.Lexeme, right)), nil
}
