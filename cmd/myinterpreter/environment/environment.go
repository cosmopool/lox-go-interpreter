package environment

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/core"
)

type Environment struct {
	variables map[string]any
	enclosing *Environment
}

func CreateEnvironment() Environment {
	return Environment{enclosing: nil, variables: map[string]any{}}
}

func CreateEnvironmentWithEnclosing(environment *Environment) Environment {
	return Environment{enclosing: environment, variables: map[string]any{}}
}

func (e *Environment) GetVariable(token *core.Token) (any, core.Error) {
	value, ok := e.variables[token.Lexeme]
	if ok {
		return value, core.Error{}
	}

	if e.enclosing != nil {
		return e.enclosing.GetVariable(token)
	}

	return nil, core.Error{Line: token.Line, Err: fmt.Errorf("Undefined variable '" + token.Lexeme + "'."), ExitCode: 70}
}

func (e *Environment) AddVariable(name string, value any) {
	if e.variables == nil {
		e.variables = map[string]any{}
	}
	e.variables[name] = value
}

func (e *Environment) AssignVariable(token *core.Token, value any) *core.Error {
	if _, ok := e.variables[token.Lexeme]; ok {
		e.variables[token.Lexeme] = value
		return nil
	}

	if e.enclosing != nil {
		return e.enclosing.AssignVariable(token, value)
	}

	return &core.Error{Line: token.Line, Err: fmt.Errorf("Undefined variable '" + token.Lexeme + "'."), ExitCode: 70}
}

func (e Environment) String() string {
  return fmt.Sprintf("enclosing: %v, variables: %v\n", e.enclosing != nil, e.variables)
}
