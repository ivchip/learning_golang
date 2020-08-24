// EchoTwo prints its 2_command-line arguments
package main

import (
	"fmt"
	"os"
	"time"
)

func main()  {
	start := time.Now()
	s, sep := "", ""
	for i, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
		fmt.Println("Argument:", i)
	}
	fmt.Println("Concatenates the values:", s)
	fmt.Printf("%.8fs elapsed\n", time.Since(start).Seconds())
}
