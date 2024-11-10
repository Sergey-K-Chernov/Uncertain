package uncertain

import (
	"math"
	"testing"
)

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

// TODO
//func TestCbrt(t *testing.T)

// TODO
//func Log2(t *testing.T) {

// TODO
//func Pow(t *testing.T) {

func TestSqrt(t *testing.T) {
	cases := [][2]Uncertain{
		{{-1, 0}, {math.NaN(), math.NaN()}},
		{{-1, 0.1}, {math.NaN(), math.NaN()}},
		{{0, 0}, {0, 0}},
		{{0, 0.1}, {0, 0.316227766016838}},
		{{1, 0}, {1, 0}},
		{{1, 0.1}, {1, 0.05}},
		{{0.64, 0}, {0.8, 0}},
		{{0.64, 0.4}, {0.8, 0.25}},
		{{0.81, 0.9}, {0.9, 0.5}},
		{{1.5, 0.3}, {1.22474487139159, 0.122474487139159}},
		{{2, 0.2}, {1.4142135623731, 0.070710678118655}},
		{{3, 0.2}, {1.73205080756888, 0.057735026918963}},
		{{4, 0.4}, {2, 0.1}},
		{{9, 0.3}, {3, 0.05}},
		{{100, 10}, {10, 0.5}},
		{{1024, 128}, {32, 2}},
		{{65536, 512}, {256, 1}},
		{{1000000, 1000}, {1000, 0.5}},
	}

	for i, the_case := range cases {
		res := Sqrt(the_case[0])

		if math.IsNaN(res.Value) {
			if !math.IsNaN(the_case[1].Value) {
				t.Fatalf("Test case %d failed: Sqrt(%f±%f) is %f±%f, got %f±%f",
					i, the_case[0].Value, the_case[0].Error, the_case[1].Value, the_case[1].Error, res.Value, res.Error)
			} else {
				continue
			}
		}

		if !almostEqual(res, the_case[1]) {
			t.Fatalf("Test case %d failed: Sqrt(%f±%f) is %f±%f, got %f±%f",
				i, the_case[0].Value, the_case[0].Error, the_case[1].Value, the_case[1].Error, res.Value, res.Error)
		}
	}
}
