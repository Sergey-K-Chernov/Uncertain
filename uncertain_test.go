package uncertain

import (
	"math"
	"testing"
)

func almostEqual(a, b Uncertain) bool {
	threshold := 0.0000000000001
	return (math.Abs(a.Value-b.Value) <= threshold) && (math.Abs(a.Error-b.Error) <= threshold)
}

func TestAdd(t *testing.T) {
	v1 := []Uncertain{
		{0, 0}, {0, 3}, {3, 3}, {1, 0.1}, {1, 0.1}, {-1, 0.1}, {-1, 0.1},
	}
	v2 := []Uncertain{
		{0, 0}, {3, 0}, {0, 3}, {1, 0.1}, {-1, 0.1}, {1, 0.1}, {-1, 0.1},
	}

	res := []Uncertain{
		{0, 0}, {3, 3}, {3, 6}, {2, 0.2}, {0, 0.2}, {0, 0.2}, {-2, 0.2},
	}

	if len(v1) != len(v2) || len(v1) != len(res) {
		t.Fatalf("Wrong contitions. Array lengths are not equal")
	}

	for i := range v1 {
		if !almostEqual(v1[i].Add(v2[i]), res[i]) {
			t.Fatalf("Test %d failed", i)
		}
	}
}

func TestSub(t *testing.T) {
	v1 := []Uncertain{
		{0, 0}, {10, 1}, {-100, 10}, {-100, 10}, {200, 50}, {200, 50},
	}
	v2 := []Uncertain{
		{0, 0}, {5, 0.5}, {100, 10}, {-100, 10}, {300, 30}, {-300, 30},
	}

	res := []Uncertain{
		{0, 0}, {5, 1.5}, {-200, 20}, {0, 20}, {-100, 80}, {500, 80},
	}

	if len(v1) != len(v2) || len(v1) != len(res) {
		t.Fatalf("Wrong contitions. Array lengths are not equal")
	}

	for i := range v1 {
		if !almostEqual(v1[i].Sub(v2[i]), res[i]) {
			t.Fatalf("Test %d failed", i)
		}
	}
}

func TestMul(t *testing.T) {
	v1 := []Uncertain{
		{7, 0}, {0, 0}, {10, 1}, {-2, 0.1}, {100, 5}, {0, 1}, {10, 1},
	}
	v2 := []Uncertain{
		{6, 0}, {0, 0}, {5, 0.5}, {2, 0.1}, {-200, 40}, {10, 1}, {0, 1},
	}

	res := []Uncertain{
		{42, 0}, {0, 0}, {50, 10}, {-4, 0.4}, {-20000, 5000}, {0, 11}, {0, 11},
	}

	if len(v1) != len(v2) || len(v1) != len(res) {
		t.Fatalf("Wrong contitions. Array lengths are not equal")
	}

	for i := range v1 {
		if !almostEqual(v1[i].Mul(v2[i]), res[i]) {
			t.Fatalf("Test %d failed", i)
		}
	}
}

func TestMul2(t *testing.T) {
	v1 := []Uncertain{
		{7, 0}, {0, 0}, {10, 1}, {-2, 0.1}, {100, 5}, {0, 1}, {10, 1},
	}
	v2 := []Uncertain{
		{6, 0}, {0, 0}, {5, 0.5}, {2, 0.1}, {-200, 40}, {10, 1}, {0, 1},
	}

	res := []Uncertain{
		{42, 0}, {0, 0}, {50, 10}, {-4, 0.4}, {-20000, 5000}, {0, 11}, {0, 11},
	}

	if len(v1) != len(v2) || len(v1) != len(res) {
		t.Fatalf("Wrong contitions. Array lengths are not equal")
	}

	for i := range v1 {
		if !almostEqual(v1[i].mul(v2[i]), res[i]) {
			t.Fatalf("Test %d failed", i)
		}
	}
}

func TestDiv(t *testing.T) {
	v1 := []Uncertain{
		{0, 10}, {0, 0}, {0, 10}, {10, 1}, {500, 25}, {-300, 20},
	}
	v2 := []Uncertain{
		{1, 0}, {1, 0}, {1, 0.1}, {10, 1}, {-10, 5}, {3, 0.5},
	}

	res := []Uncertain{
		{0, 10}, {0, 0}, {0, 11}, {1, 0.2}, {-50, 27.5}, {-100, 23.3333333333333},
	}

	if len(v1) != len(v2) || len(v1) != len(res) {
		t.Fatalf("Wrong contitions. Array lengths are not equal")
	}

	if len(v1) != len(v2) || len(v1) != len(res) {
		t.Fatalf("Wrong contitions. Array lengths are not equal")
	}

	for i := range v1 {
		if !almostEqual(v1[i].Div(v2[i]), res[i]) {
			t.Fatalf("Test %d failed", i)
		}
	}
}

func TestDiv2(t *testing.T) {
	v1 := []Uncertain{
		{0, 10}, {0, 0}, {0, 10}, {10, 1}, {500, 25}, {-300, 20},
	}
	v2 := []Uncertain{
		{1, 0}, {1, 0}, {1, 0.1}, {10, 1}, {-10, 5}, {3, 0.5},
	}

	res := []Uncertain{
		{0, 10}, {0, 0}, {0, 11}, {1, 0.2}, {-50, 27.5}, {-100, 23.3333333333333},
	}

	if len(v1) != len(v2) || len(v1) != len(res) {
		t.Fatalf("Wrong contitions. Array lengths are not equal")
	}

	for i := range v1 {
		if !almostEqual(v1[i].div(v2[i]), res[i]) {
			t.Fatalf("Test %d failed", i)
		}
	}
}

func TestAcos(t *testing.T) {
	cases := [][2]Uncertain{
		{{-1.5, 0.1}, {math.NaN(), math.NaN()}},
		{{-1.5, 0}, {math.NaN(), 0}},
		{{1.5, 0}, {math.NaN(), 0}},
		{{-1, 0}, {3.14159265358979, 0}},
		{{-1, 0.1}, {3.14159265358979, 0.451026811796262}},
		{{-1, 0.2}, {3.14159265358979, 0.643501108793284}},
		{{-0.75, 0}, {2.41885840577638, 0}},
		{{-0.75, 0.1}, {2.41885840577638, 0.151185789203691}},
		{{-0.75, 0.2}, {2.41885840577638, 0.302371578407382}},
		{{-0.5, 0}, {2.0943951023932, 0}},
		{{-0.5, 0.1}, {2.0943951023932, 0.115470053837925}},
		{{-0.5, 0.2}, {2.0943951023932, 0.23094010767585}},
		{{-0.25, 0}, {1.82347658193698, 0}},
		{{-0.25, 0.1}, {1.82347658193698, 0.103279555898864}},
		{{-0.25, 0.2}, {1.82347658193698, 0.206559111797729}},
		{{0, 0}, {1.5707963267949, 0}},
		{{0, 0.1}, {1.5707963267949, 0.1}},
		{{0, 0.2}, {1.5707963267949, 0.2}},
		{{0.25, 0}, {1.31811607165282, 0}},
		{{0.25, 0.1}, {1.31811607165282, 0.103279555898864}},
		{{0.25, 0.2}, {1.31811607165282, 0.206559111797729}},
		{{0.5, 0}, {1.0471975511966, 0}},
		{{0.5, 0.1}, {1.0471975511966, 0.115470053837925}},
		{{0.5, 0.2}, {1.0471975511966, 0.23094010767585}},
		{{0.75, 0}, {0.722734247813416, 0}},
		{{0.75, 0.1}, {0.722734247813416, 0.151185789203691}},
		{{0.75, 0.2}, {0.722734247813416, 0.302371578407382}},
		{{1, 0}, {0, 0}},
		{{1, 0.1}, {0, 0.451026811796262}},
		{{1, 0.2}, {0, 0.643501108793284}},
	}

	for i, the_case := range cases {
		res := Acos(the_case[0])

		if math.IsNaN(res.Value) {
			if !math.IsNaN(the_case[1].Value) {
				t.Fatalf("Test case %d failed: Acos(%f±%f) is %f±%f, got %f±%f",
					i, the_case[0].Value, the_case[0].Error, the_case[1].Value, the_case[1].Error, res.Value, res.Error)
			} else {
				continue
			}
		}

		if !almostEqual(res, the_case[1]) {
			t.Fatalf("Test case %d failed: Acos(%f±%f) is %f±%f, got %f±%f",
				i, the_case[0].Value, the_case[0].Error, the_case[1].Value, the_case[1].Error, res.Value, res.Error)
		}
	}
}

/*
func TestAcosh(t *testing.T) {

}
*/

func TestAsin(t *testing.T) {
	cases := [][2]Uncertain{
		{{-1.5, 0.1}, {math.NaN(), math.NaN()}},
		{{-1.5, 0}, {math.NaN(), 0}},
		{{1.5, 0}, {math.NaN(), 0}},
		{{-1, 0}, {-1.5707963267949, 0}},
		{{-1, 0.1}, {-1.5707963267949, 0.451026811796262}},
		{{-1, 0.2}, {-1.5707963267949, 0.643501108793284}},
		{{-0.75, 0}, {-0.848062078981481, 0}},
		{{-0.75, 0.1}, {-0.848062078981481, 0.151185789203691}},
		{{-0.75, 0.2}, {-0.848062078981481, 0.302371578407382}},
		{{-0.5, 0}, {-0.523598775598299, 0}},
		{{-0.5, 0.1}, {-0.523598775598299, 0.115470053837925}},
		{{-0.5, 0.2}, {-0.523598775598299, 0.23094010767585}},
		{{-0.25, 0}, {-0.252680255142079, 0}},
		{{-0.25, 0.1}, {-0.252680255142079, 0.103279555898864}},
		{{-0.25, 0.2}, {-0.252680255142079, 0.206559111797729}},
		{{0, 0}, {0, 0}},
		{{0, 0.1}, {0, 0.1}},
		{{0, 0.2}, {0, 0.2}},
		{{0.25, 0}, {0.252680255142079, 0}},
		{{0.25, 0.1}, {0.252680255142079, 0.103279555898864}},
		{{0.25, 0.2}, {0.252680255142079, 0.206559111797729}},
		{{0.5, 0}, {0.523598775598299, 0}},
		{{0.5, 0.1}, {0.523598775598299, 0.115470053837925}},
		{{0.5, 0.2}, {0.523598775598299, 0.23094010767585}},
		{{0.75, 0}, {0.848062078981481, 0}},
		{{0.75, 0.1}, {0.848062078981481, 0.151185789203691}},
		{{0.75, 0.2}, {0.848062078981481, 0.302371578407382}},
		{{1, 0}, {1.5707963267949, 0}},
		{{1, 0.1}, {1.5707963267949, 0.451026811796262}},
		{{1, 0.2}, {1.5707963267949, 0.643501108793284}},
	}

	for i, the_case := range cases {
		res := Asin(the_case[0])
		if !almostEqual(res, the_case[1]) {
			t.Fatalf("Test case %d failed: Asin(%f±%f) is %f±%f, got %f±%f",
				i, the_case[0].Value, the_case[0].Error, the_case[1].Value, the_case[1].Error, res.Value, res.Error)
		}
	}
}

/*
func TestAsinh(t *testing.T) {

}
*/

func TestAtan(t *testing.T) {
	cases := [][2]Uncertain{
		{{math.Inf(-1), 10000}, {-math.Pi / 2, 0}},
		{{math.Inf(1), 10000}, {math.Pi / 2, 0}},
		{{-1, 0}, {-0.785398163397448, 0}},
		{{-1, 0.1}, {-0.785398163397448, 0.05}},
		{{-1, 0.2}, {-0.785398163397448, 0.1}},
		{{-0.75, 0}, {-0.643501108793284, 0}},
		{{-0.75, 0.1}, {-0.643501108793284, 0.064}},
		{{-0.75, 0.2}, {-0.643501108793284, 0.128}},
		{{-0.5, 0}, {-0.463647609000806, 0}},
		{{-0.5, 0.1}, {-0.463647609000806, 0.08}},
		{{-0.5, 0.2}, {-0.463647609000806, 0.16}},
		{{-0.25, 0}, {-0.244978663126864, 0}},
		{{-0.25, 0.1}, {-0.244978663126864, 0.094117647058824}},
		{{-0.25, 0.2}, {-0.244978663126864, 0.188235294117647}},
		{{0, 0}, {0, 0}},
		{{0, 0.1}, {0, 0.1}},
		{{0, 0.2}, {0, 0.2}},
		{{0.25, 0}, {0.244978663126864, 0}},
		{{0.25, 0.1}, {0.244978663126864, 0.094117647058824}},
		{{0.25, 0.2}, {0.244978663126864, 0.188235294117647}},
		{{0.5, 0}, {0.463647609000806, 0}},
		{{0.5, 0.1}, {0.463647609000806, 0.08}},
		{{0.5, 0.2}, {0.463647609000806, 0.16}},
		{{0.75, 0}, {0.643501108793284, 0}},
		{{0.75, 0.1}, {0.643501108793284, 0.064}},
		{{0.75, 0.2}, {0.643501108793284, 0.128}},
		{{1, 0}, {0.785398163397448, 0}},
		{{1, 0.1}, {0.785398163397448, 0.05}},
		{{1, 0.2}, {0.785398163397448, 0.1}},
		{{-10, 0}, {-1.47112767430373, 0}},
		{{-10, 1}, {-1.47112767430373, 0.00990099009901}},
		{{-100, 0}, {-1.56079666010823, 0}},
		{{-100, 1}, {-1.56079666010823, 9.99900009999e-05}},
		{{-100000000, 0}, {-1.5707963167949, 0}},
		{{-100000000, 3000}, {-1.5707963167949, 3e-13}},
		{{100000000, 3000}, {1.5707963167949, 3e-13}},
	}

	for i, the_case := range cases {
		res := Atan(the_case[0])

		if math.IsNaN(res.Value) {
			if !math.IsNaN(the_case[1].Value) {
				t.Fatalf("Test case %d failed: Acos(%f±%f) is %f±%f, got %f±%f",
					i, the_case[0].Value, the_case[0].Error, the_case[1].Value, the_case[1].Error, res.Value, res.Error)
			} else {
				continue
			}
		}

		if !almostEqual(res, the_case[1]) {
			t.Fatalf("Test case %d failed: Atan(%f±%f) is %f±%f, got %f±%f",
				i, the_case[0].Value, the_case[0].Error, the_case[1].Value, the_case[1].Error, res.Value, res.Error)
		}
	}
}

/*
func TestAtan2(t *testing.T) {
	cases := [][3]Uncertain{}

	for i, the_case := range cases {
		res := Atan2(the_case[0])
		if !almostEqual(res, the_case[1]) {
			t.Fatalf("Test case %d failed: Atan2(%f±%f, %f±%f) is %f±%f, got %f±%f",
				i, the_case[0].Value, the_case[0].Error, the_case[1].Value, the_case[1].Error, res.Value, res.Error)
		}
	}
}
*/
/*
func TestAtanh(t *testing.T) {
}
*/
