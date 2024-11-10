package uncertain

import "math"

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
//  v2.Value = 0 - if the divisor value is 0, both value and error of the result are Inf.
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

// TODO
//func Cbrt(x float64) float64

// TODO
//func Pow(x, y float64) float64

// TODO
//func Pow10(n int) float64

func Sqrt(v Uncertain) (result Uncertain) {
	if v.Value == 0 {
		result.Value = 0
		result.Error = math.Sqrt(v.Error)
		return
	}
	result.Value = math.Sqrt(v.Value)
	result.Error = 1 / (2 * result.Value) * v.Error
	return
}
