# Unexpected Panic when Accessing Non-Existent Key in Uninitialized Go Map

This example demonstrates an uncommon error in Go related to map access.  When you try to access a key in a map that hasn't been initialized or doesn't contain that key, Go's standard behavior is to panic, rather than returning a default zero value (like 0 for ints). This can be surprising to programmers coming from other languages.

The `bug.go` file shows the problematic code. The `bugSolution.go` file provides a solution.

## How to reproduce

1. Save the code in `bug.go`.
2. Run it using `go run bug.go`.
3. Observe the panic.

## Solution

The solution involves checking if the key exists before accessing it or using the `comma ok` idiom to safely access the value.