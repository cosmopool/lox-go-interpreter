package visitor

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/core"
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/environment"
)

type Interpreter struct {
	environment environment.Environment
}

func CreateInterpreter() Interpreter {
	return Interpreter{environment: environment.CreateEnvironment()}
}

func (i *Interpreter) Interpret(expr core.Statement) (any, core.Error) {
	return expr.Accept(i)
}

func (i Interpreter) VisitExpressionStmt(stmt core.ExpressionStmt) (any, core.Error) {
	evaluator := CreateEvaluatorWithEnvironment(&i.environment)
	return evaluator.Evaluate(stmt.Expr)
}

func (i Interpreter) VisitPrintStmt(stmt core.PrintStmt) (any, core.Error) {
	evaluator := CreateEvaluatorWithEnvironment(&i.environment)
	value, err := evaluator.Evaluate(stmt.Expr)
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
	evaluator := CreateEvaluatorWithEnvironment(&i.environment)
	var value any
	var err core.Error

	if stmt.Initializer != nil {
		value, err = evaluator.Evaluate(stmt.Initializer)
		if err.Err != nil {
			return nil, err
		}
	}

	i.environment.AddVariable(stmt.Name.Lexeme, value)

	return nil, err
}

func (i *Interpreter) VisitBlockStmt(stmt core.BlockStmt) (any, core.Error) {
	previousEnvironment := i.environment
	i.environment = environment.CreateEnvironmentWithEnclosing(&previousEnvironment)

	for _, statement := range stmt.Statements {
		_, err := statement.Accept(i)
		if err.Err != nil {
			return nil, err
		}
	}

	i.environment = previousEnvironment
	return nil, core.Error{}
}
