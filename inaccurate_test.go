package inaccurate

import (
	"math"
	"testing"
)

func almostEqual(a, b Inaccurate) bool {
	threshold := 0.0000000000001
	return (math.Abs(a.Value-b.Value) <= threshold) && (math.Abs(a.Error-b.Error) <= threshold)
}

func TestAdd(t *testing.T) {
	v1 := []Inaccurate{
		{0, 0}, {0, 3}, {3, 3}, {1, 0.1}, {1, 0.1}, {-1, 0.1}, {-1, 0.1},
	}
	v2 := []Inaccurate{
		{0, 0}, {3, 0}, {0, 3}, {1, 0.1}, {-1, 0.1}, {1, 0.1}, {-1, 0.1},
	}

	res := []Inaccurate{
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
	v1 := []Inaccurate{
		{0, 0}, {10, 1}, {-100, 10}, {-100, 10}, {200, 50}, {200, 50},
	}
	v2 := []Inaccurate{
		{0, 0}, {5, 0.5}, {100, 10}, {-100, 10}, {300, 30}, {-300, 30},
	}

	res := []Inaccurate{
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
	v1 := []Inaccurate{
		{7, 0}, {0, 0}, {10, 1}, {-2, 0.1}, {100, 5}, {0, 1}, {10, 1},
	}
	v2 := []Inaccurate{
		{6, 0}, {0, 0}, {5, 0.5}, {2, 0.1}, {-200, 40}, {10, 1}, {0, 1},
	}

	res := []Inaccurate{
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
	v1 := []Inaccurate{
		{7, 0}, {0, 0}, {10, 1}, {-2, 0.1}, {100, 5}, {0, 1}, {10, 1},
	}
	v2 := []Inaccurate{
		{6, 0}, {0, 0}, {5, 0.5}, {2, 0.1}, {-200, 40}, {10, 1}, {0, 1},
	}

	res := []Inaccurate{
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
	v1 := []Inaccurate{
		{0, 10}, {0, 0}, {0, 10}, {10, 1}, {500, 25}, {-300, 20},
	}
	v2 := []Inaccurate{
		{1, 0}, {1, 0}, {1, 0.1}, {10, 1}, {-10, 5}, {3, 0.5},
	}

	res := []Inaccurate{
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
	v1 := []Inaccurate{
		{0, 10}, {0, 0}, {0, 10}, {10, 1}, {500, 25}, {-300, 20},
	}
	v2 := []Inaccurate{
		{1, 0}, {1, 0}, {1, 0.1}, {10, 1}, {-10, 5}, {3, 0.5},
	}

	res := []Inaccurate{
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
	cases := [][2]Inaccurate{
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
		if !almostEqual(res, the_case[1]) {
			t.Fatalf("Test case %d failed: Acos(%f±%f) is %f±%f, got %f±%f",
				i, the_case[0].Value, the_case[0].Error, the_case[1].Value, the_case[1].Error, res.Value, res.Error)
		}
	}
}

func TestAcosh(t *testing.T) {

}

func TestAsin(t *testing.T) {

}

func TestAsinh(t *testing.T) {

}

func TestAtan(t *testing.T) {
}

func TestAtan2(t *testing.T) {
}

func TestAtanh(t *testing.T) {
}
