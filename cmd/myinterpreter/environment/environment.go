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
