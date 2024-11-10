package uncertain

import (
	"math"
)

func almostEqual(a, b Uncertain) bool {
	if a.Value == math.Inf(1) || b.Value == math.Inf(1) || a.Value == math.Inf(-1) || b.Value == math.Inf(-1) {
		return (a.Value == b.Value && a.Error == b.Error)
	}
	threshold := 0.000000000001

	ok := true
	if b.Value == 0 {
		ok = ok && math.Abs(a.Value-b.Value) <= threshold
	} else {
		ok = ok && math.Abs(a.Value/b.Value-1) <= threshold
	}

	if b.Error == 0 {
		ok = ok && math.Abs(a.Error-b.Error) <= threshold
	} else {
		ok = ok && math.Abs(a.Error/b.Error-1) <= threshold
	}
	return ok
}

/*
func TestAcosh(t *testing.T) {

}
*/

/*
func TestAsinh(t *testing.T) {

}
*/

/*
// TODO
func TestAtanh(t *testing.T) {
}
*/

// TODO
//func Cosh(t *testing.T) {

// TODO
//func Exp(t *testing.T) {

// TODO
//func Exp2(t *testing.T) {

// TODO
//func Log(t *testing.T) {

// TODO
//func Log10(t *testing.T) {

// TODO
//func Pow10(t *testing.T) {

// TODO
//func TestSinh(t *testing.T) {

// TODO
//func TestTanh(t *testing.T) {
