package core

type ExpressionVisitor interface {
	VisitBinaryExpr(expr Binary) (any, Error)
	VisitGroupExpr(expr Grouping) (any, Error)
	VisitLiteralExpr(expr Literal) (any, Error)
	VisitUnaryExpr(expr Unary) (any, Error)
	VisitVariableExpr(expr Variable) (any, Error)
}

type StatementVisitor interface {
	VisitExpressionStmt(stmt ExpressionStmt) (any, Error)
	VisitPrintStmt(stmt PrintStmt) (any, Error)
	VisitVarStmt(stmt VarStmt) (any, Error)
}
