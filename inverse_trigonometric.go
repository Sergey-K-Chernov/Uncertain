package uncertain

import "math"

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

// Arccos is a synonym for Acos
func Arccos(v Uncertain) (result Uncertain) {
	return Acos(v)
}

// Asin returns the arcsine, in radians, of v.Value and propagates error.
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

// Arcsin is a synonym for Asin
func Arcsin(v Uncertain) (result Uncertain) {
	return Asin(v)
}

// Atan returns the arctangent, in radians, of v.Value and propagates error.
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

// Arctg is a synonym for Atan
func Arctg(v Uncertain) (result Uncertain) {
	return Atan(v)
}

// Atan2 returns the arc tangent of y/x, using
// the signs of the two to determine the quadrant
// of the return value, and propagates error.
//
// Special cases for values are the same as the special cases for math.Atan2
// Special cases for errors are:
//
//	Atan2({y, ey}, {NaN, ex}) = {NaN, _}
//	Atan2({NaN, ey}, {x, ex}) = {NaN, _}
//	Atan2({±0, 0}, {x>=0, ex}) = {±0, 0}
//	Atan2({±0, 0}, {x<=-0, ex}) = {±Pi, 0}
//	Atan2({y>0, ey}, {0, 0}) = {+Pi/2, 0}
//	Atan2({y<0, ey}, {0, 0}) = {-Pi/2, 0}
//	Atan2({+Inf, ey}, {+Inf, ex}) = {+Pi/4, 0}
//	Atan2({-Inf, ey}, {+Inf, ex}) = {-Pi/4, 0}
//	Atan2({+Inf, ey}, {-Inf, ex}) = {3Pi/4, 0}
//	Atan2({-Inf, ey}, {-Inf, ex}) = {-3Pi/4, 0}
//	Atan2({y, ey}, {+Inf, ex}) = {0, 0}
//	Atan2({y>0, ey}, {-Inf, ex}) = {+Pi, 0}
//	Atan2({y<0, ey}, {-Inf, ex}) = {-Pi, 0}
//	Atan2({+Inf, ey}, {x, ex}) = {+Pi/2, 0}
//	Atan2({-Inf, ey}, {x, ex}) = {-Pi/2, 0}
func Atan2(y, x Uncertain) (result Uncertain) {
	var res [4]float64
	result.Value = math.Atan2(y.Value, x.Value)

	if y.Value == math.Inf(1) || y.Value == math.Inf(-1) || x.Value == math.Inf(1) || x.Value == math.Inf(-1) {
		result.Error = 0
		return
	}

	res[0] = math.Atan2(y.Value+y.Error, x.Value+x.Error)
	res[1] = math.Atan2(y.Value+y.Error, x.Value-x.Error)
	res[2] = math.Atan2(y.Value-y.Error, x.Value+x.Error)
	res[3] = math.Atan2(y.Value-y.Error, x.Value-x.Error)

	min, max := res[0], res[0]

	for i := 1; i < 4; i++ {
		min = math.Min(min, res[i])
		max = math.Max(max, res[i])
	}
	if min == -max && max == math.Pi {
		result.Error = 0
		return
	}

	result.Error = math.Abs((max - min) / 2)
	return
}
