package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test number of generated rolls matches dice flag
func TestRollsLength(t *testing.T) {
	sides := 6
	dice := 8
	rolls := GenerateRolls(sides, dice)
	assert.Equal(t, dice, len(rolls))
}

// Test various integer mixes for expected errors
func TestInputValidation(t *testing.T) {
  var sideerr, diceerr error

	tests := []struct {
		testname      string
		sides         int
		dice          int
		expectedError []error
	}{
		{
			"test no errors",
			6,
			4,
			[]error{},
		},
		{
			"test sides error",
			-6,
			6,
			[]error{sideerr},
		},
		{
			"test dice error",
			6,
			-6,
			[]error{diceerr},
		},
		{
			"test both errors",
			-6,
			-6,
			[]error{sideerr, diceerr},
		},
	}

	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			expected := test.expectedError
			actual := ValidateInputs(test.sides, test.dice)
			assert.Equal(t, expected, actual)
		})
	}
}
