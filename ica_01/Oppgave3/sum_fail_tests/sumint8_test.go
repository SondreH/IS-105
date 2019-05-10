package sum

import (
	"testing"

	sum "github.com/SondreH/IS-105/ica_01/Oppgave3/sum"
)

var sum_tests_int8 = []struct {
	n1       int8
	n2       int8
	expected int8
}{
	{1, 2, 3},
	{4, -5, -1},
	{1267, 1, 127},
}

func TestSumInt8(t *testing.T) {
	for _, v := range sum_tests_int8 {
		if val := sum.SumInt8(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}
