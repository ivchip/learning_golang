/*
Algorithm that needs to obtain the average of
a student from his three partial grades.
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main()  {
	var notes []float64
	fmt.Print("Note 1: ")
	notes = append(notes, getNumberFloat())
	fmt.Print("Note 2: ")
	notes = append(notes, getNumberFloat())
	fmt.Print("Note 3: ")
	notes = append(notes, getNumberFloat())
	fmt.Printf("The average note is: %.2f\n", avg(notes))
}

func avg(notes []float64) float64 {
	sumValues := 0.0
	for _, value := range notes {
		sumValues += value
	}
	return sumValues / float64(len(notes))
}

func getNumberFloat() float64  {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	number, err := strconv.ParseFloat(text, 2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Not is a number: %v\n", err)
		os.Exit(1)
	}
	if number < 0 || number > 5 {
		fmt.Println("The note must be between 0 and 5.")
		os.Exit(1)
	}
	return number
}
