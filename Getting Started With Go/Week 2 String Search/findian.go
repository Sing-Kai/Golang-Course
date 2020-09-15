package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {

	fmt.Println("Enter string: ")

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	line := scanner.Text()

	if containsIan(line) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}

func containsIan(line string) bool {

	str := strings.ToLower(line)

	count := utf8.RuneCountInString(str)

	startsWithI := str[0] == 'i'

	endsWithN := str[count-1] == 'n'

	containsA := strings.ContainsRune(str, 'a')

	return startsWithI && endsWithN && containsA
}
