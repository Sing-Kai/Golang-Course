package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

/*
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
*/

func main() {

	userMap := make(map[string]string)

	fmt.Println("Please enter name: ")

	name := scanLine()

	fmt.Println("Please enter address: ")

	address := scanLine()

	userMap["name"] = name

	userMap["address"] = address

	jsonData, _ := json.Marshal(userMap)

	fmt.Println(string(jsonData))

}

func scanLine() string {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	return line

}
