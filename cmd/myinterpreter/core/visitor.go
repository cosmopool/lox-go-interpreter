package core


type ExpressionVisitor interface {
	VisitBinaryExpr(expr Binary) (any, error)
	VisitGroupExpr(expr Grouping) (any, error)
	VisitLiteralExpr(expr Literal) (any, error)
	VisitUnaryExpr(expr Unary) (any, error)
}
