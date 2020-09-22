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

func (a *animal) Eat() string {
	return a.food
}

func (a *animal) Move() string {
	return a.locomotion
}

func (a *animal) Speak() string {
	return a.noise
}

func main() {

	cow := animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := animal{food: "mice", locomotion: "sither", noise: "hsss"}

	animals := []animal{cow, bird, snake}

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

func processRequest(a animal, req string) {

	if req == "eat" {
		fmt.Printf("%s \n", a.Eat())
	} else if req == "move" {
		fmt.Printf("%s\n", a.Move())
	} else if req == "speak" {
		fmt.Printf("%s\n", a.Speak())
	}
}

func getAnimal(a string, animals []animal) animal {

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
