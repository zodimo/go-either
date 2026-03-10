package either

func MapLeft[L, R, L2 any](e Either[L, R], f func(L) L2) Either[L2, R] {
	if e.IsLeft() {
		return Left[L2, R](f(e.left))
	}
	return Right[L2, R](e.right)
}
func MapRight[L, R, R2 any](e Either[L, R], f func(R) R2) Either[L, R2] {
	if e.IsRight() {
		return Right[L, R2](f(e.right))
	}
	return Left[L, R2](e.left)
}

func FlatMapLeft[L, R, L2 any](e Either[L, R], f func(L) Either[L2, R]) Either[L2, R] {
	if e.IsLeft() {
		return f(e.left)
	}
	return Right[L2, R](e.right)
}
func FlatMapRight[L, R, R2 any](e Either[L, R], f func(R) Either[L, R2]) Either[L, R2] {
	if e.IsRight() {
		return f(e.right)
	}
	return Left[L, R2](e.left)
}

func Match[L, R, T any](e Either[L, R], f func(L) T, g func(R) T) T {
	if e.IsLeft() {
		return f(e.left)
	}
	return g(e.right)
}
