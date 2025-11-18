package either

import (
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
		
		if result.IsNone() {
			t.Error("Expected result to be Some")
		}
		
		mapped := result.UnwrapUnsafe()
		if !mapped.IsLeft() {
			t.Error("Expected mapped either to be left")
		}
		
		leftMaybe := mapped.Left()
		value := leftMaybe.UnwrapUnsafe()
		if value != 2 {
			t.Errorf("Expected mapped value to be 2, got %v", value)
		}
	})
	
	t.Run("returns None when either is right", func(t *testing.T) {
		right := Right[string, int](42)
		result := MapLeft(right, func(s string) int {
			return len(s)
		})
		
		if result.IsSome() {
			t.Error("Expected result to be None")
		}
	})
}

func TestHelperMapRight(t *testing.T) {
	t.Run("maps right value to different type when either is right", func(t *testing.T) {
		right := Right[string, int](42)
		result := MapRight(right, func(i int) string {
			return "number: " + string(rune(i+'0'))
		})
		
		if result.IsNone() {
			t.Error("Expected result to be Some")
		}
		
		mapped := result.UnwrapUnsafe()
		if !mapped.IsRight() {
			t.Error("Expected mapped either to be right")
		}
		
		rightMaybe := mapped.Right()
		value := rightMaybe.UnwrapUnsafe()
		expected := "number: " + string(rune(42+'0'))
		if value != expected {
			t.Errorf("Expected mapped value to be %v, got %v", expected, value)
		}
	})
	
	t.Run("returns None when either is left", func(t *testing.T) {
		left := Left[string, int]("error")
		result := MapRight(left, func(i int) string {
			return "should not be called"
		})
		
		if result.IsSome() {
			t.Error("Expected result to be None")
		}
	})
}

func TestHelperFlatMapLeft(t *testing.T) {
	t.Run("flat maps left value to different type when either is left", func(t *testing.T) {
		left := Left[string, int]("hello")
		result := FlatMapLeft(left, func(s string) Either[int, int] {
			return Right[int, int](len(s))
		})
		
		if result.IsNone() {
			t.Error("Expected result to be Some")
		}
		
		mapped := result.UnwrapUnsafe()
		if !mapped.IsRight() {
			t.Error("Expected flat mapped either to be right")
		}
		
		rightMaybe := mapped.Right()
		value := rightMaybe.UnwrapUnsafe()
		if value != 5 {
			t.Errorf("Expected flat mapped value to be 5, got %v", value)
		}
	})
	
	t.Run("returns None when either is right", func(t *testing.T) {
		right := Right[string, int](42)
		result := FlatMapLeft(right, func(s string) Either[int, int] {
			return Left[int, int](999)
		})
		
		if result.IsSome() {
			t.Error("Expected result to be None")
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
		
		if result.IsNone() {
			t.Error("Expected result to be Some")
		}
		
		mapped := result.UnwrapUnsafe()
		if !mapped.IsLeft() {
			t.Error("Expected flat mapped either to be left")
		}
		
		leftMaybe := mapped.Left()
		value := leftMaybe.UnwrapUnsafe()
		if value != "too big" {
			t.Errorf("Expected flat mapped value to be 'too big', got %v", value)
		}
	})
	
	t.Run("returns None when either is left", func(t *testing.T) {
		left := Left[string, int]("error")
		result := FlatMapRight(left, func(i int) Either[string, string] {
			return Right[string, string]("should not be called")
		})
		
		if result.IsSome() {
			t.Error("Expected result to be None")
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
