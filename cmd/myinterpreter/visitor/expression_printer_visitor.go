package visitor

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/core"
)

type PrinterVisitor struct {
	stringifyVisitor StringifyVisitor
}

func (p PrinterVisitor) VisitBinaryExpr(expr core.Binary) (any, error) {
	str, err := p.stringifyVisitor.VisitBinaryExpr(expr)
	if err != nil {
		return nil, err
	}

	fmt.Println(str)
	return str, nil
}

func (p PrinterVisitor) VisitGroupExpr(expr core.Grouping) (any, error) {
	str, err := p.stringifyVisitor.VisitGroupExpr(expr)
	if err != nil {
		return nil, err
	}

	fmt.Println(str)
	return str, nil
}

func (p PrinterVisitor) VisitLiteralExpr(expr core.Literal) (any, error) {
	str, err := p.stringifyVisitor.VisitLiteralExpr(expr)
	if err != nil {
		return nil, err
	}

	fmt.Println(str)
	return str, nil
}

func (p PrinterVisitor) VisitUnaryExpr(expr core.Unary) (any, error) {
	str, err := p.stringifyVisitor.VisitUnaryExpr(expr)
	if err != nil {
		return nil, err
	}

	fmt.Println(str)
	return str, nil
}
