// Package uncertain implements a value with error and provides
// some mathematical functions that implement error propagation
//
// For + and - operations absolute error is a sum of absolute errors of operands.
//
// For * and / operations relative error is a sum of relative errors of operands.
//
// For function f(x) abolute error of a result is an abolute error of an argument multiplied by a function's derivative.
package uncertain

// Uncetraint type represents an uncertain value, i.e., value with error
type Uncertain struct {
	Value float64
	Error float64
}
