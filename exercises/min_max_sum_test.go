package exercises

import (
	"testing"

	"github.com/alesr/hackerrank/utilities"
	"github.com/stretchr/testify/assert"
)

func TestMinMaxSum(t *testing.T) {
	cases := []struct{
		name string
		given []int32
		expected string
	}{
		{
			name: "success",
			given: []int32{1, 3, 5, 7, 9},
			expected: "16 24",
		},
	}
	
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			observed := utilities.CaptureOutput(func() {
				MinMaxSum(tc.given)
			})

			assert.Equal(t, tc.expected, observed)	
		})
	}
}