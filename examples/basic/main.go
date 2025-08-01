package main

import (
	"os"
)

func main() {
	// Generate the templ file
	component := ExampleComponent()
	
	// In a real application, you would render this to an HTTP response
	// For this example, we'll just write it to stdout
	component.Render(os.Stdout)
}
