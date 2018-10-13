package pyrand_test

import (
	"testing"

	"github.com/rowanseymour/pyrand"
)

// test values generated using CPython 3.7.0
var rngTests = []struct {
	seed     int64
	sequence []float64
}{
	{0, []float64{0.8444218515250481, 0.7579544029403025, 0.420571580830845, 0.25891675029296335, 0.5112747213686085}},
	{1234, []float64{0.9664535356921388, 0.4407325991753527, 0.007491470058587191, 0.9109759624491242, 0.939268997363764}},
	{-1234, []float64{0.9664535356921388, 0.4407325991753527, 0.007491470058587191, 0.9109759624491242, 0.939268997363764}},
	{12345678, []float64{0.7202671550185803, 0.6330310001166692, 0.22877255649315598, 0.25254569034434393, 0.6060686820396118}},
	{23523623435353553, []float64{0.7667782846525043, 0.2620460079415977, 0.9385746408320916, 0.10965138305592881, 0.9750957142925043}},
}

func TestPythonRand(t *testing.T) {
	for _, tc := range rngTests {
		r := &pyrand.PythonRand{}
		r.Seed(tc.seed)

		for _, expected := range tc.sequence {
			actual := r.Random()
			if expected != actual {
				t.Errorf("expected %f got %f", expected, actual)
			}
		}
	}
}
