package sum

import (
	"testing"

	sum "github.com/SondreH/IS-105/ica_01/Oppgave3/sum"
)

var sum_tests_uint32 = []struct {
	n1       uint32
	n2       uint32
	expected uint32
}{
	{1, 4, 5},
	{4, 5, 9},
	{126, 1, 127},
}

func TestSumuint32(t *testing.T) {
	for _, v := range sum_tests_uint32 {
		if val := sum.SumUint32(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}
