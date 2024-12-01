package core


type ExpressionVisitor interface {
	VisitBinaryExpr(expr Binary) (any, Error)
	VisitGroupExpr(expr Grouping) (any, Error)
	VisitLiteralExpr(expr Literal) (any, Error)
	VisitUnaryExpr(expr Unary) (any, Error)
}
