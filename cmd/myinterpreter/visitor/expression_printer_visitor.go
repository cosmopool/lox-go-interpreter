package visitor

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/core"
)

type PrinterVisitor struct {
	stringifyVisitor StringifyVisitor
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
