package core

type Expression interface {
	Accept(visitor ExpressionVisitor) (any, Error)
}

type Binary struct {
	Left     Expression
	Operator Token
	Right    Expression
}

func (b Binary) Accept(visitor ExpressionVisitor) (any, Error) {
	return visitor.VisitBinaryExpr(b)
}

type Grouping struct {
	Expr Expression
}

func (g Grouping) Accept(visitor ExpressionVisitor) (any, Error) {
	return visitor.VisitGroupExpr(g)
}

type Literal struct {
	Value any
}

func (l Literal) Accept(visitor ExpressionVisitor) (any, Error) {
	return visitor.VisitLiteralExpr(l)
}

type Unary struct {
	Operator Token
	Right    Expression
}

func (u Unary) Accept(visitor ExpressionVisitor) (any, Error) {
	return visitor.VisitUnaryExpr(u)
}

type Error struct {
	Line int
	Err  error
  ExitCode int
}
