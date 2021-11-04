package exercises

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleArraySum(t *testing.T) {
	cases := []struct {
		name     string
		given    []int32
		expected int32
	}{
		{
			name:     "success",
			given:    []int32{1, 2, 3, 4, 10, 11},
			expected: 31,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			observed := simpleArraySum(tc.given)
			assert.Equal(t, tc.expected, observed)
		})
	}
}
