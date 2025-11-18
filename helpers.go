package either

import (
	"github.com/zodimo/go-maybe"
)

func MapLeft[L, R, L2 any](e Either[L, R], f func(L) L2) maybe.Maybe[Either[L2, R]] {
	if e.IsLeft() {
		return maybe.Some(Left[L2, R](f(e.left)))
	}
	return maybe.None[Either[L2, R]]()
}
func MapRight[L, R, R2 any](e Either[L, R], f func(R) R2) maybe.Maybe[Either[L, R2]] {
	if e.IsRight() {
		return maybe.Some(Right[L, R2](f(e.right)))
	}
	return maybe.None[Either[L, R2]]()
}

func FlatMapLeft[L, R, L2 any](e Either[L, R], f func(L) Either[L2, R]) maybe.Maybe[Either[L2, R]] {
	if e.IsLeft() {
		return maybe.Some(f(e.left))
	}
	return maybe.None[Either[L2, R]]()
}
func FlatMapRight[L, R, R2 any](e Either[L, R], f func(R) Either[L, R2]) maybe.Maybe[Either[L, R2]] {
	if e.IsRight() {
		return maybe.Some(f(e.right))
	}
	return maybe.None[Either[L, R2]]()
}
