package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Enter 10 integers (e.g 23 54 23 12) : ")

	s := scanUserInput()

	fmt.Println("Sorting integers...")
	bubbleSort(s)

	fmt.Println("Result of sorted slice:")
	printSlice(s)

}

func bubbleSort(s []int) {

	for i := 1; i < len(s); i++ {
		swap(s, i)
	}
}

func swap(s []int, i int) {

	for i := 1; i < len(s); i++ {
		if s[i-1] > s[i] {
			s[i], s[i-1] = s[i-1], s[i]
		}
	}
}

func scanUserInput() []int {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	line = strings.Trim(line, " ")

	s := convertLineToSlice(line)

	return s
}

func convertLineToSlice(line string) []int {

	arr := strings.Split(line, " ")
	s := make([]int, len(arr))

	for i, str := range arr {

		if tempInt, err := strconv.Atoi(str); err != nil {
			panic(err)
		} else {
			s[i] = tempInt
		}
	}

	return s
}

func printSlice(s []int) {
	fmt.Println(s)
}
