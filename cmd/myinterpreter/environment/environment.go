package environment

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/core"
)

var variables = map[string]any{}

func GetVariable(token core.Token) (any, core.Error) {
	value, ok := variables[token.Lexeme]
	if !ok {
		return nil, core.Error{Line: token.Line, Err: fmt.Errorf("Undefined variable '" + token.Lexeme + "'."), ExitCode: 70}
	}

	return value, core.Error{}
}

func AddVariable(name string, value any) {
	variables[name] = value
}

func AssignVariable(name string, value any, line int) *core.Error {
	if _, ok := variables[name]; ok {
		variables[name] = value
		return nil
	}

	return &core.Error{Line: line, Err: fmt.Errorf("Undefined variable '" + name + "'."), ExitCode: 70}
}
