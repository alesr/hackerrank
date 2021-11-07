package exercises

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAVeryBigSum(t *testing.T) {
	cases := []struct {
		name     string
		given    []int64
		expected int64
	}{
		{
			name: "success",
			given: []int64{
				1000000001,
				1000000002,
				1000000003,
				1000000004,
				1000000005,
			},
			expected: 5000000015,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			observed := AVeryBigSum(tc.given)
			assert.Equal(t, tc.expected, observed)
		})
	}
}
