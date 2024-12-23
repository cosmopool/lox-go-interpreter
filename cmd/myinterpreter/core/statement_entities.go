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

type VarStmt struct {
	Name Token
	Initializer Expression
}

func (s VarStmt) Accept(visitor StatementVisitor) (any, Error) {
	return visitor.VisitVarStmt(s)
}

type BlockStmt struct {
	Statements []Statement
}

func (s BlockStmt) Accept(visitor StatementVisitor) (any, Error) {
	return visitor.VisitBlockStmt(s)
}
