package goasm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerated(t *testing.T) {
	tcs := []struct {
		x         []uint64
		y         []uint64
		expectedz []uint64
	}{
		{
			x:         []uint64{1, 2, 3, 4},
			y:         []uint64{1, 2, 3, 4},
			expectedz: []uint64{2, 4, 6, 8},
		},
		{
			x:         []uint64{1, 1, 1, 1, 2, 2, 2, 2},
			y:         []uint64{1, 1, 1, 1, 2, 2, 2, 2},
			expectedz: []uint64{2, 2, 2, 2, 4, 4, 4, 4},
		},
		{
			x:         []uint64{1, 2, 3, 4, 5},
			y:         []uint64{1, 2, 3, 4, 5},
			expectedz: []uint64{2, 4, 6, 8, 10},
		},
		{
			x:         []uint64{1},
			y:         []uint64{1},
			expectedz: []uint64{2},
		},
	}

	for i, tc := range tcs {
		tc := tc
		t.Run(fmt.Sprintf("test-%v", i), func(t *testing.T) {
			z := make([]uint64, len(tc.expectedz))
			Add(tc.x, tc.y, z)
			assert.Equal(t, tc.expectedz, z)
		})
	}
}
