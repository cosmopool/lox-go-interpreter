package visitor

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/core"
)

type EvaluatorVisitor struct{}

func (p EvaluatorVisitor) VisitBinaryExpr(expr core.Binary) (any, core.Error) {
	rightExpr, err := expr.Right.Accept(p)
	if err.Err != nil {
		return nil, err
	}
	leftExpr, err := expr.Left.Accept(p)
	if err.Err != nil {
		return nil, err
	}

	switch expr.Operator.Type {
	case core.MINUS:
		left, right, err := getMultipleFloat(leftExpr, rightExpr)
		if err != nil {
			return nil, core.Error{Line: expr.Operator.Line, Err: err}
		}
		return left - right, core.Error{}

	case core.STAR:
		left, right, err := getMultipleFloat(leftExpr, rightExpr)
		if err != nil {
			return nil, core.Error{Line: expr.Operator.Line, Err: err}
		}
		return left * right, core.Error{}

	case core.SLASH:
		left, right, err := getMultipleFloat(leftExpr, rightExpr)
		if err != nil {
			return nil, core.Error{Line: expr.Operator.Line, Err: err}
		}
		return left / right, core.Error{}

	case core.PLUS:
		leftStr, leftIsString := leftExpr.(string)
		rightStr, rightIsString := rightExpr.(string)
		if leftIsString && rightIsString {
			return fmt.Sprintf("%s%s", leftStr, rightStr), core.Error{}
		}

		left, right, err := getMultipleFloat(leftExpr, rightExpr)
		if err != nil {
			return nil, core.Error{Line: expr.Operator.Line, Err: fmt.Errorf("Operands must be two numbers or two strings.")}
		}
		return left + right, core.Error{}

	case core.GREATER:
		left, right, err := getMultipleFloat(leftExpr, rightExpr)
		if err != nil {
			return nil, core.Error{Line: expr.Operator.Line, Err: err}
		}
		return left > right, core.Error{}

	case core.GREATER_EQUAL:
		left, right, err := getMultipleFloat(leftExpr, rightExpr)
		if err != nil {
			return nil, core.Error{Line: expr.Operator.Line, Err: err}
		}
		return left >= right, core.Error{}

	case core.LESS:
		left, right, err := getMultipleFloat(leftExpr, rightExpr)
		if err != nil {
			return nil, core.Error{Line: expr.Operator.Line, Err: err}
		}
		return left < right, core.Error{}

	case core.LESS_EQUAL:
		left, right, err := getMultipleFloat(leftExpr, rightExpr)
		if err != nil {
			return nil, core.Error{Line: expr.Operator.Line, Err: err}
		}
		return left <= right, core.Error{}

	case core.EQUAL_EQUAL:
		return isEqual(leftExpr, rightExpr), core.Error{}

	case core.BANG_EQUAL:
		return !isEqual(leftExpr, rightExpr), core.Error{}

	default:
		return nil, core.Error{}
	}
}

func (p EvaluatorVisitor) VisitGroupExpr(expr core.Grouping) (any, core.Error) {
	value, err := expr.Expr.Accept(p)
	if err.Err != nil {
		return nil, err
	}

	return value, core.Error{}
}

func (p EvaluatorVisitor) VisitLiteralExpr(expr core.Literal) (any, core.Error) {
	return expr.Value, core.Error{}
}

func (p EvaluatorVisitor) VisitUnaryExpr(expr core.Unary) (any, core.Error) {
	right, err := expr.Right.Accept(p)
	if err.Err != nil {
		return nil, err
	}

	switch expr.Operator.Type {
	case core.MINUS:
		float, err := getFloat(right)
		if err != nil {
			return nil, core.Error{Line: expr.Operator.Line, Err: err}
		}
		return -float, core.Error{}

	case core.BANG:
		return !isTruthy(right), core.Error{}

	default:
		return nil, core.Error{}
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
	}

	return 0, fmt.Errorf("Operand must be a number.")
}

func getMultipleFloat(a any, b any) (float64, float64, error) {
	aFloat, err := getFloat(a)
	if err != nil {
		return 0, 0, err
	}
	bFloat, err := getFloat(b)
	if err != nil {
		return 0, 0, err
	}

	return aFloat, bFloat, nil
}
