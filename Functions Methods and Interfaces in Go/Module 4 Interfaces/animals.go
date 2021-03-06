package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type cmdType int

const (
	newAnimal cmdType = 0
	query     cmdType = 1
)

type userRequest struct {
	cmd      string
	inputOne string
	inputTwo string
}

//Animal interface
type Animal interface {
	Eat()
	Move()
	Speak()
}

//Cow struct
type Cow struct {
	food       string
	locomotion string
	noise      string
}

//Eat return what animal eats
func (a Cow) Eat() {
	fmt.Printf("%s \n", a.food)
}

//Move return how animal moves
func (a Cow) Move() {
	fmt.Printf("%s \n", a.locomotion)
}

//Speak return what noise animal makes
func (a Cow) Speak() {
	fmt.Printf("%s \n", a.noise)
}

//Bird struct
type Bird struct {
	food       string
	locomotion string
	noise      string
}

//Eat return what animal eats
func (a Bird) Eat() {
	fmt.Printf("%s \n", a.food)
}

//Move return how animal moves
func (a Bird) Move() {
	fmt.Printf("%s \n", a.locomotion)
}

//Speak return what noise animal makes
func (a Bird) Speak() {
	fmt.Printf("%s \n", a.noise)
}

//Snake struct
type Snake struct {
	food       string
	locomotion string
	noise      string
}

//Eat return what animal eats
func (a Snake) Eat() {
	fmt.Printf("%s \n", a.food)
}

//Move return how animal moves
func (a Snake) Move() {
	fmt.Printf("%s \n", a.locomotion)
}

//Speak return what noise animal makes
func (a Snake) Speak() {
	fmt.Printf("%s \n", a.noise)
}

func main() {

	var a animalCollection

	a.animals = map[string]Animal{}

	for {

		fmt.Print("> ")
		userRequest := scanUserRequest()

		if !validCmd(userRequest.cmd) {
			fmt.Println("> incorrect command entered... ")
			fmt.Println("> new animal example: newanimal james cow")
			fmt.Println("> query animal exmaple: query james eat")
		} else {
			processRequest(userRequest, a)
		}
	}

}

//create animal collection struct to store map of animals
type animalCollection struct {
	animals map[string]Animal
}

//method add animal
func (a *animalCollection) newAnimal(name string, animal Animal) error {

	if _, keyPresent := a.animals[name]; keyPresent {
		return errors.New("animal name has been used already, please choose another")
	}

	a.animals[name] = animal
	return nil
}

func processRequest(userRequest userRequest, a animalCollection) {

	cmdType := commandType(userRequest.cmd)

	if cmdType == newAnimal {
		processNewAnimal(userRequest, a)
	}

	if cmdType == query {
		processQuery(userRequest, a)
	}
}

func processNewAnimal(userRequest userRequest, a animalCollection) {

	if animal, err := getAnimal(userRequest.inputTwo); err != nil {
		fmt.Println(err)
	} else {

		if err := a.newAnimal(userRequest.inputOne, animal); err != nil {
			fmt.Println(err)
		}
	}
}

func processQuery(userRequest userRequest, a animalCollection) {

	name := userRequest.inputOne

	if animal, ok := a.animals[name]; ok {
		animalAction(userRequest.inputTwo, animal)
	} else {
		println("name entered was not found")
	}
}

//process to decide if its a new animal or query
func commandType(str string) cmdType {

	if str == "newanimal" {
		return newAnimal
	}
	return query

}

func validCmd(cmd string) bool {
	return strings.ToLower(cmd) == "newanimal" || strings.ToLower(cmd) == "query"
}

func getAnimal(str string) (Animal, error) {

	str = strings.ToLower(str)

	if str == "cow" {
		cow := Cow{food: "grass", locomotion: "walk", noise: "moo"}
		return cow, nil
	}

	if str == "bird" {
		bird := Bird{food: "worms", locomotion: "fly", noise: "peep"}
		return bird, nil
	}

	if str == "snake" {
		snake := Snake{food: "mice", locomotion: "slither", noise: "hsss"}
		return snake, nil
	}

	return nil, errors.New("animal type was unrecognised")
}

func animalAction(input string, a Animal) {

	if input == "eat" {
		a.Eat()
	} else if input == "move" {
		a.Move()
	} else if input == "speak" {
		a.Speak()
	} else {
		println("input incorrect must be: eat, move or speak")
	}
}

//function scan user input
func scanUserRequest() userRequest {

	emptyStr := " "
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	line = strings.Trim(line, emptyStr)
	arr := strings.Split(line, emptyStr)

	var userRequest userRequest

	if len(arr) != 3 {

		userRequest.cmd = emptyStr
		userRequest.inputOne = emptyStr
		userRequest.inputTwo = emptyStr

		return userRequest
	}

	userRequest.cmd = arr[0]
	userRequest.inputOne = arr[1]
	userRequest.inputTwo = arr[2]

	return userRequest

}
