package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	userMap := make(map[string]string)

	//get user name
	fmt.Println("Please enter name: ")
	name := scanLine()

	//get user address
	fmt.Println("Please enter address: ")
	address := scanLine()

	//add details to map
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
