package inaccurate

import "math"

type Inaccurate struct {
	Value float64
	Error float64
}

func (v1 Inaccurate) Add(v2 Inaccurate) (sum Inaccurate) {
	sum.Value = v1.Value + v2.Value
	sum.Error = v1.Error + v2.Error
	return
}

func (v1 Inaccurate) Sub(v2 Inaccurate) (diff Inaccurate) {
	diff.Value = v1.Value - v2.Value
	diff.Error = v1.Error + v2.Error
	return
}

func (v1 Inaccurate) Mul(v2 Inaccurate) (product Inaccurate) {
	if v1.Value == 0 || v2.Value == 0 {
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

func (v1 Inaccurate) mul(v2 Inaccurate) (product Inaccurate) {
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

func (v1 Inaccurate) Div(v2 Inaccurate) (quotient Inaccurate) {
	if v1.Value == 0 {
		return v1.div(v2)
	}

	quotient.Value = v1.Value / v2.Value

	relError := 0.0
	relError += v1.Error / math.Abs(v1.Value)
	relError += v2.Error / math.Abs(v2.Value) // No  need to check for division by zero: in this case function will fail in the first line

	quotient.Error = relError * math.Abs(quotient.Value)
	return
}

func (v1 Inaccurate) div(v2 Inaccurate) (quotient Inaccurate) {
	rel := v2.Error / v2.Value
	v2.Value = 1 / v2.Value
	v2.Error = v2.Value * rel

	return v1.mul(v2)
}
