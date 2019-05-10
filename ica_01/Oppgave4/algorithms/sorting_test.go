package algorithms

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

//For å sjekke benchmarkene:
//Go test -bench=.

// https://golang.org/doc/effective_go.html#init
func init() {
	seed := time.Now().Unix()
	rand.Seed(seed)
}

func perm(n int) (out []int) {
	for _, v := range rand.Perm(n) {
		out = append(out, v)
	}
	return
}

// Skriv "benchmark"-tester for benchmarkBSortModified funksjonen
// Skriv en ny testfunksjon benchmarkBSortModified

func BenchmarkBSort100(b *testing.B) {
	benchmarkBSort(100, b)
}

func BenchmarkBSort1000(b *testing.B) {
	benchmarkBSort(1000, b)
}

func BenchmarkBSort10000(b *testing.B) {
	benchmarkBSort(10000, b)
}

func benchmarkBSort(i int, b *testing.B) {
	for j := 0; j < b.N; j++ {
		b.StopTimer()
		values := perm(i)
		b.StartTimer()
		Bubble_sort(values)
	}
}

func benchmarkBSortModifisert(i int, b *testing.B) {
	for j := 0; j < b.N; j++ {
		b.StopTimer()
		values := perm(i)
		b.StartTimer()
		Bubble_sort_modified(values)
	}
}

func BenchmarkBSortMod100(b *testing.B) {
	benchmarkBSortModifisert(100, b)
}

func BenchmarkBSortMod1000(b *testing.B) {
	benchmarkBSortModifisert(1000, b)
}

func BenchmarkBSortMod10000(b *testing.B) {
	benchmarkBSortModifisert(10000, b)
}

// Implementasjon av testfunksjoner for Quicksort algoritmen
func TestQSort(t *testing.T) {
	values := []int{9, 1, 20, 3, 6, 7}
	expected := []int{1, 3, 6, 7, 9, 20}

	QSort(values)

	if !reflect.DeepEqual(values, expected) {
		t.Fatalf("expected %d, actual is %d", 1, values[0])
	}
}
func BenchmarkQSort100(b *testing.B) {
	benchmarkQSort(100, b)
}

func BenchmarkQSort1000(b *testing.B) {
	benchmarkQSort(1000, b)
}

func BenchmarkQSort10000(b *testing.B) {
	benchmarkQSort(10000, b)
}

func benchmarkQSort(i int, b *testing.B) {
	for j := 0; j < b.N; j++ {
		b.StopTimer()
		values := perm(i)
		b.StartTimer()
		QSort(values)
	}
}
