package visitor

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/core"
)

type StringifyVisitor struct{}

func (p StringifyVisitor) VisitBinaryExpr(expr core.Binary) (any, error) {
	right, err := expr.Right.Accept(p)
	if err != nil {
		return nil, err
	}

	left, err := expr.Left.Accept(p)
	if err != nil {
		return nil, err
	}

	str := fmt.Sprintf("(%s %s %s)", expr.Operator.Lexeme, left, right)
	return str, nil
}

func (p StringifyVisitor) VisitGroupExpr(expr core.Grouping) (any, error) {
	group, err := expr.Expr.Accept(p)
	if err != nil {
		return nil, err
	}

	str := fmt.Sprintf("(group %s)", group)
	return str, nil
}

func (p StringifyVisitor) VisitLiteralExpr(expr core.Literal) (any, error) {
	if expr.Value == nil {
		return "nil", nil
	}

	_, isFloat := expr.Value.(float64)
	if !isFloat {
		return fmt.Sprint(expr.Value), nil
	}

	separated := strings.Split(fmt.Sprint(expr.Value), ".")
	if len(separated) == 1 {
		return fmt.Sprintf("%.1f", expr.Value), nil
	}

	decimalPart := separated[len(separated)-1]
	decimalPart = strings.ReplaceAll(decimalPart, "0", "")

	if decimalPart == "" {
		return fmt.Sprintf("%.1f", expr.Value), nil
	}

	return fmt.Sprintf("%g", expr.Value), nil
}

func (p StringifyVisitor) VisitUnaryExpr(expr core.Unary) (any, error) {
	right, err := expr.Right.Accept(p)
	if err != nil {
		return nil, err
	}

	str := fmt.Sprintf("(%s %v)", expr.Operator.Lexeme, right)
	return str, nil
}
