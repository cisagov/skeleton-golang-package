/*
Go has a testing package from the stdlib. But here
we are also using Testify - a third-party package to
provide more options. This may change in the future.
*/
package example_test

import (
	"example/src"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHello(t *testing.T) {
	// Define the input string
	input := "Person"

	// Call the Hello function with the input string
	output := example.Hello(input)

	// Define the expected output string
	expected := "Hello, Person!"

	// Compare the expected output with the actual output
	assert.Equal(t, expected, output, "Unexpected output")
}

func TestHelloFancy(t *testing.T) {
	// Set up test cases
	testCases := []struct {
		// Define additional test cases here
		output string
	}{
		{output: "ℋℯ𝓁𝓁ℴ, 𝒲ℴ𝓇𝓁𝒹!"},
	}

	// Run test cases
	for _, tc := range testCases {
		// Define the expected result
		result := example.HelloFancy()

		// Assert equality
		assert.Contains(t, tc.output, result, "Output: %d, Expected: %d, Got: %d", tc.output, result)
	}
}

func TestAddNumbers(t *testing.T) {
	// Define the expected result
	result := example.AddNumbers(5, 6)

	// Assert equality
	assert.Equal(t, 11, result, "Expected 5 + 6 to equal 11")
}
