package main

import (
	"testing"
	"urls/pkg/handlers"

	"github.com/stretchr/testify/assert"
)


func TestIsValidUrl(t *testing.T) {
	testTable := []struct {
		url string
		expected bool
	}{
		{
			url: "https://test.com",
			expected: true,
		},
		{
			url: "test.com",
			expected: false,
		},
		{
			url: "https://test1.com",
			expected: true,
		},
		{
			url: "test1.com",
			expected: false,
		},
	}

	for _, testCase := range testTable {
		result := handlers.IsValidUrl(testCase.url)

		t.Logf("Is valid url: %s, result: %t\n",
				testCase.url, result)
				
		assert.Equal(t, testCase.expected, result)
	}
}