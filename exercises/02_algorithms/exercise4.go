/*
Algorithm to read 3 different integers from each other
and determine the largest number of the three.
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main()  {
	var numbers []int
	fmt.Print("Number 1: ")
	numbers = append(numbers, getValueInt(numbers))
	fmt.Print("Number 2: ")
	numbers = append(numbers, getValueInt(numbers))
	fmt.Print("Number 3: ")
	numbers = append(numbers, getValueInt(numbers))
	fmt.Printf("The max number is: %d\n", max(numbers))
}

func getValueInt(array []int) int  {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	number, err := strconv.Atoi(text)
	if err != nil {
		fmt.Fprintf(os.Stderr,  "Not is a number: %v\n", err)
		os.Exit(1)
	}
	for _, value := range array {
		if value == number {
			fmt.Println("The number is repeat")
			os.Exit(1)
		}
	}
	return number
}

func max(array []int) int {
	max := array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
	}
	return max
}
