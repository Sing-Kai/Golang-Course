package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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

//set the name lengths a max of 20
func setNameLen(name string) string {

	if len(name) > maxLength {

		return name[:maxLength]
	}

	return name
}

func main() {

	//prompt the user for the name of the text file
	fmt.Println("enter name of file (include .txt extension) :")

	fileName := scanFileName()
	file, err := os.Open(fileName)

	if err != nil {

		fmt.Printf("Error %s /n", err)

	} else {
		//convert file into slice of name struct
		nameSlice := convertFileToSlice(file)
		//print out slice of struct
		printNameStructs(nameSlice)
	}
}

func printNameStructs(names []name) {

	for _, name := range names {
		fmt.Printf("%s %s \n", name.fname, name.lname)
	}
}

func convertFileToSlice(file io.Reader) []name {

	scanner := bufio.NewScanner(file)
	names := make([]name, 1)

	//read each line and create a name struct with details and appends to slice
	for scanner.Scan() {

		var text = scanner.Text()
		n := createNameStruct(text)
		names = append(names, n)

	}

	return names
}

func createNameStruct(line string) name {

	fname, lname := getNameDetails(line)
	var n name
	n.SetLength(fname, lname)

	return n
}

//reads line and gets name details
func getNameDetails(line string) (string, string) {

	detailsArr := strings.Split(line, " ")

	emptyStr := ""

	//empty line
	if line == emptyStr {
		fmt.Println("empty line")
		return emptyStr, emptyStr
	}

	//missing last name
	if len(detailsArr) < 2 {
		return detailsArr[0], emptyStr
	}

	return detailsArr[0], detailsArr[1]
}

func scanFileName() string {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	return line

}
