package main

import (
	"fmt"
	"testing"
	"urls/pkg/hash"

	"github.com/stretchr/testify/assert"
)


func TestHashFunction(t *testing.T) {
	testTable := []struct {
		url string
		expected string
	}{
		{
			url: "http://test.com",
			expected: "mD2SasNU8O",
		},
		{
			url: "http://test1.com",
			expected: "ADYSqs2UQb",
		},
		{
			url: "http://test2.com",
			expected: "06nS1s6Ubo",
		},
	}

	for _, testCase := range testTable {
		result := hash.HashFunction(testCase.url)

		t.Logf("Calling HashFunction(%s), result: %s\n", 
				testCase.url, result)

		assert.Equal(t, testCase.expected, result, 
			fmt.Sprintf("Incorrect value. Expect: %s, got %s",
				testCase.expected, result))
	}
	
}
