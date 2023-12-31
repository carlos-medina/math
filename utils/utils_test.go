package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		name           string
		n              int
		expectedResult int
		expectedError  string
	}{
		{
			name:           "Zero",
			n:              0,
			expectedResult: 1,
		},
		{
			name:           "One",
			n:              1,
			expectedResult: 1,
		},
		{
			name:           "Two",
			n:              2,
			expectedResult: 2,
		},
		{
			name:           "Three",
			n:              3,
			expectedResult: 6,
		},
		{
			name:           "Four",
			n:              4,
			expectedResult: 24,
		},
		{
			name:           "Negative number",
			n:              -1,
			expectedResult: 0,
			expectedError:  "Factorial is defined only for positive numbers; -1 is negative",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := Factorial(test.n)
			if test.expectedError != "" {
				assert.EqualError(t, err, test.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expectedResult, result)
			}
		})
	}
}
