// Package uncertain implements a value with error and provides
// some mathematical functions that implement error propagation
package uncertain

import "math"

// Uncetraint type represents an uncertain value, i.e., value with error
type Uncertain struct {
	Value float64
	Error float64
}

// Add method returs the sum of its receiver and its argument. Absolute error is a sum of absolute errors.
func (v1 Uncertain) Add(v2 Uncertain) (sum Uncertain) {
	sum.Value = v1.Value + v2.Value
	sum.Error = v1.Error + v2.Error
	return
}

// Sub method returs the difference between its receiver and its argument. Absolute error is a sum of absolute errors.
func (v1 Uncertain) Sub(v2 Uncertain) (diff Uncertain) {
	diff.Value = v1.Value - v2.Value
	diff.Error = v1.Error + v2.Error
	return
}

// Mul method returs the product of its receiver and its argument. Relative error is a sum of relative errors.
//
// Special cases are:
//
//  v1.Value = 0
//  v2.Value = 0
//  - impossible to calculate relative error if the value is zero. Error is calculated by special function.
func (v1 Uncertain) Mul(v2 Uncertain) (product Uncertain) {
	if v1.Value*v2.Value == 0 {
		return v1.mul(v2)
	}

	product.Value = v1.Value * v2.Value

	relError := 0.0
	if v1.Value != 0 {
		relError += v1.Error / math.Abs(v1.Value)
	}
	if v2.Value != 0 {
		relError += v2.Error / math.Abs(v2.Value)
	}

	product.Error = relError * math.Abs(product.Value)
	return
}

// mul is a special case function for Mul.
// It works fine for any values including zeroes
func (v1 Uncertain) mul(v2 Uncertain) (product Uncertain) {
	product.Value = v1.Value * v2.Value

	var val [4]float64

	val[0] = (v1.Value + v1.Error) * (v2.Value + v2.Error)
	val[1] = (v1.Value + v1.Error) * (v2.Value - v2.Error)
	val[2] = (v1.Value - v1.Error) * (v2.Value - v2.Error)
	val[3] = (v1.Value - v1.Error) * (v2.Value + v2.Error)

	max, min := val[0], val[0]

	for _, v := range val {
		max = math.Max(max, v)
		min = math.Min(min, v)
	}
	product.Error = (max - min) / 2

	return
}

// Div method returs the quotient of its receiver-dividend and its argument-divisor. Relative error is a sum of relative errors.
//
// Special cases are:
//
//  v1.Value = 0 - impossible to calculate relative error if the value is zero. Error is calculated by special function.
//  v2.Value = 0 - if he divisor value is 0, both value and error of the result are NaN.
func (v1 Uncertain) Div(v2 Uncertain) (quotient Uncertain) {
	if v1.Value == 0 {
		return v1.div(v2)
	}

	quotient.Value = v1.Value / v2.Value

	relError := 0.0
	relError += v1.Error / math.Abs(v1.Value)
	relError += v2.Error / math.Abs(v2.Value) // No  need to check for division by zero: in this case function will fail in the fourth line

	quotient.Error = relError * math.Abs(quotient.Value)
	return
}

// mul is a special case function for Mul.
// It works fine for any values but zero divisor
func (v1 Uncertain) div(v2 Uncertain) (quotient Uncertain) {
	rel := v2.Error / v2.Value
	v2.Value = 1 / v2.Value
	v2.Error = v2.Value * rel

	return v1.mul(v2)
}

// Acos returns the arccosine, in radians, of v.Value and propagates error.
//
// Special case is:
//
//  Acos({x, e}) = {NaN, _} if x < -1 or x > 1
func Acos(v Uncertain) (result Uncertain) {
	if v.Error == 0 {
		result.Value = math.Acos(v.Value)
		result.Error = 0
		return
	}

	if v.Value == -1 {
		r := math.Acos(-1 + v.Error)
		result.Value = math.Pi
		result.Error = math.Pi - r
		return
	}
	if v.Value == 1 {
		result.Value = 0
		result.Error = math.Acos(1 - v.Error)
		return
	}

	result.Value = math.Acos(v.Value)
	result.Error = math.Abs(-1 / (math.Sqrt(1 - v.Value*v.Value)) * v.Error)
	return
}

//func Acosh(x float64) float64

// Asin returns the arcsine, in radians, of x.
//
// Special cases are:
//
//	Asin({±0, e}) = {±0, e}
//	Asin({x, e}) = {NaN, _} if x < -1 or x > 1
func Asin(v Uncertain) (result Uncertain) {
	result.Value = math.Asin(v.Value)

	if v.Error == 0 {
		result.Error = 0
		return
	}

	if v.Value == -1 {
		result.Error = math.Asin(-1+v.Error) - result.Value
		return
	}
	if v.Value == 1 {
		result.Error = result.Value - math.Asin(1-v.Error)
		return
	}

	result.Error = math.Abs(1 / (math.Sqrt(1 - v.Value*v.Value)) * v.Error)
	return
}

//func Asinh(x float64) float64

// Atan returns the arctangent, in radians, of x.
//
// Special cases are:
//
//	Atan({±0, e}) = {±0, e}
//	Atan({±Inf, e}) = {±Pi/2, 0}
func Atan(v Uncertain) (result Uncertain) {
	result.Value = math.Atan(v.Value)
	result.Error = v.Error * (1.0 / (1.0 + v.Value*v.Value))
	return
}

func Atan2(v Uncertain) (result Uncertain) {

	return
}

//func Atanh(x float64) float64

//func Cbrt(x float64) float64

//func Cos(x float64) float64

//func Cosh(x float64) float64

//func Exp(x float64) float64

//func Exp2(x float64) float64

//func Log(x float64) float64

//func Log10(x float64) float64

//func Log2(x float64) float64

//func Pow(x, y float64) float64

//func Pow10(n int) float64

//func Sin(x float64) float64

//func Sinh(x float64) float64

//func Sqrt(x float64) float64

//func Tan(x float64) float64

//func Tanh(x float64) float64