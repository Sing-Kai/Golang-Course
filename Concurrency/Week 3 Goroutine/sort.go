package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	c := make(chan []int)
	m := make(chan []int)
	result := make(chan []int)

	// numbers := []int{8, 7, 6, 5, 4, 3, 2, 1, 0, 1, 2, 3, 3, 10, 6, 39}
	// numbers := []int{8, 7, 6}
	numbers := scanUserInput()

	if len(numbers) < 4 {
		fmt.Println("Not enough integers entered, please enter 4 or more")
		os.Exit(1)
	}

	n := len(numbers) / 4
	remainder := len(numbers) % 4

	start := 0
	end := 0

	for i := 0; i < 4; i++ {
		end += n
		if i < 3 {
			go createSubSlice(numbers, start, end, c)
		} else {
			go createSubSlice(numbers, start, end+remainder, c)
		}
		start += n
	}

	subSliceOne := <-c
	subSliceTwo := <-c

	go merge(subSliceOne, subSliceTwo, m)

	subSliceThree := <-c
	subSliceFour := <-c

	go merge(subSliceThree, subSliceFour, m)

	go merge(<-m, <-m, result)

	fmt.Println(<-result)

}

func merge(sOne []int, sTwo []int, m chan []int) {

	res := append(sOne, sTwo...)
	sort.Ints(res)
	m <- res
}

func createSubSlice(slice []int, i int, j int, c chan []int) {

	res := slice[i:j]

	fmt.Println("Sorting slice: ", res)

	sort.Ints(res)
	c <- res
}

func scanUserInput() []int {

	fmt.Println("Enter a list of integers separated by spaces")

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
