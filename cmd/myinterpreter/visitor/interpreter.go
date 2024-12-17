package visitor

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/core"
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/environment"
)

type Interpreter struct{}

func (i Interpreter) Interpret(expr core.Statement) (any, core.Error) {
	return expr.Accept(i)
}

func (i Interpreter) VisitExpressionStmt(stmt core.ExpressionStmt) (any, core.Error) {
	return Evaluator{}.Evaluate(stmt.Expr)
}

func (i Interpreter) VisitPrintStmt(stmt core.PrintStmt) (any, core.Error) {
	value, err := Evaluator{}.Evaluate(stmt.Expr)
	if err.Err != nil {
		return nil, err
	}

	if value == nil {
		fmt.Println("nil")
		return nil, core.Error{}
	}

	_, isFloat := value.(float64)
	if !isFloat {
		fmt.Println(value)
		return nil, core.Error{}
	}

	separated := strings.Split(fmt.Sprint(value), ".")
	if len(separated) == 1 {
		fmt.Println(value)
		return nil, core.Error{}
	}

	decimalPart := separated[len(separated)-1]
	decimalPart = strings.ReplaceAll(decimalPart, "0", "")

	if decimalPart == "" {
		fmt.Printf("%.1f", value)
		return nil, core.Error{}
	}

	return nil, core.Error{}
}

func (i Interpreter) VisitVarStmt(stmt core.VarStmt) (any, core.Error) {
	var value any
	var err core.Error

	if stmt.Initializer != nil {
		value, err = Evaluator{}.Evaluate(stmt.Initializer)
		if err.Err != nil {
			return nil, err
		}
	}

  environment.AddVariable(stmt.Name.Lexeme, value)
	return nil, err
}
