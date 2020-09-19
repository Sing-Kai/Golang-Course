package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*

Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
The program should be written as a loop. Before entering the loop,
the program should create an empty integer slice of size (length) 3.
During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
The slice must grow in size to accommodate any number of integers which the user decides to enter.
The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.

*/

func main() {

	slice := make([]int, 3)

	for {
		fmt.Println("Enter an integer: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()

		if isExit(line) {
			fmt.Println("Exiting!")
			break
		}

		if tempInt, err := strconv.Atoi(line); err == nil {

			// fmt.Printf("int entered %s : \n", line)

			slice = append(slice, tempInt)

			sortSlice(slice)

		} else {

			fmt.Println("Invalid data type entered, please enter an integer or press X to exit")

		}
	}
}

func sortSlice(s []int) {
	sort.Ints(s)
	printSlice(s)
}

func isExit(line string) bool {

	res := strings.ToUpper(line)

	return res == "X"

}

func printSlice(s []int) {
	// fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	fmt.Printf("%v\n", s)
}
