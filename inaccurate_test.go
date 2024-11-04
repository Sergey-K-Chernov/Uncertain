package inaccurate

import (
	"math"
	"testing"
)

func AlmostEqual(a, b Inaccurate) bool {
	threshold := 0.000000000001
	return (math.Abs(a.Value-b.Value) <= threshold) && (math.Abs(a.Error-b.Error) <= threshold)
}

func TestAdd(t *testing.T) {
	v1 := []Inaccurate{
		{0, 0}, {1, 0.1}, {1, 0.1}, {-1, 0.1}, {-1, 0.1},
	}
	v2 := []Inaccurate{
		{0, 0}, {1, 0.1}, {-1, 0.1}, {1, 0.1}, {-1, 0.1},
	}

	res := []Inaccurate{
		{0, 0}, {2, 0.2}, {0, 0.2}, {0, 0.2}, {-2, 0.2},
	}

	if len(v1) != len(v2) || len(v1) != len(res) {
		t.Fatalf("Wrong contitions. Array lengths are not equal")
	}

	for i := range v1 {
		if !AlmostEqual(v1[i].Add(v2[i]), res[i]) {
			t.Fatalf("Test %d failed", i)
		}
	}
}

func TestSub(t *testing.T) {
	v1 := []Inaccurate{
		{0, 0}, {10, 1},
	}
	v2 := []Inaccurate{
		{0, 0}, {5, 0.5},
	}

	res := []Inaccurate{
		{0, 0}, {5, 1.5},
	}

	if len(v1) != len(v2) || len(v1) != len(res) {
		t.Fatalf("Wrong contitions. Array lengths are not equal")
	}

	for i := range v1 {
		if !AlmostEqual(v1[i].Sub(v2[i]), res[i]) {
			t.Fatalf("Test %d failed", i)
		}
	}
}

func TestMul(t *testing.T) {
	v1 := []Inaccurate{
		{0, 0}, {10, 1}, {-2, 0.1}, {100, 5}, {0, 1},
	}
	v2 := []Inaccurate{
		{0, 0}, {5, 0.5}, {2, 0.1}, {-200, 40}, {10, 1},
	}

	res := []Inaccurate{
		{0, 0}, {50, 10}, {-4, 0.4}, {-20000, 5000}, {0, 11},
	}

	if len(v1) != len(v2) || len(v1) != len(res) {
		t.Fatalf("Wrong contitions. Array lengths are not equal")
	}

	for i := range v1 {
		if !AlmostEqual(v1[i].Mul(v2[i]), res[i]) {
			t.Fatalf("Test %d failed", i)
		}
	}
}

func TestMul2(t *testing.T) {
	v1 := []Inaccurate{
		{0, 0}, {10, 1}, {-2, 0.1}, {100, 5}, {0, 1},
	}
	v2 := []Inaccurate{
		{0, 0}, {5, 0.5}, {2, 0.1}, {-200, 40}, {10, 1},
	}

	res := []Inaccurate{
		{0, 0}, {50, 10}, {-4, 0.4}, {-20000, 5000}, {0, 11},
	}

	if len(v1) != len(v2) || len(v1) != len(res) {
		t.Fatalf("Wrong contitions. Array lengths are not equal")
	}

	for i := range v1 {
		if !AlmostEqual(v1[i].mul(v2[i]), res[i]) {
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
		if !AlmostEqual(v1[i].Div(v2[i]), res[i]) {
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
		if !AlmostEqual(v1[i].div(v2[i]), res[i]) {
			t.Fatalf("Test %d failed", i)
		}
	}
}
