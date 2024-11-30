package core


type ExpressionVisitor interface {
	VisitBinaryExpr(expr Binary) error
	VisitGroupExpr(expr Grouping) error
	VisitLiteralExpr(expr Literal) error
	VisitUnaryExpr(expr Unary) error
}
