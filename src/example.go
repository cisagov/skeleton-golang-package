// example is an example Go package. We begin by importing
// a third-party package dependency called unistyle which provides
// Unicode text styles for Go.

package example

import (
	"fmt"

	"github.com/danielgtaylor/unistyle" // This is a third-party package
)

// Hello returns a greeting message based on the given parameter
func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// HelloFancy returns a message using unistyle - a third-party package
func HelloFancy() string {
	return fmt.Sprint(unistyle.Cursive("Hello, World!")) // Use a function from the third-party module
}

// AddNumbers adds two numbers together and returns the result
func AddNumbers(a, b int) int {
	return a + b
}
