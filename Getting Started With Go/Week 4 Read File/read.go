package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type name struct {
	fname string
	lname string
}

const (
	maxLength = 20
)

func (n *name) SetLength(fname string, lname string) {

	n.fname = setNameLen(fname)

	n.lname = setNameLen(lname)

}

func setNameLen(name string) string {

	if len(name) > maxLength {

		return name[:maxLength]
	}

	return name
}

func main() {

	//prompt the user for the name of the text file
	fmt.Println("enter name of file (include .txt extension) :")

	fileName := scanLine()

	file, err := os.Open(fileName)

	if err != nil {

		fmt.Printf("Error %s /n", err)

	} else {

		//convert file into slice of name struct
		nameSlice := convertNamesToSlice(file)

		//print out slice of struct
		printNameStructs(nameSlice)

	}
}

func printNameStructs(names []name) {

	for _, name := range names {

		fmt.Printf("first name: %s last name: %s /n", name.fname, name.lname)
	}
}

func convertNamesToSlice(file io.Reader) []name {

	scanner := bufio.NewScanner(file)

	names := make([]name, 1)

	for scanner.Scan() {

		//convert line into name struct

	}

	return names
}

func scanLine() string {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	return line

}
