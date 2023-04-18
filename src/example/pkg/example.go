package pkg

import (
	"fmt"

	"github.com/danielgtaylor/unistyle" // this is a third-party package
)

// Hello returns a greeting message
func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// HelloFancy returns a message using unixstyle - a third-party package
func HelloFancy() string {
	return fmt.Sprintf(unistyle.Cursive("Hello, World!")) // using a function from the third-party module
}

// AddNumbers adds two numbers together and returns the result
func AddNumbers(a, b int) int {
	return a + b
}
