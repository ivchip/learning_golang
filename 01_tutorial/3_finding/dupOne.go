/*
DupOne prints the text of each line that appears more than
once in the standard input, preceded by its count.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("With \"q\" quit program.")
	for input.Scan() {
		if input.Text() == "q" {
			break
		}
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
