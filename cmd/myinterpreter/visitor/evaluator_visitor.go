package visitor

import (
	"fmt"
	"strconv"

	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/core"
)

type EvaluatorVisitor struct{}

func (p EvaluatorVisitor) VisitBinaryExpr(expr core.Binary) (any, error) {
	rightExpr, err := expr.Right.Accept(p)
	if err != nil {
		return nil, err
	}
	right, err := getFloat(rightExpr)
	if err != nil {
		return nil, err
	}
	leftExpr, err := expr.Left.Accept(p)
	if err != nil {
		return nil, err
	}
	left, err := getFloat(leftExpr)
	if err != nil {
		return nil, err
	}

	switch expr.Operator.Type {
	case core.MINUS:
		return left - right, nil

	case core.STAR:
		return left * right, nil

	case core.SLASH:
		return left / right, nil

	case core.BANG:
		return !isTruthy(right), nil

	default:
		return nil, nil
	}
}

func (p EvaluatorVisitor) VisitGroupExpr(expr core.Grouping) (any, error) {
	value, err := expr.Expr.Accept(p)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (p EvaluatorVisitor) VisitLiteralExpr(expr core.Literal) (any, error) {
	return expr.Value, nil
}

func (p EvaluatorVisitor) VisitUnaryExpr(expr core.Unary) (any, error) {
	right, err := expr.Right.Accept(p)
	if err != nil {
		return nil, err
	}

	switch expr.Operator.Type {
	case core.MINUS:
		float, err := getFloat(right)
		if err != nil {
			return nil, err
		}
		return -float, nil

	case core.BANG:
		return !isTruthy(right), nil

	default:
		return nil, nil
	}
}

func isTruthy(value any) bool {
	if value == nil {
		return false
	}

	if boolean, ok := value.(bool); ok {
		return boolean
	}

	return true
}

func isEqual(a any, b any) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil {
		return true
	}

	return a == b
}

func getFloat(unk any) (float64, error) {
	switch i := unk.(type) {
	case float64:
		return i, nil
	case float32:
		return float64(i), nil
	case int64:
		return float64(i), nil
	case int32:
		return float64(i), nil
	case int:
		return float64(i), nil
	case uint64:
		return float64(i), nil
	case uint32:
		return float64(i), nil
	case uint:
		return float64(i), nil
	case string:
		return strconv.ParseFloat(i, 64)
	}

	return 0, fmt.Errorf("Could not parse given literal to float %v", unk)
}
