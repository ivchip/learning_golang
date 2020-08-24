// Algorithm that asks for 2 numbers and shows the average of both.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main()  {
	fmt.Print("Number 1: ")
	number1 := getNumber()
	fmt.Print("Number 1: ")
	number2 := getNumber()
	fmt.Printf("The average is: %.2f\n", (number1 + number2) / 2.0)
}

func getNumber() float64  {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	number, err := strconv.ParseFloat(text, 2)
	if err != nil {
		fmt.Fprintf(os.Stderr,  "Not is a number: %v\n", err)
		os.Exit(1)
	}
	return number
}
