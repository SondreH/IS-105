package sum

import (
	"testing"

	sum "github.com/SondreH/IS-105/ica_01/Oppgave3/sum"
)

var sum_tests_int32 = []struct {
	n1       int32
	n2       int32
	expected int32
}{
	{1, 2, 3},
	{4, -5, -1},
	{126, 1, 127},
}

func TestSumInt32(t *testing.T) {
	for _, v := range sum_tests_int32 {
		if val := sum.SumInt32(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}
