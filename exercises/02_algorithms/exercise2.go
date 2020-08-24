/*
Algorithm that requests an integer and displays
the name of the corresponding month.
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main()  {
	months := map[int]string{
		1: "January",
		2: "February",
		3: "March",
		4: "April",
		5: "May",
		6: "June",
		7: "July",
		8: "August",
		9: "September",
		10: "October",
		11: "November",
		12: "December",
	}
	fmt.Print("Month number: ")
	number1 := getNumberInt()
	fmt.Printf("The month is: %s\n", months[number1])
}

func getNumberInt() int  {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	number, err := strconv.Atoi(text)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Not is a number: %v\n", err)
		os.Exit(1)
	}
	if number < 1 || number > 12 {
		fmt.Println("The number must be between 1 and 12.")
		os.Exit(1)
	}
	return number
}
