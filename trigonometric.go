package uncertain

import "math"

// Cos returns the cosine of the radian argument v.Value and propagates error.
//
// Special cases are:
//
//	Cos({±Inf, e}) = {NaN, _}
//	Cos({NaN, e}) = {NaN, _}
func Cos(v Uncertain) (result Uncertain) {
	s, c := math.Sincos(v.Value)

	result.Value = c
	result.Error = math.Abs(-s * v.Error)
	return
}

// Sin returns the sine of the radian argument v.Value and propagates error.
//
// Special cases are:
//
//	Sin({±0, e}) = {±0, re}
//	Sin({±Inf, e}) = {NaN, _}
//	Sin({NaN, e}) = {NaN, _}
func Sin(v Uncertain) (result Uncertain) {
	s, c := math.Sincos(v.Value)

	result.Value = s
	result.Error = math.Abs(c * v.Error)
	return
}

// Sincos returns Sin(x), Cos(x).
//
// Special cases are:
//
//	Sin({±0, e}) = {±0, re}, {1, re}
//	Sin({±Inf, e}) = {NaN, _}, {NaN, _}
//	Sin({NaN, e}) = {NaN, _}, {NaN, _}
func Sincos(v Uncertain) (sin, cos Uncertain) {
	s, c := math.Sincos(v.Value)

	sin.Value = s
	cos.Value = c

	sin.Error = math.Abs(c * v.Error)
	cos.Error = math.Abs(-s * v.Error)

	return
}

// Tan returns the tangent of the radian argument v.Value and propagates error.
//
// Special cases are:
//
//	Tan({±0, e}) = {±0, re}
//	Tan({±Inf, e}) = {NaN, _}
//	Tan({NaN, e}) = {NaN, _}
func Tan(v Uncertain) (result Uncertain) {
	result.Value = math.Tan(v.Value)
	result.Error = (result.Value*result.Value + 1) * v.Error
	return
}
