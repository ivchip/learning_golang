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
	number1, err := getNumber()
	if err != nil {
		fmt.Fprintf(os.Stderr,  "Not is a number: %v\n", err)
		os.Exit(1)
	}
	fmt.Print("Number 1: ")
	number2, err1 := getNumber()
	if err1 != nil {
		fmt.Fprintf(os.Stderr,  "Not is a number: %v\n", err1)
		os.Exit(1)
	}
	fmt.Printf("The average is: %.2f\n", (number1 + number2) / 2.0)
}

func getNumber() (float64, error)  {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	number, err := strconv.ParseFloat(text, 2)
	return number, err
}
