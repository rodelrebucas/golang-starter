package main

import "fmt"

func main() {
	// create variable of type Writer
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))
}

// Interface - collection of function signatures; describes *behavior* unlike *struct* which which describes data
// Naming convention should be suffix with *er* with single method
type Writer interface {
	// describe the method signature
	Write([]byte) (int, error)
}

// Implicitly implements the Writer interface
type ConsoleWriter struct{}

// Create the method signature *Write*
func (vw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
