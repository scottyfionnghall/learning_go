package math

import "testing"

type testpair struct {
	values  []float64
	avarage float64
	min     float64
	max     float64
}

var test = []testpair{
	{[]float64{1, 2}, 1.5, 1, 2},
	{[]float64{1, 1, 1, 1, 1, 1}, 1, 1, 1},
	{[]float64{-1, 1}, 0, -1, 1},
	{[]float64{}, 0, 0, 0},
}

func TestAvarage(t *testing.T) {
	for _, pair := range test {
		v := Avarage(pair.values)
		if v != pair.avarage {
			t.Error(
				"For", pair.values,
				"expected", pair.avarage,
				"got", v,
			)
		}
	}
}

func TestMin(t *testing.T) {
	for _, pair := range test {
		v := Min(pair.values)
		if v != pair.min {
			t.Error(
				"For", pair.values,
				"expected", pair.min,
				"got", v,
			)
		}
	}
}

func TestMax(t *testing.T) {
	for _, pair := range test {
		v := Max(pair.values)
		if v != pair.max {
			t.Error(
				"For", pair.values,
				"expected", pair.max,
				"got", v,
			)
		}
	}
}
