package main

import (
	"fmt"
)

func main() {

	fmt.Println("Enter float number:")

	var res float64

	fmt.Scanln(&res)

	var value int = int(res)

	fmt.Printf("Float value truncated to integer: %d\n", value)

}
