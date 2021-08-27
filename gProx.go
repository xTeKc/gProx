package main 

import (
	"fmt"
	"log"
	"os"
)

//FooReader defines an io.Reader to read from stdin
type FooReader struct{}

//<Read> reads data from stdin
func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}