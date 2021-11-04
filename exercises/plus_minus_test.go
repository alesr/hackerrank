package exercises

import (
	"testing"

	"github.com/alesr/hackerrank/utilities"
	"github.com/stretchr/testify/assert"
)

func TestPlusMinus(t *testing.T) {
	cases := []struct {
		name     string
		given    []int32
		expected string
	}{
		{
			name:     "success",
			given:    []int32{-4, 3, -9, 0, 4, 1},
			expected: "0.500000\n0.333333\n0.166667\n",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			observed := utilities.CaptureOutput(func() {
				PlusMinus(tc.given)
			})

			assert.Equal(t, tc.expected, observed)
		})
	}
}
