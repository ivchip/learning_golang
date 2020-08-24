/*
Algorithm that requests 2 integers and an arithmetic operator
and then must show the result of the corresponding operation.
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main()  {
	fmt.Print("Number 1: ")
	number1, err1 := strconv.Atoi(getString())
	validateErr(err1)
	fmt.Print("Number 2: ")
	number2, err2 := strconv.Atoi(getString())
	validateErr(err2)
	fmt.Print("Operation: ")
	operation := getString()
	operations(number1, number2, operation)
}

func getString() string  {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	return text
}

func validateErr(err error)  {
	if err != nil {
		fmt.Fprintf(os.Stderr,  "Not is a number: %v\n", err)
		os.Exit(1)
	}
}

func operations(number1 int, number2 int, opt string)  {
	switch opt {
	case "+" :
		fmt.Printf("The sum between %d and %d equal: %d\n", number1, number2, number1 + number2)
	case "-" :
		fmt.Printf("The rest between %d and %d equal: %d\n", number1, number2, number1 - number2)
	case "*" :
		fmt.Printf("The multiply between %d and %d equal: %d\n", number1, number2, number1 * number2)
	case "/" :
		fmt.Printf("The division between %d and %d equal: %.3f\n", number1, number2, float64(number1) / float64(number2))
	default:
		fmt.Println("No operation was selected.")
	}
}
