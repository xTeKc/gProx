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

//<Write> writes data to Stdout
func (fooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out > ")
	return os.Stdout.Write(b)
}

func main() {
	//instantiate reader and writer
	var (
		reader FooReader
		writer FooWriter
	)

	//create buffer to hold input/output
	input := make([]byte, 4096)

	//use reader to read input
	s, err := reader.Read(input)
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	fmt.Printf("Read %d bytes from stdin\n", s)
}