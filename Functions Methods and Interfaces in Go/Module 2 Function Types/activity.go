package main

import (
	"fmt"
	"math"
)

func main() {

	//prompt user to enter acceleration, velocity and displacement
	acceleration := getUserInput("enter acceleration: ")
	velocity := getUserInput("enter verlocity: ")
	displacement := getUserInput("enter displacement: ")

	fn := genDisplaceFn(acceleration, velocity, displacement)

	//promot user to enter time one
	t1 := getUserInput("enter time value: ")
	fmt.Print("Result: ")
	fmt.Println(fn(t1))

	//promot user to enter time one
	t2 := getUserInput("enter another time value: ")
	fmt.Print("Result: ")
	fmt.Println(fn(t2))

}

func genDisplaceFn(a float64, v float64, d float64) func(t float64) float64 {

	fn := func(t float64) float64 {
		//formula
		s := (a*math.Pow(t, 2))/2 + v*t + d

		return s
	}

	return fn
}

func getUserInput(message string) float64 {

	fmt.Printf("%s ", message)

	var val float64
	_, err := fmt.Scan(&val)

	if err != nil {
		fmt.Println("value entered must be float62")
	}

	return val
}
