package uncertain

import (
	"math"
	"testing"
)

func almostEqual(a, b Uncertain) bool {
	if a.Value == math.Inf(1) || b.Value == math.Inf(1) || a.Value == math.Inf(-1) || b.Value == math.Inf(-1) {
		return (a.Value == b.Value && a.Error == b.Error)
	}
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
		{0, 10}, {0, 0}, {0, 10}, {10, 1}, {500, 25}, {-300, 20}, {10, 3}, {-10, 3},
	}
	v2 := []Uncertain{
		{1, 0}, {1, 0}, {1, 0.1}, {10, 1}, {-10, 5}, {3, 0.5}, {0, 4}, {0, 4},
	}

	res := []Uncertain{
		{0, 10}, {0, 0}, {0, 11}, {1, 0.2}, {-50, 27.5}, {-100, 23.3333333333333}, {math.Inf(1), math.Inf(1)}, {math.Inf(-1), math.Inf(1)},
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
		{{-1, 0}, {math.Pi, 0}},
		{{-1, 0.1}, {math.Pi, 0.451026811796262}},
		{{-1, 0.2}, {math.Pi, 0.643501108793284}},
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
		{{-1, 0}, {-math.Pi / 4, 0}},
		{{-1, 0.1}, {-math.Pi / 4, 0.05}},
		{{-1, 0.2}, {-math.Pi / 4, 0.1}},
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
		{{1, 0}, {math.Pi / 4, 0}},
		{{1, 0.1}, {math.Pi / 4, 0.05}},
		{{1, 0.2}, {math.Pi / 4, 0.1}},
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
				t.Fatalf("Test case %d failed: Atan(%f±%f) is %f±%f, got %f±%f",
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

func TestAtan2(t *testing.T) {
	negative_zero := math.Copysign(0.0, -1)

	cases := [][3]Uncertain{
		{{0, 0}, {math.NaN(), 0}, {math.NaN(), 0}}, //	Atan2({y, ey}, {NaN, ex}) = {NaN, _}
		{{math.NaN(), 0}, {0, 0}, {math.NaN(), 0}}, //	Atan2({NaN, ey}, {x, ex}) = {NaN, _}

		{{0, 0}, {10, 1}, {0, 0}}, //	Atan2({0, 0}, {x>=0, ex}) = {0, 0}
		{{0, 1}, {10, 1}, {0, 0.110657221173896}},
		{{negative_zero, 0}, {10, 1}, {-0, 0}}, //	Atan2({-0, 0}, {x>=0, ex}) = {-0, 0}
		{{negative_zero, 1}, {10, 1}, {-0, 0.110657221173896}},

		{{0, 0}, {-10, 1}, {math.Pi, 0}}, //	Atan2({0, 0}, {x<=-0, ex}) = {Pi, 0}
		{{0, 1}, {-10, 1}, {math.Pi, 3.05093276638905}},

		{{negative_zero, 0}, {-10, 1}, {-math.Pi, 0}}, //	Atan2({-0, 0}, {x<=-0, ex}) = {-Pi, 0}
		{{negative_zero, 1}, {-10, 1}, {-math.Pi, 3.05093276638905}},

		{{10, 0}, {0, 0}, {math.Pi / 2, 0}}, //	Atan2({y>0, ey}, {0, 0}) = {+Pi/2, 0}
		{{10, 1}, {0, 1}, {math.Pi / 2, 0.110657221173896}},

		{{-10, 0}, {0, 0}, {-math.Pi / 2, 0}}, //	Atan2({y<0, ey}, {0, 0}) = {-Pi/2, 0}
		{{-10, 1}, {0, 1}, {-math.Pi / 2, 0.110657221173896}},

		{{math.Inf(1), 0}, {math.Inf(1), 0}, {math.Pi / 4, 0}},   //	Atan2({+Inf, ey}, {+Inf, ex}) = {+Pi/4, 0}
		{{math.Inf(1), 10}, {math.Inf(1), 10}, {math.Pi / 4, 0}}, //	Atan2({+Inf, ey}, {+Inf, ex}) = {+Pi/4, 0}

		{{math.Inf(-1), 0}, {math.Inf(1), 0}, {-math.Pi / 4, 0}},   //	Atan2({-Inf, ey}, {+Inf, ex}) = {-Pi/4, 0}
		{{math.Inf(-1), 10}, {math.Inf(1), 10}, {-math.Pi / 4, 0}}, //	Atan2({-Inf, ey}, {+Inf, ex}) = {-Pi/4, 0}

		{{math.Inf(1), 0}, {math.Inf(-1), 0}, {3 * math.Pi / 4, 0}},   //	Atan2({+Inf, ey}, {-Inf, ex}) = {3Pi/4, 0}
		{{math.Inf(1), 10}, {math.Inf(-1), 10}, {3 * math.Pi / 4, 0}}, //	Atan2({+Inf, ey}, {-Inf, ex}) = {3Pi/4, 0}

		{{math.Inf(-1), 0}, {math.Inf(-1), 0}, {-3 * math.Pi / 4, 0}},   //	Atan2({-Inf, ey}, {-Inf, ex}) = {-3Pi/4, 0}
		{{math.Inf(-1), 10}, {math.Inf(-1), 10}, {-3 * math.Pi / 4, 0}}, //	Atan2({-Inf, ey}, {-Inf, ex}) = {-3Pi/4, 0}

		{{0, 0}, {math.Inf(1), 0}, {0, 0}},  //	Atan2({y, ey}, {+Inf, ex}) = {0, 0}
		{{0, 1}, {math.Inf(1), 10}, {0, 0}}, //	Atan2({y, ey}, {+Inf, ex}) = {0, 0}

		{{5, 0}, {math.Inf(1), 0}, {0, 0}},   //	Atan2({y, ey}, {+Inf, ex}) = {0, 0}
		{{-5, 1}, {math.Inf(1), 10}, {0, 0}}, //	Atan2({y, ey}, {+Inf, ex}) = {0, 0}

		{{100, 0}, {math.Inf(-1), 0}, {math.Pi, 0}},   //	Atan2({y>0, ey}, {-Inf, ex}) = {+Pi, 0}
		{{100, 10}, {math.Inf(-1), 10}, {math.Pi, 0}}, //	Atan2({y>0, ey}, {-Inf, ex}) = {+Pi, 0}

		{{-100, 0}, {math.Inf(-1), 0}, {-math.Pi, 0}},   //	Atan2({y<0, ey}, {-Inf, ex}) = {-Pi, 0}
		{{-100, 10}, {math.Inf(-1), 10}, {-math.Pi, 0}}, //	Atan2({y<0, ey}, {-Inf, ex}) = {-Pi, 0}

		{{math.Inf(1), 0}, {100, 0}, {math.Pi / 2, 0}},   //	Atan2({+Inf, ey}, {x, ex}) = {+Pi/2, 0}
		{{math.Inf(1), 10}, {100, 10}, {math.Pi / 2, 0}}, //	Atan2({+Inf, ey}, {x, ex}) = {+Pi/2, 0}

		{{math.Inf(-1), 0}, {100, 0}, {-math.Pi / 2, 0}},   //	Atan2({-Inf, ey}, {x, ex}) = {-Pi/2, 0}
		{{math.Inf(-1), 10}, {100, 10}, {-math.Pi / 2, 0}}, //	Atan2({-Inf, ey}, {x, ex}) = {-Pi/2, 0}

		{{300, 0}, {300, 0}, {math.Pi / 4, 0}},
		{{300, 10}, {300, 10}, {math.Pi / 4, 0.033320995878247}},
		{{300, 10}, {300, 0}, {math.Pi / 4, 0.016669752057205}},
		{{300, 0}, {300, 10}, {math.Pi / 4, 0.016669752057205}},

		{{500, 0}, {-500, 0}, {3 * math.Pi / 4, 0}},
		{{500, 50}, {-500, 0}, {3 * math.Pi / 4, 0.050083082443963}},
		{{500, 0}, {-500, 50}, {3 * math.Pi / 4, 0.050083082443963}},
		{{500, 50}, {-500, 50}, {3 * math.Pi / 4, 0.099668652491162}},

		{{-200, 0}, {200, 0}, {-math.Pi / 4, 0}},
		{{-200, 30}, {200, 0}, {-math.Pi / 4, 0.075279336441899}},
		{{-200, 0}, {200, 30}, {-math.Pi / 4, 0.0752793364419}},
		{{-200, 30}, {200, 30}, {-math.Pi / 4, 0.148889947609497}},

		{{-400, 0}, {-400, 0}, {-3 * math.Pi / 4, 0}},
		{{-400, 70}, {-400, 0}, {-3 * math.Pi / 4, 0.087942466973428}},
		{{-400, 0}, {-400, 70}, {-3 * math.Pi / 4, 0.087942466973429}},
		{{-400, 70}, {-400, 70}, {-3 * math.Pi / 4, 0.173245666452365}},

		{{300, 0}, {300 * math.Sqrt(3), 0}, {math.Pi / 6, 0}},
		{{300, 30}, {300 * math.Sqrt(3), 70}, {math.Pi / 6, 0.101865701678475}},
	}

	for i, the_case := range cases {
		res := Atan2(the_case[0], the_case[1])

		if math.IsNaN(res.Value) {
			if !math.IsNaN(the_case[2].Value) {
				t.Fatalf("Test case %d failed: Acos(%f±%f) is %f±%f, got %f±%f",
					i, the_case[0].Value, the_case[0].Error, the_case[1].Value, the_case[1].Error, res.Value, res.Error)
			} else {
				continue
			}
		}

		if !almostEqual(res, the_case[2]) {
			t.Fatalf("Test case %d failed: Atan2(%f±%f, %f±%f) is %f±%f, got %f±%f",
				i, the_case[0].Value, the_case[0].Error, the_case[1].Value, the_case[1].Error, the_case[2].Value, the_case[2].Error, res.Value, res.Error)
		}
	}
}

/*
// TODO
func TestAtanh(t *testing.T) {
}
*/

// TODO
//func TestCbrt(t *testing.T)

func TestCos(t *testing.T) {
	cases := [][2]Uncertain{
		{{math.Inf(-1), 10000}, {math.NaN(), 0}},
		{{math.Inf(1), 10000}, {math.NaN(), 0}},
		{{math.NaN(), 10000}, {math.NaN(), 0}},

		{{0, 0}, {1, 0}},
		{{0, 0.05}, {1, 0}},
		{{0.1, 0}, {0.995004165278026, 0}},
		{{0.1, 0.01}, {0.995004165278026, 0.000998334166468}},
		{{-0.1, 0}, {0.995004165278026, 0}},
		{{-0.1, 0.01}, {0.995004165278026, 0.000998334166468}},
		{{1, 0.05}, {0.54030230586814, 0.042073549240395}},
		{{-1, 0.05}, {0.54030230586814, 0.042073549240395}},
		{{math.Pi / 6, 0}, {0.866025403784439, 0}},
		{{math.Pi / 6, 0.1}, {0.866025403784439, 0.05}},
		{{-math.Pi / 6, 0}, {0.866025403784439, 0}},
		{{-math.Pi / 6, 0.1}, {0.866025403784439, 0.05}},
		{{math.Pi / 4, 0}, {0.707106781186548, 0}},
		{{math.Pi / 4, 0.1}, {0.707106781186548, 0.070710678118655}},
		{{-math.Pi / 4, 0}, {0.707106781186548, 0}},
		{{-math.Pi / 4, 0.1}, {0.707106781186548, 0.070710678118655}},
		{{math.Pi / 3, 0}, {0.5, 0}},
		{{math.Pi / 3, 0.1}, {0.5, 0.086602540378444}},
		{{-math.Pi / 3, 0}, {0.5, 0}},
		{{-math.Pi / 3, 0.1}, {0.5, 0.086602540378444}},
		{{math.Pi / 2, 0}, {6.12323399573677e-17, 0}},
		{{math.Pi / 2, 0.1}, {6.12323399573677e-17, 0.1}},
		{{-math.Pi / 2, 0}, {6.12323399573677e-17, 0}},
		{{-math.Pi / 2, 0.1}, {6.12323399573677e-17, 0.1}},

		{{3 * math.Pi / 4, 0}, {-0.707106781186547, 0}},
		{{3 * math.Pi / 4, 0.1}, {-0.707106781186547, 0.070710678118655}},
		{{-3 * math.Pi / 4, 0}, {-0.707106781186547, 0}},
		{{-3 * math.Pi / 4, 0.1}, {-0.707106781186547, 0.070710678118655}},
		{{math.Pi, 0}, {-1, 0}},
		{{math.Pi, 0.1}, {-1, 1.22464679914735e-17}},
		{{-math.Pi, 0}, {-1, 0}},
		{{-math.Pi, 0.1}, {-1, 1.22464679914735e-17}},

		{{3 * math.Pi / 2, 0}, {-1.83697019872103e-16, 0}},
		{{3 * math.Pi / 2, 0.1}, {-1.83697019872103e-16, 0.1}},
		{{-3 * math.Pi / 2, 0}, {-1.83697019872103e-16, 0}},
		{{-3 * math.Pi / 2, 0.1}, {-1.83697019872103e-16, 0.1}},
		{{2 * math.Pi, 0}, {1, 0}},
		{{2 * math.Pi, 0.1}, {1, 2.44929359829471e-17}},
		{{3 * math.Pi, 0}, {-1, 0}},
		{{3 * math.Pi, 0.1}, {-1, 3.67394039744206e-17}},
		{{4 * math.Pi, 0.1}, {1, 4.89858719658941e-17}},
		{{-4 * math.Pi, 0.1}, {1, 4.89858719658941e-17}},
	}

	for i, the_case := range cases {
		res := Cos(the_case[0])

		if math.IsNaN(res.Value) {
			if !math.IsNaN(the_case[1].Value) {
				t.Fatalf("Test case %d failed: Cos(%f±%f) is %f±%f, got %f±%f",
					i, the_case[0].Value, the_case[0].Error, the_case[1].Value, the_case[1].Error, res.Value, res.Error)
			} else {
				continue
			}
		}

		if !almostEqual(res, the_case[1]) {
			t.Fatalf("Test case %d failed: Cos(%f±%f) is %f±%f, got %f±%f",
				i, the_case[0].Value, the_case[0].Error, the_case[1].Value, the_case[1].Error, res.Value, res.Error)
		}
	}
}

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
//func Log2(t *testing.T) {

// TODO
//func Pow(t *testing.T) {

// TODO
//func Pow10(t *testing.T) {

func TestSin(t *testing.T) {
	negative_zero := math.Copysign(0.0, -1)

	cases := [][2]Uncertain{
		{{0, 0}, {0, 0}},
		{{negative_zero, 0}, {negative_zero, 0}},
		{{math.Inf(-1), 10000}, {math.NaN(), 0}},
		{{math.Inf(1), 10000}, {math.NaN(), 0}},
		{{math.NaN(), 10000}, {math.NaN(), 0}},

		{{0, 0.05}, {0, 0.05}},
		{{0.1, 0}, {0.099833416646828, 0}},
		{{0.1, 0.01}, {0.099833416646828, 0.00995004165278}},
		{{-0.1, 0}, {-0.099833416646828, 0}},
		{{-0.1, 0.01}, {-0.099833416646828, 0.00995004165278}},
		{{1, 0.05}, {0.841470984807896, 0.027015115293407}},
		{{-1, 0.05}, {-0.841470984807896, 0.027015115293407}},
		{{math.Pi / 6, 0}, {0.5, 0}},
		{{math.Pi / 6, 0.1}, {0.5, 0.086602540378444}},
		{{math.Pi / -6, 0}, {-0.5, 0}},
		{{math.Pi / -6, 0.1}, {-0.5, 0.086602540378444}},
		{{math.Pi / 4, 0}, {0.707106781186548, 0}},
		{{math.Pi / 4, 0.1}, {0.707106781186548, 0.070710678118655}},
		{{math.Pi / -4, 0}, {-0.707106781186548, 0}},
		{{math.Pi / -4, 0.1}, {-0.707106781186548, 0.070710678118655}},
		{{math.Pi / 3, 0}, {0.866025403784439, 0}},
		{{math.Pi / 3, 0.1}, {0.866025403784439, 0.05}},
		{{math.Pi / -3, 0}, {-0.866025403784439, 0}},
		{{math.Pi / -3, 0.1}, {-0.866025403784439, 0.05}},
		{{math.Pi / 2, 0}, {1, 0}},
		{{math.Pi / 2, 0.1}, {1, 6.12323399573677e-18}},
		{{math.Pi / -2, 0}, {-1, 0}},
		{{math.Pi / -2, 0.1}, {-1, 6.12323399573677e-18}},
		{{3 * math.Pi / 4, 0}, {0.707106781186548, 0}},
		{{3 * math.Pi / 4, 0.1}, {0.707106781186548, 0.070710678118655}},
		{{-3 * math.Pi / 4, 0}, {-0.707106781186548, 0}},
		{{-3 * math.Pi / 4, 0.1}, {-0.707106781186548, 0.070710678118655}},
		{{math.Pi, 0}, {1.22464679914735e-16, 0}},
		{{math.Pi, 0.1}, {1.22464679914735e-16, 0.1}},
		{{-math.Pi, 0}, {-1.22464679914735e-16, 0}},
		{{-math.Pi, 0.1}, {-1.22464679914735e-16, 0.1}},
		{{3 * math.Pi / 2, 0}, {-1, 0}},
		{{3 * math.Pi / 2, 0.1}, {-1, 1.83697019872103e-17}},
		{{-3 * math.Pi / 2, 0}, {1, 0}},
		{{-3 * math.Pi / 2, 0.1}, {1, 1.83697019872103e-17}},
		{{2 * math.Pi, 0}, {-2.44929359829471e-16, 0}},
		{{2 * math.Pi, 0.1}, {-2.44929359829471e-16, 0.1}},
		{{3 * math.Pi, 0}, {3.67394039744206e-16, 0}},
		{{3 * math.Pi, 0.1}, {3.67394039744206e-16, 0.1}},
		{{4 * math.Pi, 0.1}, {-4.89858719658941e-16, 0.1}},
		{{-4 * math.Pi, 0.1}, {4.89858719658941e-16, 0.1}},
	}

	for i, the_case := range cases {
		res := Sin(the_case[0])

		if math.IsNaN(res.Value) {
			if !math.IsNaN(the_case[1].Value) {
				t.Fatalf("Test case %d failed: Sin(%f±%f) is %f±%f, got %f±%f",
					i, the_case[0].Value, the_case[0].Error, the_case[1].Value, the_case[1].Error, res.Value, res.Error)
			} else {
				continue
			}
		}

		if !almostEqual(res, the_case[1]) {
			t.Fatalf("Test case %d failed: Sin(%f±%f) is %f±%f, got %f±%f",
				i, the_case[0].Value, the_case[0].Error, the_case[1].Value, the_case[1].Error, res.Value, res.Error)
		}
	}
}
