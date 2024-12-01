package visitor

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/core"
)

type EvaluatorVisitor struct{}

func (p EvaluatorVisitor) VisitBinaryExpr(expr core.Binary) (any, error) {
	return nil, nil
}

func (p EvaluatorVisitor) VisitGroupExpr(expr core.Grouping) (any, error) {
	return nil, nil
}

func (p EvaluatorVisitor) VisitLiteralExpr(expr core.Literal) (any, error) {
	if expr.Value == nil {

		fmt.Println("nil")
		return nil, nil
	}

	fmt.Println(expr.Value)
	return nil, nil
}

func (p EvaluatorVisitor) VisitUnaryExpr(expr core.Unary) (any, error) {
	return nil, nil
}
