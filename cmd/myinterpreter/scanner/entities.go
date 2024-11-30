package scanner

type tokenType = string
type keyword = string

const (
	LEFT_PAREN    tokenType = "LEFT_PAREN"
	RIGHT_PAREN   tokenType = "RIGHT_PAREN"
	LEFT_BRACE    tokenType = "LEFT_BRACE"
	RIGHT_BRACE   tokenType = "RIGHT_BRACE"
	COMMA         tokenType = "COMMA"
	DOT           tokenType = "DOT"
	MINUS         tokenType = "MINUS"
	PLUS          tokenType = "PLUS"
	SEMICOLON     tokenType = "SEMICOLON"
	SLASH         tokenType = "SLASH"
	STAR          tokenType = "STAR"
	BANG          tokenType = "BANG"
	BANG_EQUAL    tokenType = "BANG_EQUAL"
	EQUAL         tokenType = "EQUAL"
	EQUAL_EQUAL   tokenType = "EQUAL_EQUAL"
	GREATER       tokenType = "GREATER"
	GREATER_EQUAL tokenType = "GREATER_EQUAL"
	LESS          tokenType = "LESS"
	LESS_EQUAL    tokenType = "LESS_EQUAL"
	EOF           tokenType = "EOF"
	STRING        tokenType = "STRING"
	NUMBER        tokenType = "NUMBER"
	IDENTIFIER    tokenType = "IDENTIFIER"
	AND           keyword   = "AND"
	CLASS         keyword   = "CLASS"
	ELSE          keyword   = "ELSE"
	FALSE         keyword   = "FALSE"
	FOR           keyword   = "FOR"
	FUN           keyword   = "FUN"
	IF            keyword   = "IF"
	NIL           keyword   = "NIL"
	OR            keyword   = "OR"
	PRINT         keyword   = "PRINT"
	RETURN        keyword   = "RETURN"
	SUPER         keyword   = "SUPER"
	THIS          keyword   = "THIS"
	TRUE          keyword   = "TRUE"
	VAR           keyword   = "VAR"
	WHILE         keyword   = "WHILE"
)

func Keywords() map[string]tokenType {
	return map[string]tokenType{
		"and":    AND,
		"class":  CLASS,
		"else":   ELSE,
		"false":  FALSE,
		"for":    FOR,
		"fun":    FUN,
		"if":     IF,
		"nil":    NIL,
		"or":     OR,
		"print":  PRINT,
		"return": RETURN,
		"super":  SUPER,
		"this":   THIS,
		"true":   TRUE,
		"var":    VAR,
		"while":  WHILE,
	}
}

type Token struct {
	Type    tokenType
	Lexeme  string
	Literal any
  Line int
}

type Error struct {
	Line int
	Err  error
}
