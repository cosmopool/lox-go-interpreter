package parser

type ExpressionVisitor[T any] interface {
	VisitBinaryExpr(expr Binary[T]) (T, error)
	VisitGroupExpr(expr Grouping[T]) (T, error)
	VisitLiteralExpr(expr Literal[T]) (T, error)
	VisitUnaryExpr(expr Unary[T]) (T, error)
}
