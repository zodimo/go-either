package either

import (
	"fmt"
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

func (e Either[L, R]) MapLeft[L2 any](f func(L) L2) Either[L2, R] {
	return MapLeft(e, f)
}

func (e Either[L, R]) MapRight[R2 any](f func(R) R2) Either[L, R2] {
	return MapRight(e, f)
}

func (e Either[L, R]) Map[R2 any](f func(R) R2) Either[L, R2] {
	return MapRight(e, f)
}

func (e Either[L, R]) FlatMapLeft[L2 any](f func(L) Either[L2, R]) Either[L2, R] {
	return FlatMapLeft(e, f)
}

func (e Either[L, R]) FlatMap[R2 any](f func(R) Either[L, R2]) Either[L, R2] {
	return FlatMapRight(e, f)
}

func (e Either[L, R]) FlatMapRight[R2 any](f func(R) Either[L, R2]) Either[L, R2] {
	return FlatMapRight(e, f)
}

func (e Either[L, R]) Match[T any](left func(L) T, right func(R) T) T {
	return Match(e, left, right)
}

func (e Either[L, R]) String() string {
	if e.isLeft {
		return fmt.Sprintf("Left[%T](%v)", e.left, e.left)
	}
	return fmt.Sprintf("Right[%T](%v)", e.right, e.right)
}
