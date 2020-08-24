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
	ages = append(ages, getValue())
	fmt.Print("Age 2: ")
	ages = append(ages, getValue())
	fmt.Printf("The oldest person has: %d\n", majorAge(ages))
}

func getValue() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	number, err := strconv.Atoi(text)
	if err != nil {
		fmt.Fprintf(os.Stderr,  "Not is a age: %v\n", err)
		os.Exit(1)
	}
	return number
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
