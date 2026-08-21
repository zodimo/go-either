package either

import (
	"fmt"
	"testing"
)

func TestLeft(t *testing.T) {
	t.Run("creates left either with correct value", func(t *testing.T) {
		left := Left[string, int]("error")

		if !left.IsLeft() {
			t.Error("Expected IsLeft() to be true")
		}
		if left.IsRight() {
			t.Error("Expected IsRight() to be false")
		}

		leftValue := left.Left()
		if leftValue.IsNone() {
			t.Error("Expected Left() to return Some")
		}
		if leftValue.UnwrapUnsafe() != "error" {
			t.Errorf("Expected left value to be 'error', got %v", leftValue.UnwrapUnsafe())
		}

		rightValue := left.Right()
		if rightValue.IsSome() {
			t.Error("Expected Right() to return None")
		}
	})
}

func TestRight(t *testing.T) {
	t.Run("creates right either with correct value", func(t *testing.T) {
		right := Right[string, int](42)

		if right.IsLeft() {
			t.Error("Expected IsLeft() to be false")
		}
		if !right.IsRight() {
			t.Error("Expected IsRight() to be true")
		}

		rightValue := right.Right()
		if rightValue.IsNone() {
			t.Error("Expected Right() to return Some")
		}
		if rightValue.UnwrapUnsafe() != 42 {
			t.Errorf("Expected right value to be 42, got %v", rightValue.UnwrapUnsafe())
		}

		leftValue := right.Left()
		if leftValue.IsSome() {
			t.Error("Expected Left() to return None")
		}
	})
}

func TestMapLeft(t *testing.T) {
	t.Run("maps left value when either is left", func(t *testing.T) {
		left := Left[string, int]("hello")
		mapped := left.MapLeft(func(s string) string {
			return s + " world"
		})

		if !mapped.IsLeft() {
			t.Error("Expected mapped either to be left")
		}

		leftMaybe := mapped.Left()
		value := leftMaybe.UnwrapUnsafe()
		if value != "hello world" {
			t.Errorf("Expected mapped value to be 'hello world', got %v", value)
		}
	})

	t.Run("does not map when either is right", func(t *testing.T) {
		right := Right[string, int](42)
		mapped := right.MapLeft(func(s string) string {
			return s + " world"
		})

		if !mapped.IsRight() {
			t.Error("Expected mapped either to remain right")
		}

		rightMaybe := mapped.Right()
		value := rightMaybe.UnwrapUnsafe()
		if value != 42 {
			t.Errorf("Expected right value to remain 42, got %v", value)
		}
	})
}

func TestMapRight(t *testing.T) {
	t.Run("maps right value when either is right", func(t *testing.T) {
		right := Right[string, int](21)
		mapped := right.MapRight(func(i int) int {
			return i * 2
		})

		if !mapped.IsRight() {
			t.Error("Expected mapped either to be right")
		}

		rightMaybe := mapped.Right()
		value := rightMaybe.UnwrapUnsafe()
		if value != 42 {
			t.Errorf("Expected mapped value to be 42, got %v", value)
		}
	})

	t.Run("does not map when either is left", func(t *testing.T) {
		left := Left[string, int]("error")
		mapped := left.MapRight(func(i int) int {
			return i * 2
		})

		if !mapped.IsLeft() {
			t.Error("Expected mapped either to remain left")
		}

		leftMaybe := mapped.Left()
		value := leftMaybe.UnwrapUnsafe()
		if value != "error" {
			t.Errorf("Expected left value to remain 'error', got %v", value)
		}
	})
}

func TestFlatMapLeft(t *testing.T) {
	t.Run("flat maps left value when either is left", func(t *testing.T) {
		left := Left[string, int]("hello")
		mapped := left.FlatMapLeft(func(s string) Either[string, int] {
			if s == "hello" {
				return Right[string, int](100)
			}
			return Left[string, int]("error")
		})

		if !mapped.IsRight() {
			t.Error("Expected flat mapped either to be right")
		}

		rightMaybe := mapped.Right()
		value := rightMaybe.UnwrapUnsafe()
		if value != 100 {
			t.Errorf("Expected flat mapped value to be 100, got %v", value)
		}
	})

	t.Run("does not flat map when either is right", func(t *testing.T) {
		right := Right[string, int](42)
		mapped := right.FlatMapLeft(func(s string) Either[string, int] {
			return Left[string, int]("should not be called")
		})

		if !mapped.IsRight() {
			t.Error("Expected flat mapped either to remain right")
		}

		rightMaybe := mapped.Right()
		value := rightMaybe.UnwrapUnsafe()
		if value != 42 {
			t.Errorf("Expected right value to remain 42, got %v", value)
		}
	})
}

func TestFlatMapRight(t *testing.T) {
	t.Run("flat maps right value when either is right", func(t *testing.T) {
		right := Right[string, int](42)
		mapped := right.FlatMapRight(func(i int) Either[string, int] {
			if i > 40 {
				return Left[string, int]("too big")
			}
			return Right[string, int](i * 2)
		})

		if !mapped.IsLeft() {
			t.Error("Expected flat mapped either to be left")
		}

		leftMaybe := mapped.Left()
		value := leftMaybe.UnwrapUnsafe()
		if value != "too big" {
			t.Errorf("Expected flat mapped value to be 'too big', got %v", value)
		}
	})

	t.Run("does not flat map when either is left", func(t *testing.T) {
		left := Left[string, int]("error")
		mapped := left.FlatMapRight(func(i int) Either[string, int] {
			return Right[string, int](999)
		})

		if !mapped.IsLeft() {
			t.Error("Expected flat mapped either to remain left")
		}

		leftMaybe := mapped.Left()
		value := leftMaybe.UnwrapUnsafe()
		if value != "error" {
			t.Errorf("Expected left value to remain 'error', got %v", value)
		}
	})
}

// Test helper functions from helpers.go
func TestHelperMapLeft(t *testing.T) {
	t.Run("maps left value to different type when either is left", func(t *testing.T) {
		left := Left[string, int]("42")
		result := MapLeft(left, func(s string) int {
			return len(s)
		})

		if !result.IsLeft() {
			t.Error("Expected mapped either to be left")
		}

		leftMaybe := result.Left()
		value := leftMaybe.UnwrapUnsafe()
		if value != 2 {
			t.Errorf("Expected mapped value to be 2, got %v", value)
		}
	})

	t.Run("does not map when either is right", func(t *testing.T) {
		right := Right[string, int](42)
		result := MapLeft(right, func(s string) int {
			return len(s)
		})

		if !result.IsRight() {
			t.Error("Expected result to remain right")
		}

		rightMaybe := result.Right()
		value := rightMaybe.UnwrapUnsafe()
		if value != 42 {
			t.Errorf("Expected right value to remain 42, got %v", value)
		}
	})
}

func TestHelperMapRight(t *testing.T) {
	t.Run("maps right value to different type when either is right", func(t *testing.T) {
		right := Right[string, int](42)
		result := MapRight(right, func(i int) string {
			return "number: " + string(rune(i+'0'))
		})

		if !result.IsRight() {
			t.Error("Expected mapped either to be right")
		}

		rightMaybe := result.Right()
		value := rightMaybe.UnwrapUnsafe()
		expected := "number: " + string(rune(42+'0'))
		if value != expected {
			t.Errorf("Expected mapped value to be %v, got %v", expected, value)
		}
	})

	t.Run("does not map when either is left", func(t *testing.T) {
		left := Left[string, int]("error")
		result := MapRight(left, func(i int) string {
			return "should not be called"
		})

		if !result.IsLeft() {
			t.Error("Expected result to remain left")
		}

		leftMaybe := result.Left()
		value := leftMaybe.UnwrapUnsafe()
		if value != "error" {
			t.Errorf("Expected left value to remain 'error', got %v", value)
		}
	})
}

func TestHelperFlatMapLeft(t *testing.T) {
	t.Run("flat maps left value to different type when either is left", func(t *testing.T) {
		left := Left[string, int]("hello")
		result := FlatMapLeft(left, func(s string) Either[int, int] {
			return Right[int, int](len(s))
		})

		if !result.IsRight() {
			t.Error("Expected flat mapped either to be right")
		}

		rightMaybe := result.Right()
		value := rightMaybe.UnwrapUnsafe()
		if value != 5 {
			t.Errorf("Expected flat mapped value to be 5, got %v", value)
		}
	})

	t.Run("does not flat map when either is right", func(t *testing.T) {
		right := Right[string, int](42)
		result := FlatMapLeft(right, func(s string) Either[int, int] {
			return Left[int, int](999)
		})

		if !result.IsRight() {
			t.Error("Expected result to remain right")
		}

		rightMaybe := result.Right()
		value := rightMaybe.UnwrapUnsafe()
		if value != 42 {
			t.Errorf("Expected right value to remain 42, got %v", value)
		}
	})
}

func TestHelperFlatMapRight(t *testing.T) {
	t.Run("flat maps right value to different type when either is right", func(t *testing.T) {
		right := Right[string, int](42)
		result := FlatMapRight(right, func(i int) Either[string, string] {
			if i > 40 {
				return Left[string, string]("too big")
			}
			return Right[string, string]("ok")
		})

		if !result.IsLeft() {
			t.Error("Expected flat mapped either to be left")
		}

		leftMaybe := result.Left()
		value := leftMaybe.UnwrapUnsafe()
		if value != "too big" {
			t.Errorf("Expected flat mapped value to be 'too big', got %v", value)
		}
	})

	t.Run("does not flat map when either is left", func(t *testing.T) {
		left := Left[string, int]("error")
		result := FlatMapRight(left, func(i int) Either[string, string] {
			return Right[string, string]("should not be called")
		})

		if !result.IsLeft() {
			t.Error("Expected result to remain left")
		}

		leftMaybe := result.Left()
		value := leftMaybe.UnwrapUnsafe()
		if value != "error" {
			t.Errorf("Expected left value to remain 'error', got %v", value)
		}
	})
}

func TestMatch(t *testing.T) {
	t.Run("matches left value when either is left", func(t *testing.T) {
		left := Left[string, int]("error")
		result := left.Match(
			func(s string) int {
				return len(s)
			},
			func(i int) int {
				return i * 2
			},
		)

		if result != 5 {
			t.Errorf("Expected match result to be 5 (length of 'error'), got %v", result)
		}
	})

	t.Run("matches right value when either is right", func(t *testing.T) {
		right := Right[string, int](42)
		result := right.Match(
			func(s string) int {
				return len(s)
			},
			func(i int) int {
				return i * 2
			},
		)

		if result != 84 {
			t.Errorf("Expected match result to be 84 (42 * 2), got %v", result)
		}
	})

	t.Run("match with string return type", func(t *testing.T) {
		left := Left[int, string](100)
		result := left.Match(
			func(i int) string {
				return "left: " + string(rune(i+'0'))
			},
			func(s string) string {
				return "right: " + s
			},
		)

		expected := "left: " + string(rune(100+'0'))
		if result != expected {
			t.Errorf("Expected match result to be %v, got %v", expected, result)
		}
	})

	t.Run("match with complex transformation", func(t *testing.T) {
		right := Right[string, int](10)
		result := right.Match(
			func(s string) int {
				return -1
			},
			func(i int) int {
				return i*i + i
			},
		)

		if result != 110 {
			t.Errorf("Expected match result to be 110 (10*10 + 10), got %v", result)
		}
	})

	t.Run("match preserves left value", func(t *testing.T) {
		left := Left[string, int]("test")
		result := left.Match(
			func(s string) int {
				if s != "test" {
					t.Errorf("Expected left value to be 'test', got %v", s)
				}
				return 0
			},
			func(i int) int {
				t.Error("Right function should not be called for left either")
				return -1
			},
		)

		if result != 0 {
			t.Errorf("Expected match result to be 0, got %v", result)
		}
	})

	t.Run("match preserves right value", func(t *testing.T) {
		right := Right[string, int](99)
		result := right.Match(
			func(s string) int {
				t.Error("Left function should not be called for right either")
				return -1
			},
			func(i int) int {
				if i != 99 {
					t.Errorf("Expected right value to be 99, got %v", i)
				}
				return i
			},
		)

		if result != 99 {
			t.Errorf("Expected match result to be 99, got %v", result)
		}
	})
}

// Integration tests
func TestEitherChaining(t *testing.T) {
	t.Run("chains operations on left either", func(t *testing.T) {
		result := Left[string, int]("hello")

		// Chain map operations
		mapped := result.MapLeft(func(s string) string {
			return s + " world"
		}).MapLeft(func(s string) string {
			return "greeting: " + s
		})

		if !mapped.IsLeft() {
			t.Error("Expected chained result to be left")
		}

		leftMaybe := mapped.Left()
		value := leftMaybe.UnwrapUnsafe()
		if value != "greeting: hello world" {
			t.Errorf("Expected chained value to be 'greeting: hello world', got %v", value)
		}
	})

	t.Run("chains operations on right either", func(t *testing.T) {
		result := Right[string, int](10)

		// Chain map operations
		mapped := result.MapRight(func(i int) int {
			return i * 2
		}).MapRight(func(i int) int {
			return i + 2
		})

		if !mapped.IsRight() {
			t.Error("Expected chained result to be right")
		}

		rightMaybe := mapped.Right()
		value := rightMaybe.UnwrapUnsafe()
		if value != 22 {
			t.Errorf("Expected chained value to be 22, got %v", value)
		}
	})
}

// Method versions support type-changing mappings (L2 / R2).
func TestMethodMapLeftChangingType(t *testing.T) {
	t.Run("maps left value to a different type when either is left", func(t *testing.T) {
		left := Left[string, int]("hello")
		mapped := left.MapLeft(func(s string) int {
			return len(s)
		})

		if !mapped.IsLeft() {
			t.Error("Expected mapped either to be left")
		}

		leftMaybe := mapped.Left()
		value := leftMaybe.UnwrapUnsafe()
		if value != 5 {
			t.Errorf("Expected mapped value to be 5, got %v", value)
		}
	})

	t.Run("preserves right value when either is right", func(t *testing.T) {
		right := Right[string, int](42)
		mapped := right.MapLeft(func(s string) int {
			return len(s)
		})

		if !mapped.IsRight() {
			t.Error("Expected mapped either to remain right")
		}

		rightMaybe := mapped.Right()
		value := rightMaybe.UnwrapUnsafe()
		if value != 42 {
			t.Errorf("Expected right value to remain 42, got %v", value)
		}
	})
}

func TestMethodMapRightChangingType(t *testing.T) {
	t.Run("maps right value to a different type when either is right", func(t *testing.T) {
		right := Right[string, int](42)
		mapped := right.MapRight(func(i int) string {
			return fmt.Sprintf("%d", i)
		})

		if !mapped.IsRight() {
			t.Error("Expected mapped either to be right")
		}

		rightMaybe := mapped.Right()
		value := rightMaybe.UnwrapUnsafe()
		if value != "42" {
			t.Errorf("Expected mapped value to be \"42\", got %q", value)
		}
	})

	t.Run("preserves left value when either is left", func(t *testing.T) {
		left := Left[string, int]("error")
		mapped := left.MapRight(func(i int) string {
			return "should not be called"
		})

		if !mapped.IsLeft() {
			t.Error("Expected mapped either to remain left")
		}

		leftMaybe := mapped.Left()
		value := leftMaybe.UnwrapUnsafe()
		if value != "error" {
			t.Errorf("Expected left value to remain \"error\", got %q", value)
		}
	})
}

// Map defaults to MapRight; FlatMap defaults to FlatMapRight.
func TestMethodMap(t *testing.T) {
	t.Run("maps right value to a different type when either is right", func(t *testing.T) {
		right := Right[string, int](42)
		mapped := right.Map(func(i int) string {
			return fmt.Sprintf("%d", i)
		})

		if !mapped.IsRight() {
			t.Error("Expected mapped either to be right")
		}

		rightMaybe := mapped.Right()
		value := rightMaybe.UnwrapUnsafe()
		if value != "42" {
			t.Errorf("Expected mapped value to be \"42\", got %q", value)
		}
	})

	t.Run("preserves left value when either is left", func(t *testing.T) {
		left := Left[string, int]("error")
		mapped := left.Map(func(i int) string {
			return "should not be called"
		})

		if !mapped.IsLeft() {
			t.Error("Expected mapped either to remain left")
		}

		leftMaybe := mapped.Left()
		value := leftMaybe.UnwrapUnsafe()
		if value != "error" {
			t.Errorf("Expected left value to remain \"error\", got %q", value)
		}
	})
}

func TestMethodFlatMap(t *testing.T) {
	t.Run("flat maps right value to a different right type when either is right", func(t *testing.T) {
		right := Right[string, int](42)
		mapped := right.FlatMap(func(i int) Either[string, string] {
			return Right[string, string](fmt.Sprintf("%d", i))
		})

		if !mapped.IsRight() {
			t.Error("Expected flat mapped either to be right")
		}

		rightMaybe := mapped.Right()
		value := rightMaybe.UnwrapUnsafe()
		if value != "42" {
			t.Errorf("Expected flat mapped value to be \"42\", got %q", value)
		}
	})

	t.Run("preserves left value when either is left", func(t *testing.T) {
		left := Left[string, int]("error")
		mapped := left.FlatMap(func(i int) Either[string, string] {
			return Right[string, string]("should not be called")
		})

		if !mapped.IsLeft() {
			t.Error("Expected result to remain left")
		}

		leftMaybe := mapped.Left()
		value := leftMaybe.UnwrapUnsafe()
		if value != "error" {
			t.Errorf("Expected left value to remain \"error\", got %q", value)
		}
	})
}

func TestMethodFlatMapLeftChangingType(t *testing.T) {
	t.Run("flat maps left value to a different left type when either is left", func(t *testing.T) {
		left := Left[string, int]("hello")
		mapped := left.FlatMapLeft(func(s string) Either[int, int] {
			return Right[int, int](len(s))
		})

		if !mapped.IsRight() {
			t.Error("Expected flat mapped either to be right")
		}

		rightMaybe := mapped.Right()
		value := rightMaybe.UnwrapUnsafe()
		if value != 5 {
			t.Errorf("Expected flat mapped value to be 5, got %v", value)
		}
	})

	t.Run("preserves right value when either is right", func(t *testing.T) {
		right := Right[string, int](42)
		mapped := right.FlatMapLeft(func(s string) Either[int, int] {
			return Left[int, int](999)
		})

		if !mapped.IsRight() {
			t.Error("Expected result to remain right")
		}

		rightMaybe := mapped.Right()
		value := rightMaybe.UnwrapUnsafe()
		if value != 42 {
			t.Errorf("Expected right value to remain 42, got %v", value)
		}
	})
}

func TestMethodFlatMapRightChangingType(t *testing.T) {
	t.Run("flat maps right value to a different right type when either is right", func(t *testing.T) {
		right := Right[string, int](42)
		mapped := right.FlatMapRight(func(i int) Either[string, string] {
			return Right[string, string](fmt.Sprintf("%d", i))
		})

		if !mapped.IsRight() {
			t.Error("Expected flat mapped either to be right")
		}

		rightMaybe := mapped.Right()
		value := rightMaybe.UnwrapUnsafe()
		if value != "42" {
			t.Errorf("Expected flat mapped value to be \"42\", got %q", value)
		}
	})

	t.Run("preserves left value when either is left", func(t *testing.T) {
		left := Left[string, int]("error")
		mapped := left.FlatMapRight(func(i int) Either[string, string] {
			return Right[string, string]("should not be called")
		})

		if !mapped.IsLeft() {
			t.Error("Expected result to remain left")
		}

		leftMaybe := mapped.Left()
		value := leftMaybe.UnwrapUnsafe()
		if value != "error" {
			t.Errorf("Expected left value to remain \"error\", got %q", value)
		}
	})
}

// The package-level Match helper has no direct test yet; the method delegates
// to it, so verify it directly.
func TestHelperMatch(t *testing.T) {
	t.Run("matches left value when either is left", func(t *testing.T) {
		left := Left[string, int]("error")
		result := Match(left,
			func(s string) int {
				return len(s)
			},
			func(i int) int {
				return i * 2
			},
		)

		if result != 5 {
			t.Errorf("Expected match result to be 5, got %v", result)
		}
	})

	t.Run("matches right value when either is right", func(t *testing.T) {
		right := Right[string, int](42)
		result := Match(right,
			func(s string) int {
				return len(s)
			},
			func(i int) int {
				return i * 2
			},
		)

		if result != 84 {
			t.Errorf("Expected match result to be 84, got %v", result)
		}
	})
}

// Under Option A the methods are documented as delegating to the package-level
// helpers; this test asserts they produce identical results.
func TestMethodsDelegateToHelpers(t *testing.T) {
	t.Run("MapLeft", func(t *testing.T) {
		e := Left[string, int]("hello")
		assertEqual(t, "MapLeft",
			e.MapLeft(func(s string) int { return len(s) }),
			MapLeft(e, func(s string) int { return len(s) }),
		)
	})

	t.Run("MapRight", func(t *testing.T) {
		e := Right[string, int](42)
		assertEqual(t, "MapRight",
			e.MapRight(func(i int) string { return fmt.Sprintf("%d", i) }),
			MapRight(e, func(i int) string { return fmt.Sprintf("%d", i) }),
		)
	})

	t.Run("FlatMapLeft", func(t *testing.T) {
		e := Left[string, int]("hello")
		assertEqual(t, "FlatMapLeft",
			e.FlatMapLeft(func(s string) Either[int, int] { return Right[int, int](len(s)) }),
			FlatMapLeft(e, func(s string) Either[int, int] { return Right[int, int](len(s)) }),
		)
	})

	t.Run("FlatMapRight", func(t *testing.T) {
		e := Right[string, int](42)
		assertEqual(t, "FlatMapRight",
			e.FlatMapRight(func(i int) Either[string, string] { return Left[string, string]("big") }),
			FlatMapRight(e, func(i int) Either[string, string] { return Left[string, string]("big") }),
		)
	})

	t.Run("Map delegates to MapRight", func(t *testing.T) {
		e := Right[string, int](42)
		assertEqual(t, "Map",
			e.Map(func(i int) string { return fmt.Sprintf("%d", i) }),
			MapRight(e, func(i int) string { return fmt.Sprintf("%d", i) }),
		)
	})

	t.Run("FlatMap delegates to FlatMapRight", func(t *testing.T) {
		e := Right[string, int](42)
		assertEqual(t, "FlatMap",
			e.FlatMap(func(i int) Either[string, string] { return Left[string, string]("big") }),
			FlatMapRight(e, func(i int) Either[string, string] { return Left[string, string]("big") }),
		)
	})
}

func assertEqual[A, B comparable](t *testing.T, name string, got, want Either[A, B]) {
	t.Helper()
	if got.IsLeft() != want.IsLeft() || got.IsRight() != want.IsRight() {
		t.Fatalf("%s: side mismatch, got left=%v right=%v, want left=%v right=%v",
			name, got.IsLeft(), got.IsRight(), want.IsLeft(), want.IsRight())
	}
	if got.IsLeft() {
		if got.Left().UnwrapUnsafe() != want.Left().UnwrapUnsafe() {
			t.Fatalf("%s: left value mismatch, got %v, want %v",
				name, got.Left().UnwrapUnsafe(), want.Left().UnwrapUnsafe())
		}
	}
	if got.IsRight() {
		if got.Right().UnwrapUnsafe() != want.Right().UnwrapUnsafe() {
			t.Fatalf("%s: right value mismatch, got %v, want %v",
				name, got.Right().UnwrapUnsafe(), want.Right().UnwrapUnsafe())
		}
	}
}
