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
		processRequest(userRequest, a)
	}

}

//create animal collection struct to store map of animals
type animalCollection struct {
	animals map[string]Animal
}

//method name is unique
func (a *animalCollection) hasName(name string) bool {

	_, ok := a.animals[name]

	return ok
}

//method add animal
func (a *animalCollection) newAnimal(name string, animal Animal) {

	if _, keyPresent := a.animals[name]; !keyPresent {
		a.animals[name] = animal
	}
}

//method return animal
func (a *animalCollection) getAnimalByName(name string) Animal {

	return a.animals[name]
}

func processRequest(userRequest userRequest, a animalCollection) {

	if !validCmd(userRequest.cmd) {
		fmt.Println("incorrect command: must be newanimal or query")
	}

	cmdType := commandType(userRequest.cmd)

	if cmdType == newAnimal {

		if animal, err := getAnimal(userRequest.inputTwo); err != nil {
			fmt.Println(err)
		} else {
			a.newAnimal(userRequest.inputOne, animal)
		}
	}

	if cmdType == query {
		name := userRequest.inputOne
		animal := a.animals[name]
		animalAction(userRequest.inputTwo, animal)
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
	return cmd == "newanimal" || cmd == "query"
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

func getAnimal(str string) (Animal, error) {

	str = strings.ToLower(str)

	if str == "cow" {
		cow := Cow{food: "grass", locomotion: "walk", noise: "mo"}
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

func validAnimal(a string) bool {

	return strings.ToLower(a) == "cow" || strings.ToLower(a) == "bird" || strings.ToLower(a) == "snake"
}

func validRequest(request string) bool {

	return strings.ToLower(request) == "eat" || strings.ToLower(request) == "move" || strings.ToLower(request) == "speak"
}
