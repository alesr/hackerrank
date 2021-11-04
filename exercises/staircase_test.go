package exercises

import (
	"testing"

	"github.com/alesr/hackerrank/utilities"
	"github.com/stretchr/testify/assert"
)

func TestStaircase(t *testing.T) {
	cases := []struct {
		name     string
		given    int32
		expected string
	}{
		{
			name:     "success",
			given:    5,
			expected: "    #\n   ##\n  ###\n ####\n#####\n",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			observed := utilities.CaptureOutput(func() {
				Staircase(tc.given)
			})

			assert.Equal(t, tc.expected, observed)
		})
	}
}
