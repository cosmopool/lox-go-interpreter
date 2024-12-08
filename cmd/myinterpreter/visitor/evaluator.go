package visitor

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/core"
)

type Evaluator struct{}

func (i Evaluator) Evaluate(expr core.Expression) (any, core.Error) {
	return expr.Accept(i)
}

func (i Evaluator) VisitBinaryExpr(expr core.Binary) (any, core.Error) {
	rightExpr, err := expr.Right.Accept(i)
	if err.Err != nil {
		return nil, err
	}
	leftExpr, err := expr.Left.Accept(i)
	if err.Err != nil {
		return nil, err
	}

	switch expr.Operator.Type {
	case core.MINUS:
		left, right, err := getMultipleFloat(expr.Operator, leftExpr, rightExpr)
		if err != nil {
			return nil, core.Error{Line: expr.Operator.Line, Err: fmt.Errorf("Operands must be numbers.")}
		}
		return left - right, core.Error{}

	case core.STAR:
		left, right, err := getMultipleFloat(expr.Operator, leftExpr, rightExpr)
		if err != nil {
			return nil, core.Error{Line: expr.Operator.Line, Err: fmt.Errorf("Operands must be numbers.")}
		}
		return left * right, core.Error{}

	case core.SLASH:
		left, right, err := getMultipleFloat(expr.Operator, leftExpr, rightExpr)
		if err != nil {
			return nil, core.Error{Line: expr.Operator.Line, Err: fmt.Errorf("Operands must be numbers.")}
		}
		return left / right, core.Error{}

	case core.PLUS:
		leftStr, leftIsString := leftExpr.(string)
		rightStr, rightIsString := rightExpr.(string)
		if leftIsString && rightIsString {
			return fmt.Sprintf("%s%s", leftStr, rightStr), core.Error{}
		}

		left, right, err := getMultipleFloat(expr.Operator, leftExpr, rightExpr)
		if err != nil {
			return nil, core.Error{Line: expr.Operator.Line, Err: fmt.Errorf("Operands must be two numbers or two strings.")}
		}
		return left + right, core.Error{}

	case core.GREATER:
		left, right, err := getMultipleFloat(expr.Operator, leftExpr, rightExpr)
		if err != nil {
			return nil, core.Error{Line: expr.Operator.Line, Err: fmt.Errorf("Operands must be numbers.")}
		}
		return left > right, core.Error{}

	case core.GREATER_EQUAL:
		left, right, err := getMultipleFloat(expr.Operator, leftExpr, rightExpr)
		if err != nil {
			return nil, core.Error{Line: expr.Operator.Line, Err: fmt.Errorf("Operands must be numbers.")}
		}
		return left >= right, core.Error{}

	case core.LESS:
		left, right, err := getMultipleFloat(expr.Operator, leftExpr, rightExpr)
		if err != nil {
			return nil, core.Error{Line: expr.Operator.Line, Err: fmt.Errorf("Operands must be numbers.")}
		}
		return left < right, core.Error{}

	case core.LESS_EQUAL:
		left, right, err := getMultipleFloat(expr.Operator, leftExpr, rightExpr)
		if err != nil {
			return nil, core.Error{Line: expr.Operator.Line, Err: fmt.Errorf("Operands must be numbers.")}
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

func (i Evaluator) VisitGroupExpr(expr core.Grouping) (any, core.Error) {
	value, err := expr.Expr.Accept(i)
	if err.Err != nil {
		return nil, err
	}

	return value, core.Error{}
}

func (i Evaluator) VisitLiteralExpr(expr core.Literal) (any, core.Error) {
	return expr.Value, core.Error{}
}

func (i Evaluator) VisitUnaryExpr(expr core.Unary) (any, core.Error) {
	right, err := expr.Right.Accept(i)
	if err.Err != nil {
		return nil, err
	}

	switch expr.Operator.Type {
	case core.MINUS:
		float, err := getFloat(expr.Operator, right)
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

func getFloat(operator core.Token, operand any) (float64, error) {
	switch i := operand.(type) {
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

	return 0, fmt.Errorf("%v Operand must be a number.", operator)
}

func getMultipleFloat(operator core.Token, a any, b any) (float64, float64, error) {
	aFloat, err := getFloat(operator, a)
	if err != nil {
		return 0, 0, err
	}
	bFloat, err := getFloat(operator, b)
	if err != nil {
		return 0, 0, err
	}

	return aFloat, bFloat, nil
}
