// EchoTwo prints its command-line arguments
package main

import (
	"fmt"
	"os"
)

func main()  {
	s, sep := "", ""
	for i, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
		fmt.Println("Argument:", i)
	}
	fmt.Println("Concatenates the values:", s)
}
