package sum

import (
	"testing"

	sum "github.com/SondreH/IS-105/ica_01/Oppgave3/sum"
)

var sum_tests_float64 = []struct {
	n1       float64
	n2       float64
	expected float64
}{
	{1, 2, 3},
	{4, -5, -1},
	{126.5, 1.5, 128},
}

func TestSumfloat64(t *testing.T) {
	for _, v := range sum_tests_float64 {
		if val := sum.SumFloat64(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%f, %f) returned %f, expected %f", v.n1, v.n2, val, v.expected)
		}
	}
}
