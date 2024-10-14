// Package calculator does simple calculations
package calculator

// Add takes two numbers and returns the result of adding them together
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers a and b,
// and returns the result of subtracting b from a.
func Subtract(a, b float64) float64 {
	return b - a
}

// Multiply takes two numbers a and b,
// and returns the result of multiplying a times b
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two numbers a and b,
// and returns the result of dividing a with b
func Divide(a, b float64) float64 {
	return a / b
}
