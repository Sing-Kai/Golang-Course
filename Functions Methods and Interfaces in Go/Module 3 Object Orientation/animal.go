package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Animal class
type Animal struct {
	food       string
	locomotion string
	noise      string
}

//Eat return what animal eats
func (a *Animal) Eat() string {
	return a.food
}

//Move return how animal moves
func (a *Animal) Move() string {
	return a.locomotion
}

//Speak return what noise animal makes
func (a *Animal) Speak() string {
	return a.noise
}

func main() {

	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "sither", noise: "hsss"}

	animals := []Animal{cow, bird, snake}

	fmt.Println("Enter request. Press ctrl + c to exit")

	for {

		fmt.Print("> ")
		a, req := scanUserRequest()

		if validEntry(a, req) {
			selectedAnimal := getAnimal(a, animals)
			processRequest(selectedAnimal, req)
		}
	}
}

func processRequest(a Animal, req string) {

	if req == "eat" {
		fmt.Printf("%s \n", a.Eat())
	} else if req == "move" {
		fmt.Printf("%s\n", a.Move())
	} else if req == "speak" {
		fmt.Printf("%s\n", a.Speak())
	}
}

func getAnimal(a string, animals []Animal) Animal {

	if a == "cow" {
		return animals[0]
	} else if a == "bird" {
		return animals[1]
	} else {
		return animals[2]
	}
}

func scanUserRequest() (string, string) {

	emptyStr := " "
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	line = strings.Trim(line, emptyStr)
	arr := strings.Split(line, emptyStr)

	if len(arr) != 2 {
		return emptyStr, emptyStr
	}

	return arr[0], arr[1]

}

func validEntry(a string, req string) bool {

	if !validAnimal(a) || !validRequest(req) {
		fmt.Println("Invalid entry")
		fmt.Println("Please enter animal followed by action. Example: cow eat")
		fmt.Println("Valid animals: cow, bird or snake")
		fmt.Println("Valid actions: eat, move or speak")
		return false
	}

	return true
}

func validAnimal(a string) bool {

	return strings.ToLower(a) == "cow" || strings.ToLower(a) == "bird" || strings.ToLower(a) == "snake"
}

func validRequest(request string) bool {

	return strings.ToLower(request) == "eat" || strings.ToLower(request) == "move" || strings.ToLower(request) == "speak"
}
