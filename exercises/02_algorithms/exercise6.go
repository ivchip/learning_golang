/*
Algorithm that requests the age of two people and
displays the age of the oldest on the screen.
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main()  {
	var ages []int
	fmt.Print("Age 1: ")
	age1, err1 := getValue()
	validateNumber(err1)
	ages = append(ages, age1)
	fmt.Print("Age 2: ")
	age2, err2 := getValue()
	validateNumber(err2)
	ages = append(ages, age2)
	fmt.Printf("The oldest person has: %d\n", majorAge(ages))
}

func getValue() (int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	number, err := strconv.Atoi(text)
	return number, err
}

func majorAge(array []int) int {
	max := array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
	}
	return max
}

func validateNumber(err error)  {
	if err != nil {
		fmt.Fprintf(os.Stderr,  "Not is a age: %v\n", err)
		os.Exit(1)
	}
}