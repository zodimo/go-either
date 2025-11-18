package either

import (
	"reflect"

	"github.com/zodimo/go-maybe"
)

type Either[L, R any] struct {
	left    L
	right   R
	lType   reflect.Type
	rType   reflect.Type
	isLeft  bool
	isRight bool
}

func Left[L, R any](left L) Either[L, R] {
	rType := reflect.TypeOf((*R)(nil)).Elem()
	lType := reflect.TypeOf((*L)(nil)).Elem()
	return Either[L, R]{
		left:    left,
		isLeft:  true,
		isRight: false,
		rType:   rType,
		lType:   lType,
	}
}
func Right[L, R any](right R) Either[L, R] {
	rType := reflect.TypeOf((*R)(nil)).Elem()
	lType := reflect.TypeOf((*L)(nil)).Elem()
	return Either[L, R]{
		right:   right,
		rType:   rType,
		lType:   lType,
		isLeft:  false,
		isRight: true,
	}
}

func (e Either[L, R]) IsLeft() bool {
	return e.isLeft
}
func (e Either[L, R]) IsRight() bool {
	return e.isRight
}
func (e Either[L, R]) Left() maybe.Maybe[L] {
	if e.IsLeft() {
		return maybe.Some(e.left)
	}
	return maybe.None[L]()
}
func (e Either[L, R]) Right() maybe.Maybe[R] {
	if e.IsRight() {
		return maybe.Some(e.right)
	}
	return maybe.None[R]()
}

func (e Either[L, R]) MapLeft(f func(L) L) Either[L, R] {
	if e.IsLeft() {
		return Left[L, R](f(e.left))
	}
	return Right[L, R](e.right)
}
func (e Either[L, R]) MapRight(f func(R) R) Either[L, R] {
	if e.IsRight() {
		return Right[L, R](f(e.right))
	}
	return Left[L, R](e.left)
}
func (e Either[L, R]) FlatMapLeft(f func(L) Either[L, R]) Either[L, R] {
	if e.IsLeft() {
		return f(e.left)
	}
	return Right[L, R](e.right)
}
func (e Either[L, R]) FlatMapRight(f func(R) Either[L, R]) Either[L, R] {
	if e.IsRight() {
		return f(e.right)
	}
	return Left[L, R](e.left)
}
