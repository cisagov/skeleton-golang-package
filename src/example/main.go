package main // The package main tells the Go compiler that the package should compile as an executable program and not a shared library

// example is an example Golang package
import (
	"example/pkg"
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
