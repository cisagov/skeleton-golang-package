package example_test

import (
	"example/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHello(t *testing.T) {
	// Define the input string
	input := "Person"

	// Call the Hello function with the input string
	output := pkg.Hello(input)

	// Define the expected output string
	expected := "Hello, Person!"

	// Compare the expected output with the actual output
	assert.Equal(t, expected, output, "Unexpected output")
}

func TestHelloFancy(t *testing.T) {
	// Set up test cases
	testCases := []struct {
		// Define additional cases here
		output string
	}{
		{output: "ℋℯ𝓁𝓁ℴ, 𝒲ℴ𝓇𝓁𝒹!"},
	}

	// Run test cases
	for _, tc := range testCases {
		result := pkg.HelloFancy()
		assert.Contains(t, tc.output, result, "Output: %d, Expected: %d, Got: %d", tc.output, result)
	}
}

func TestAddNumbers(t *testing.T) {
	result := pkg.AddNumbers(5, 6)
	assert.Equal(t, 11, result, "Expected 5 + 6 to equal 11")
}
