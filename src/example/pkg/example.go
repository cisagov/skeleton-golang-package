package pkg

import (
	"fmt"

	"github.com/danielgtaylor/unistyle" // This is a third-party package
)

// Hello returns a greeting message
func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// HelloFancy returns a message using unistyle - a third-party package
func HelloFancy() string {
	return fmt.Sprintf(unistyle.Cursive("Hello, World!")) // Use a function from the third-party module
}

// AddNumbers adds two numbers together and returns the result
func AddNumbers(a, b int) int {
	return a + b
}
