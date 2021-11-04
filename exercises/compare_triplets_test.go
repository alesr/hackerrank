package exercises

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareTriplets(t *testing.T) {
	cases := []struct {
		name     string
		givenA   []int32
		givenB   []int32
		expected []int32
	}{
		{
			name:     "success",
			givenA:   []int32{5, 6, 7},
			givenB:   []int32{3, 6, 10},
			expected: []int32{1, 1},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			observed := CompareTriplets(tc.givenA, tc.givenB)
			assert.Equal(t, tc.expected, observed)
		})
	}
}
