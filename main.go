/*
This is the main package for our program. It contains
a main function which shows that the file is executable.
main() is the program's entry point. It is not called
and no value is passed as an argument or returned from
the function.

Consider this a template on where to start when writing
a Go package. The addition of directories such as `pkg`
and `test` is purposeful even if they hold single files.
This is to provide an example of what structure to use
when organizing code.

We begin by importing the required packages. Notice our
example package is present. We must import it in order
to call its functions.

Next, a logger is created to begin logging before and after we
begin printing from our example package.

Considering the `example.go` package only returns and does not
print, we make use of Go's `fmt`` package to print the
output of the example package functions we call.
*/

package main /* The package main tells the Go compiler that the package should
   compile as an executable program and not a shared library */

// example is our example Golang package
import (
	"example/src"
	"fmt"
	"log"
	"os"
)

func main() {
	// Set up a logger with a timestamp format of "YYYY-MM-DD HH:mm:ss"
	logger := log.New(os.Stdout, "", log.LstdFlags)

	logger.Printf("Running program entrypoint")

	fmt.Println(pkg.Hello("Person"))
	fmt.Println(pkg.HelloFancy())
	fmt.Println(pkg.AddNumbers(10, 6))

	logger.Printf("Program has ended")
}
