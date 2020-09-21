package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *animal) SetFields(f string, l string, n string) {
	a.food = f
	a.locomotion = l
	a.noise = n
}

func (a *animal) eat() string {
	return a.food
}

func (a *animal) move() string {
	return a.locomotion
}

func (a *animal) speak() string {
	return a.noise
}

func main() {

	cow := new(animal)
	bird := new(animal)
	snake := new(animal)

	cow.SetFields("grass", "walk", "moo")
	bird.SetFields("worms", "fly", "peep")
	snake.SetFields("mice", "sither", "hsss")

	for {

		fmt.Println("Enter")
		fmt.Print("> ")

		a, req := scanUserRequest()

		if validEntry(a, req) {
			fmt.Printf("%s %s\n", a, req)
		}
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

	if !validAnimal(a) {
		fmt.Println("Animal not found! Please enter: cow, bird or snake")
		return false
	}

	if !validRequest(req) {
		fmt.Println("Request Not found! Please enter: eat, move or speak")
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
