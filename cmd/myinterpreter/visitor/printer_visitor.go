package visitor

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/core"
)

type PrinterVisitor struct {
	stringifyVisitor StringifyVisitor
}

func (p PrinterVisitor) Print(expr core.Statement) (any, core.Error) {
	return expr.Accept(p)
}

func (p PrinterVisitor) PrintExpression(expr core.Expression) (any, core.Error) {
	return expr.Accept(p)
}

func (p PrinterVisitor) VisitExpressionStmt(stmt core.ExpressionStmt) (any, core.Error) {
	str, err := p.stringifyVisitor.VisitExpressionStmt(stmt)
	if err.Err != nil {
		return nil, err
	}

	fmt.Println(str)
	return str, core.Error{}
}

func (p PrinterVisitor) VisitVarStmt(stmt core.VarStmt) (any, core.Error) {
	str, err := p.stringifyVisitor.VisitVarStmt(stmt)
	if err.Err != nil {
		return nil, err
	}

	fmt.Println(str)
	return str, core.Error{}
}

func (p PrinterVisitor) VisitPrintStmt(stmt core.PrintStmt) (any, core.Error) {
	str, err := p.stringifyVisitor.VisitPrintStmt(stmt)
	if err.Err != nil {
		return nil, err
	}

	fmt.Println(str)
	return str, core.Error{}
}

func (p PrinterVisitor) VisitBinaryExpr(expr core.Binary) (any, core.Error) {
	str, err := p.stringifyVisitor.VisitBinaryExpr(expr)
	if err.Err != nil {
		return nil, err
	}

	fmt.Println(str)
	return str, core.Error{}
}

func (p PrinterVisitor) VisitGroupExpr(expr core.Grouping) (any, core.Error) {
	str, err := p.stringifyVisitor.VisitGroupExpr(expr)
	if err.Err != nil {
		return nil, err
	}

	fmt.Println(str)
	return str, core.Error{}
}

func (p PrinterVisitor) VisitLiteralExpr(expr core.Literal) (any, core.Error) {
	str, err := p.stringifyVisitor.VisitLiteralExpr(expr)
	if err.Err != nil {
		return nil, err
	}

	fmt.Println(str)
	return str, core.Error{}
}

func (p PrinterVisitor) VisitUnaryExpr(expr core.Unary) (any, core.Error) {
	str, err := p.stringifyVisitor.VisitUnaryExpr(expr)
	if err.Err != nil {
		return nil, err
	}

	fmt.Println(str)
	return str, core.Error{}
}

func (p PrinterVisitor) VisitVariableExpr(expr core.Variable) (any, core.Error) {
	str, err := p.stringifyVisitor.VisitVariableExpr(expr)
	if err.Err != nil {
		return nil, err
	}

	fmt.Println(str)
	return str, core.Error{}
}

func (p PrinterVisitor) VisitAssignExpr(expr core.Assign) (any, core.Error) {
	str, err := p.stringifyVisitor.VisitAssignExpr(expr)
	if err.Err != nil {
		return nil, err
	}

	fmt.Println(str)
	return str, core.Error{}
}

// VisitBlockStmt implements core.StatementVisitor.
func (p PrinterVisitor) VisitBlockStmt(stmt core.BlockStmt) (any, core.Error) {
	panic("unimplemented")
}
