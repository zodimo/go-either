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
if e.IsLeft() {
    // Handle left case
}

if e.IsRight() {
    // Handle right case
}
```

### Extracting Values

```go
// Get the left value (returns Maybe[L])
leftValue := e.Left()
if leftValue.IsSome() {
    value := leftValue.UnwrapUnsafe()
    // Use value
}

// Get the right value (returns Maybe[R])
rightValue := e.Right()
if rightValue.IsSome() {
    value := rightValue.UnwrapUnsafe()
    // Use value
}
```

### Mapping

`MapLeft` and `MapRight` transform the value held in the Either. They can optionally change the mapped type (from `L` to `L2`, or `R` to `R2`) and are available both as methods and as package-level helper functions; the methods delegate to the helpers, so the two styles are interchangeable.

```go
// e is an Either[L, R]; map over the left value, optionally changing its type to L2
mapped := e.MapLeft(func(l L) L2 {
    return transformed
})

// Map over the right value, optionally changing its type to R2
mapped := e.MapRight(func(r R) R2 {
    return transformed
})
```

Mapping only runs when the Either holds a value on the mapped side; the other side is returned unchanged.

#### `Map` (defaults to `MapRight`)

`Map` is a right-biased convenience alias for `MapRight` — useful when Right represents the success value (the common case), and for forward-compatibility if the semantics ever change.

```go
mapped := e.Map(func(r R) R2 {
    return transformed
})
```

### Flat Mapping

`FlatMapLeft` and `FlatMapRight` chain operations that themselves return an `Either`, flattening the result. They support type changes and, like the map functions, exist as both methods and package-level helpers.

```go
// Flat map over the left value, returning a new Either with a different left type
flatMapped := e.FlatMapLeft(func(l L) Either[L2, R] {
    return either.Right[L2, R](someValue)
})

// Flat map over the right value, returning a new Either with a different right type
flatMapped := e.FlatMapRight(func(r R) Either[L, R2] {
    return either.Left[L, R2]("error")
})
```

#### `FlatMap` (defaults to `FlatMapRight`)

`FlatMap` is a right-biased convenience alias for `FlatMapRight`:

```go
flatMapped := e.FlatMap(func(r R) Either[L, R2] {
    return either.Left[L, R2]("error")
})
```

### Matching

The `Match` method (and its package-level helper) extract and transform the value from an Either, regardless of whether it's Left or Right. Both functions must return the same type `T`:

```go
// Match on either value, transforming both to the same type T
result := e.Match(
    func(l L) T {
        return transformedLeft
    },
    func(r R) T {
        return transformedRight
    },
)
```

### Package-Level Helper Functions

Every method has a package-level helper equivalent. The methods simply delegate to these helpers, so both styles are interchangeable — use whichever reads better in your code.

```go
// Map left to a different type
result := either.MapLeft(e, func(l L) L2 {
    return newValue
})

// Map right to a different type
result := either.MapRight(e, func(r R) R2 {
    return newValue
})

// Flat map left to a different type
result := either.FlatMapLeft(e, func(l L) Either[L2, R] {
    return either.Right[L2, R](value)
})

// Flat map right to a different type
result := either.FlatMapRight(e, func(r R) Either[L, R2] {
    return either.Left[L, R2]("error")
})

// Match either value to a single result
result := either.Match(e,
    func(l L) T {
        return transformedLeft
    },
    func(r R) T {
        return transformedRight
    },
)
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

