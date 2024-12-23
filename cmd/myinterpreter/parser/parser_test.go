package parser

// // isAtEnd() should return FALSE if NOT reached EOF token
// func TestIsAtEndReturnsFalseIfNotEOFToken(t *testing.T) {
// 	tokens := []core.Token{
// 		{Type: core.NUMBER, Lexeme: "2", Literal: 2},
// 	}
//
// 	parser := Parser{Tokens: tokens}
// 	if parser.isAtEnd() == true {
// 		t.Fatal("isAtEnd() should return FALSE if NOT reached EOF token")
// 	}
// }
//
// // isAtEnd() should return TRUE if reached EOF token
// func TestIsAtEndReturnsTrueReachedEOFToken(t *testing.T) {
// 	tokens := []core.Token{
// 		{Type: core.EOF, Lexeme: "EOF", Literal: nil},
// 	}
//
// 	parser := Parser{Tokens: tokens}
// 	if parser.isAtEnd() == false {
// 		t.Fatal("isAtEnd() should return if reached EOF token")
// 	}
// }
//
// // advance() should increase position if not reached the end tokens[]
// func TestAdvanceShouldIncreasePositionWhenNotOnEnd(t *testing.T) {
// 	tokens := []core.Token{
// 		{Type: core.NUMBER, Lexeme: "2", Literal: 2},
// 		{Type: core.EOF, Lexeme: "EOF", Literal: nil},
// 	}
//
// 	parser := Parser{Tokens: tokens}
// 	if parser.position != 0 {
// 		t.Fatalf("position shoul be 0, but found: %v", parser.position)
// 	}
//
// 	token := parser.advance()
// 	if parser.position != 1 {
// 		t.Fatalf("position shoul be 1, but found: %v", parser.position)
// 	}
// 	if !reflect.DeepEqual(token, tokens[0]) {
// 		t.Fatalf("expecting first advance to return: %v, but found: %v", tokens[0], token)
// 	}
//
// 	token = parser.advance()
// 	if !reflect.DeepEqual(token, tokens[1]) {
// 		t.Fatalf("expecting second advance to return: %v, but found: %v", tokens[1], token)
// 	}
// 	if parser.position != 1 {
// 		t.Fatalf("position shoul be 1, but found: %v", parser.position)
// 	}
// }
//
// func TestCallingTermShouldIncreasePosition(t *testing.T) {
// 	tokens := []core.Token{
// 		{Type: core.NUMBER, Lexeme: "2", Literal: 2},
// 		{Type: core.EQUAL, Lexeme: "=", Literal: nil},
// 		{Type: core.NUMBER, Lexeme: "3", Literal: 3},
// 	}
//
// 	parser := Parser{Tokens: tokens}
// 	parser.term()
// 	if parser.position == 0 {
// 		t.Fatal("term() should increase position")
// 	}
// }
//
// func TestPrimaryIterateOverTokens(t *testing.T) {
// 	tokens := []core.Token{
// 		{Type: core.LEFT_PAREN, Lexeme: "(", Literal: nil},
// 		{Type: core.STRING, Lexeme: "foo", Literal: "foo"},
// 		{Type: core.RIGHT_PAREN, Lexeme: ")", Literal: nil},
// 	}
//
// 	parser := Parser{Tokens: tokens}
// 	expr, err := parser.primary()
// 	if err != nil {
// 		t.Fatalf("was not expecting any errors, but got: %v", err)
// 	}
// 	_, isGroup := expr.(Grouping)
// 	if !isGroup {
// 		t.Fatalf("should receive a group, but got: %v", expr)
// 	}
// }
//
// func TestGroupingEmptyParentheses(t *testing.T) {
// 	tokens := []core.Token{
// 		{Type: core.LEFT_PAREN, Lexeme: "(", Literal: nil},
// 		{Type: core.RIGHT_PAREN, Lexeme: ")", Literal: nil},
// 	}
//
// 	parser := Parser{Tokens: tokens}
// 	_, err := parser.expression()
// 	if err == fmt.Errorf("Empty group") {
// 		t.Fatal("was expecting a empty group error, but didn't get one")
// 	}
// }
//
// func TestFactor(t *testing.T) {
// 	tokens := []core.Token{
// 		{Type: core.NUMBER, Lexeme: "43.0", Literal: 43},
// 		{Type: core.STAR, Lexeme: "*", Literal: nil},
// 		{Type: core.NUMBER, Lexeme: "72.0", Literal: 72},
// 		{Type: core.SLASH, Lexeme: "/", Literal: nil},
// 		{Type: core.NUMBER, Lexeme: "48.0", Literal: 48},
// 		{Type: core.EOF, Lexeme: "EOF", Literal: nil},
// 	}
//
// 	parser := Parser{Tokens: tokens}
// 	_, err := parser.expression()
// 	if err != nil {
// 		t.Fatalf("was not expecting any errors, but got: %v", err)
// 	}
// }
//
// func TestFactorWithParentheses(t *testing.T) {
// 	tokens := []core.Token{
// 		{Type: core.LEFT_PAREN, Lexeme: "(", Literal: nil},
// 		{Type: core.NUMBER, Lexeme: "22.0", Literal: 22},
// 		{Type: core.STAR, Lexeme: "*", Literal: nil},
// 		{Type: core.MINUS, Lexeme: "-", Literal: nil},
// 		{Type: core.NUMBER, Lexeme: "98.0", Literal: 98},
// 		{Type: core.LEFT_PAREN, Lexeme: "(", Literal: nil},
// 		{Type: core.NUMBER, Lexeme: "51.0", Literal: 51},
// 		{Type: core.SLASH, Lexeme: "/", Literal: nil},
// 		{Type: core.NUMBER, Lexeme: "95.0", Literal: 95},
// 		{Type: core.RIGHT_PAREN, Lexeme: ")", Literal: nil},
// 		{Type: core.RIGHT_PAREN, Lexeme: ")", Literal: nil},
// 		{Type: core.EOF, Lexeme: "EOF", Literal: nil},
// 	}
//
// 	parser := Parser{Tokens: tokens}
// 	e, err := parser.expression()
// 	if err != nil {
// 		t.Fatalf("was not expecting any errors, but got: %v", err)
// 	}
//
// 	_, ok := e.(Grouping)
// 	if !ok {
// 		t.Fatalf("was expecting an Group expression, but got: %v", e)
// 	}
// }
//
// func TestParserFactorWithParentheses(t *testing.T) {
// 	tokens := []core.Token{
// 		{Type: core.LEFT_PAREN, Lexeme: "(", Literal: nil},
// 		{Type: core.NUMBER, Lexeme: "22.0", Literal: 22},
// 		{Type: core.STAR, Lexeme: "*", Literal: nil},
// 		{Type: core.MINUS, Lexeme: "-", Literal: nil},
// 		{Type: core.NUMBER, Lexeme: "98.0", Literal: 98},
// 		{Type: core.SLASH, Lexeme: "/", Literal: nil},
// 		{Type: core.LEFT_PAREN, Lexeme: "(", Literal: nil},
// 		{Type: core.NUMBER, Lexeme: "51.0", Literal: 51},
// 		{Type: core.STAR, Lexeme: "*", Literal: nil},
// 		{Type: core.NUMBER, Lexeme: "95.0", Literal: 95},
// 		{Type: core.RIGHT_PAREN, Lexeme: ")", Literal: nil},
// 		{Type: core.RIGHT_PAREN, Lexeme: ")", Literal: nil},
// 		{Type: core.EOF, Lexeme: "EOF", Literal: nil},
// 	}
//
// 	parser := Parser{Tokens: tokens}
// 	e, err := parser.Parse()
// 	if err != nil {
// 		t.Fatalf("was not expecting any errors, but got: %v", err)
// 	}
//
// 	if len(e) != 1 {
// 		t.Fatalf("there should be 1 expression, but got: %d", len(e))
// 	}
//
// 	_, ok := e[0].(Grouping)
// 	if !ok {
// 		t.Fatalf("was expecting an Group expression, but got: %v", e[0])
// 	}
// }
//
// func TestParserComparison(t *testing.T) {
// 	tokens := []core.Token{
// 		{Type: core.LEFT_PAREN, Lexeme: "(", Literal: nil},
// 		{Type: core.NUMBER, Lexeme: "23.0", Literal: 23},
// 		{Type: core.MINUS, Lexeme: "-", Literal: nil},
// 		{Type: core.NUMBER, Lexeme: "98.0", Literal: 98},
// 		{Type: core.RIGHT_PAREN, Lexeme: ")", Literal: nil},
// 		{Type: core.GREATER_EQUAL, Lexeme: ">=", Literal: nil},
// 		{Type: core.MINUS, Lexeme: "-", Literal: nil},
// 		{Type: core.LEFT_PAREN, Lexeme: "(", Literal: nil},
// 		{Type: core.NUMBER, Lexeme: "22.0", Literal: 22},
// 		{Type: core.SLASH, Lexeme: "/", Literal: nil},
// 		{Type: core.NUMBER, Lexeme: "51.0", Literal: 51},
// 		{Type: core.PLUS, Lexeme: "+", Literal: nil},
// 		{Type: core.NUMBER, Lexeme: "95.0", Literal: 95},
// 		{Type: core.RIGHT_PAREN, Lexeme: ")", Literal: nil},
// 		{Type: core.EOF, Lexeme: "EOF", Literal: nil},
// 	}
//
// 	parser := Parser{Tokens: tokens}
// 	e, err := parser.Parse()
// 	if err != nil {
// 		t.Fatalf("was not expecting any errors, but got: %v", err)
// 	}
//
// 	if len(e) != 1 {
// 		t.Fatalf("there should be 1 expression, but got: %d", len(e))
// 	}
//
// 	_, ok := e[0].(Binary)
// 	if !ok {
// 		t.Fatalf("was expecting an Group expression, but got: %v", e[0])
// 	}
// }

// func TestParserMissingExpression(t *testing.T) {
// 	tokens := []core.Token{
// 		{Type: core.LEFT_PAREN, Lexeme: "(", Literal: nil},
// 		{Type: core.NUMBER, Lexeme: "51.0", Literal: 51},
// 		{Type: core.PLUS, Lexeme: "+", Literal: nil},
// 		{Type: core.RIGHT_PAREN, Lexeme: ")", Literal: nil},
// 		{Type: core.EOF, Lexeme: "EOF", Literal: nil},
// 	}
//
// 	e, err := Parse(tokens)
// 	if err != nil {
// 		t.Fatalf("was not expecting any errors, but got: %v", err)
// 	}
//
// 	if len(e) != 1 {
// 		t.Fatalf("there should be 1 expression, but got: %d", len(e))
// 	}
//
// 	_, ok := e[0].(core.Binary)
// 	if !ok {
// 		t.Fatalf("was expecting an Group expression, but got: %v", e[0])
// 	}
// }
