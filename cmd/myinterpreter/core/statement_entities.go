package core

type Statement interface {
	Accept(visitor StatementVisitor) (any, Error)
}

type ExpressionStmt struct {
	Expr Expression
}

func (s ExpressionStmt) Accept(visitor StatementVisitor) (any, Error) {
	return visitor.VisitExpressionStmt(s)
}

type PrintStmt struct {
	Expr Expression
}

func (s PrintStmt) Accept(visitor StatementVisitor) (any, Error) {
	return visitor.VisitPrintStmt(s)
}
