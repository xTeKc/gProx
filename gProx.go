package main 

import (
	"fmt"
	"log"
	"os"
)

//FooReader defines an io.Reader to read from Stdin
type FooReader struct{}

//<Read> reads data from Stdin
func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

//FooWriter defines an io.Writer to write to Stdout
type FooWriter struct{}