package core

type Expression interface {
	Accept(visitor ExpressionVisitor) error
}

type Binary struct {
	Left     Expression
	Operator Token
	Right    Expression
}

func (b Binary) Accept(visitor ExpressionVisitor) error {
	return visitor.VisitBinaryExpr(b)
}

type Grouping struct {
	Expr Expression
}

func (g Grouping) Accept(visitor ExpressionVisitor) error {
	return visitor.VisitGroupExpr(g)
}

type Literal struct {
	Value any
}

func (l Literal) Accept(visitor ExpressionVisitor) error {
	return visitor.VisitLiteralExpr(l)
}

type Unary struct {
	Operator Token
	Right    Expression
}

func (u Unary) Accept(visitor ExpressionVisitor) error {
	return visitor.VisitUnaryExpr(u)
}

type Error struct {
	Line int
	Err  error
}