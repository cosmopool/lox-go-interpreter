package visitor

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/core"
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/utils"
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
	str := utils.VariableToString(expr.Value, false)
	return str, nil
}

func (p StringifyVisitor) VisitUnaryExpr(expr core.Unary) (any, error) {
	right, err := expr.Right.Accept(p)
	if err != nil {
		return nil, err
	}

	str := fmt.Sprintf("(%s %v)", expr.Operator.Lexeme, right)
	return str, nil
}
