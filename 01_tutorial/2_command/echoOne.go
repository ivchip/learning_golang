// EchoOne prints its 2_command-line arguments
package main

import (
	"fmt"
	"os"
	"time"
)

func main()  {
	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
		fmt.Println("Argument:", i)
	}
	fmt.Println("Concatenates the values:", s)
	fmt.Printf("%.8fs elapsed\n", time.Since(start).Seconds())
}
