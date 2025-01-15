package main

import (
	"fmt"
	"github.com/anonymou0719/ecosystem-example_a/example"
)

func main() {
	// Call the ExampleAPI function from the api_project.
	message := example.ExampleAPI("John")
	fmt.Println(message)
}
