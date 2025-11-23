# go-either

A Go implementation of the Either monad, providing a type-safe way to represent values that can be one of two types (Left or Right).

## Overview

`go-either` provides an `Either[L, R]` type that represents a value that is either of type `L` (Left) or type `R` (Right). This is useful for error handling, where Left typically represents an error and Right represents a success value, though it can be used for any scenario where you need to represent one of two possible types.

## Installation

```bash
go get github.com/zodimo/go-either
```

## Usage

### Creating Either Values

```go
import "github.com/zodimo/go-either"

// Create a Left value (typically used for errors)
left := either.Left[string, int]("error message")

// Create a Right value (typically used for success)
right := either.Right[string, int](42)
```

### Checking the Value

```go
if either.IsLeft() {
    // Handle left case
}

if either.IsRight() {
    // Handle right case
}
```

### Extracting Values

```go
// Get the left value (returns Maybe[L])
leftValue := either.Left()
if leftValue.IsSome() {
    value := leftValue.UnwrapUnsafe()
    // Use value
}

// Get the right value (returns Maybe[R])
rightValue := either.Right()
if rightValue.IsSome() {
    value := rightValue.UnwrapUnsafe()
    // Use value
}
```

### Mapping

```go
// Map over the left value
mapped := either.MapLeft(func(l L) L {
    // Transform left value
    return transformed
})

// Map over the right value
mapped := either.MapRight(func(r R) R {
    // Transform right value
    return transformed
})
```

### Flat Mapping

```go
// Flat map over the left value
flatMapped := either.FlatMapLeft(func(l L) Either[L, R] {
    // Return a new Either based on left value
    return either.Right[L, R](someValue)
})

// Flat map over the right value
flatMapped := either.FlatMapRight(func(r R) Either[L, R] {
    // Return a new Either based on right value
    return either.Left[L, R]("error")
})
```

### Matching

The `Match` function allows you to extract and transform the value from an Either, regardless of whether it's Left or Right. Both functions must return the same type:

```go
// Match on either value, transforming both to the same type
result := either.Match(
    func(l L) R {
        // Transform left value to R
        return transformedLeft
    },
    func(r R) R {
        // Transform right value to R
        return transformedRight
    },
)
```

### Helper Functions

The package also provides helper functions that return `Maybe[Either]` for type-changing operations:

```go
// Map left to a different type
result := either.MapLeft(either, func(l L) L2 {
    // Transform to different type
    return newValue
})

// Map right to a different type
result := either.MapRight(either, func(r R) R2 {
    // Transform to different type
    return newValue
})

// Flat map left to a different type
result := either.FlatMapLeft(either, func(l L) Either[L2, R] {
    // Return Either with different left type
    return either.Right[L2, R](value)
})

// Flat map right to a different type
result := either.FlatMapRight(either, func(r R) Either[L, R2] {
    // Return Either with different right type
    return either.Left[L, R2]("error")
})
```

## Example: Error Handling

```go
func divide(a, b int) either.Either[string, int] {
    if b == 0 {
        return either.Left[string, int]("division by zero")
    }
    return either.Right[string, int](a / b)
}

result := divide(10, 2)
if result.IsRight() {
    value := result.Right().UnwrapUnsafe()
    fmt.Printf("Result: %d\n", value)
} else {
    error := result.Left().UnwrapUnsafe()
    fmt.Printf("Error: %s\n", error)
}
```

## Example: Using Match

```go
func divide(a, b int) either.Either[string, int] {
    if b == 0 {
        return either.Left[string, int]("division by zero")
    }
    return either.Right[string, int](a / b)
}

result := divide(10, 2)
// Use Match to handle both cases and return a single value
message := result.Match(
    func(err string) string {
        return "Error: " + err
    },
    func(value int) string {
        return fmt.Sprintf("Success: %d", value)
    },
)
fmt.Println(message) // Output: "Success: 5"
```

## Dependencies

- `github.com/zodimo/go-maybe` - Used for optional value handling

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

