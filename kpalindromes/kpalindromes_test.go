package kpalindromes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKpalindromes(t *testing.T) {
	testCases := []struct {
		description    string
		k              int
		s              string
		expectedResult bool
	}{
		{
			description:    "Happy path",
			k:              3,
			s:              "aabbaa",
			expectedResult: true,
		},
		{
			description:    "Happy path",
			k:              2,
			s:              "aabbc",
			expectedResult: true,
		},
		{
			description:    "Sad path - k too high",
			k:              3,
			s:              "aabbc",
			expectedResult: false,
		},
		{
			description:    "No duplicates",
			k:              3,
			s:              "abcdefg",
			expectedResult: false,
		},
		{
			description:    "k == 0",
			k:              0,
			s:              "abcdefg",
			expectedResult: true,
		},
		{
			description:    "n <= 1 && k <= 1",
			k:              1,
			s:              "a",
			expectedResult: true,
		},
		{
			description:    "k > n",
			k:              2,
			s:              "a",
			expectedResult: false,
		},
		{
			description:    "oddCharacterCount > 1 && n > 2",
			k:              1,
			s:              "aaaaaaaaaaaabc",
			expectedResult: false,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.description, func(t *testing.T) {
			asserter := assert.New(t)
			result := Kpalindromes(tC.k, tC.s)
			asserter.Equal(tC.expectedResult, result)
		})
	}
}
