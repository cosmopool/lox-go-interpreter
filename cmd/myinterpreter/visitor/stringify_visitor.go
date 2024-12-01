package visitor

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/core"
)

type StringifyVisitor struct{}

func (p StringifyVisitor) VisitBinaryExpr(expr core.Binary) (any, core.Error) {
	right, err := expr.Right.Accept(p)
	if err.Err != nil {
		return nil, err
	}

	left, err := expr.Left.Accept(p)
	if err.Err != nil {
		return nil, err
	}

	str := fmt.Sprintf("(%s %s %s)", expr.Operator.Lexeme, left, right)
	return str, core.Error{}
}

func (p StringifyVisitor) VisitGroupExpr(expr core.Grouping) (any, core.Error) {
	group, err := expr.Expr.Accept(p)
	if err.Err != nil {
		return nil, err
	}

	str := fmt.Sprintf("(group %s)", group)
	return str, core.Error{}
}

func (p StringifyVisitor) VisitLiteralExpr(expr core.Literal) (any, core.Error) {
	if expr.Value == nil {
		return "nil", core.Error{}
	}

	_, isFloat := expr.Value.(float64)
	if !isFloat {
		return fmt.Sprint(expr.Value), core.Error{}
	}

	separated := strings.Split(fmt.Sprint(expr.Value), ".")
	if len(separated) == 1 {
		return fmt.Sprintf("%.1f", expr.Value), core.Error{}
	}

	decimalPart := separated[len(separated)-1]
	decimalPart = strings.ReplaceAll(decimalPart, "0", "")

	if decimalPart == "" {
		return fmt.Sprintf("%.1f", expr.Value), core.Error{}
	}

	return fmt.Sprintf("%g", expr.Value), core.Error{}
}

func (p StringifyVisitor) VisitUnaryExpr(expr core.Unary) (any, core.Error) {
	right, err := expr.Right.Accept(p)
	if err.Err != nil {
		return nil, err
	}

	str := fmt.Sprintf("(%s %v)", expr.Operator.Lexeme, right)
	return str, core.Error{}
}
