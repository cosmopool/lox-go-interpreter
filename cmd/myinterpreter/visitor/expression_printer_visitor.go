package visitor

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/core"
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/utils"
)

type PrinterVisitor struct{}

func (p PrinterVisitor) VisitBinaryExpr(expr core.Binary) error {
	fmt.Println(fmt.Sprintf("(%s %s %s)", expr.Operator.Lexeme, expr.Left, expr.Right))
	return nil
}

func (p PrinterVisitor) VisitGroupExpr(expr core.Grouping) error {
	fmt.Println(fmt.Sprintf("(group %s)", expr.Expr))
	return nil
}

func (p PrinterVisitor) VisitLiteralExpr(expr core.Literal) error {
  fmt.Println(fmt.Sprintf(utils.VariableToString(expr.Value, false)))
	return nil
}

func (p PrinterVisitor) VisitUnaryExpr(expr core.Unary) error {
	fmt.Println(fmt.Sprintf("(%s %v)", expr.Operator.Lexeme, expr.Right))
	return nil
}
