package goasm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func FuzzGenerated(f *testing.F) {
	f.Fuzz(func(t *testing.T, xb, yb []byte) {
		x := prepare(xb)
		y := prepare(yb)

		if len(x) != len(y) {
			t.Skip()
		}

		z1 := make([]uint64, len(x))
		z2 := make([]uint64, len(x))

		referenceAdd(x, y, z1)
		Add(x, y, z2)

		assert.Equal(t, z1, z2)
	})
}

func referenceAdd(x, y, z []uint64) {
	for i := range x {
		z[i] = x[i] + y[i]
	}
}

func TestPrepare(t *testing.T) {
	assert.Len(t, prepare([]byte("000")), 1)
	assert.Len(t, prepare([]byte("0000")), 1)
	assert.Len(t, prepare([]byte("00000")), 2)
}

func prepare(data []byte) []uint64 {
	j := 0
	s := uint64(0)
	n := len(data) / 4
	if len(data)%4 != 0 {
		n += 1
	}
	res := make([]uint64, n)
	for i, v := range data {
		s <<= 8
		s += uint64(v)
		if i%4 == 3 {
			res[j] = s
			s = 0
			j++
		}
	}
	if j < len(res) {
		res[j] = s
	}
	return res
}
