package exercises

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiagonalDifference(t *testing.T) {
	cases := []struct {
		name     string
		given    [][]int32
		expected int32
	}{
		{
			name:     "success",
			given:    [][]int32{{1, 3}, {6, 5}},
			expected: 3,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			observed := DiagonalDifference(tc.given)
			assert.Equal(t, tc.expected, observed)
		})
	}
}
