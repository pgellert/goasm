package goasm

import (
	"math/rand"
	"testing"
)

func BenchmarkAdd10(b *testing.B)    { benchmarkAddN(10, b) }
func BenchmarkAdd100(b *testing.B)   { benchmarkAddN(100, b) }
func BenchmarkAdd1000(b *testing.B)  { benchmarkAddN(1000, b) }
func BenchmarkAdd10000(b *testing.B) { benchmarkAddN(10000, b) }

func BenchmarkBaselineAdd10(b *testing.B)    { benchmarkBaselineAddN(10, b) }
func BenchmarkBaselineAdd100(b *testing.B)   { benchmarkBaselineAddN(100, b) }
func BenchmarkBaselineAdd1000(b *testing.B)  { benchmarkBaselineAddN(1000, b) }
func BenchmarkBaselineAdd10000(b *testing.B) { benchmarkBaselineAddN(10000, b) }

func benchmarkAddN(l int, b *testing.B) {
	x, y, z := generateData(l)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		Add(x, y, z)
	}
}

func benchmarkBaselineAddN(l int, b *testing.B) {
	x, y, z := generateData(l)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		referenceAdd(x, y, z)
	}
}
func generateData(l int) (x, y, z []uint64) {
	x = makeRandom(l)
	y = makeRandom(l)
	z = make([]uint64, l)
	return
}

func makeRandom(l int) []uint64 {
	res := make([]uint64, l)
	for i := range res {
		res[i] = rand.Uint64()
	}
	return res
}
