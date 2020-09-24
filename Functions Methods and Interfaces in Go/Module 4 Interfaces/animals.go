package main

import (
	"bufio"
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
func (a *Cow) Eat() string {
	return a.food
}

//Move return how animal moves
func (a *Cow) Move() string {
	return a.locomotion
}

//Speak return what noise animal makes
func (a *Cow) Speak() string {
	return a.noise
}

//Bird struct
type Bird struct {
	food       string
	locomotion string
	noise      string
}

//Eat return what animal eats
func (a *Bird) Eat() string {
	return a.food
}

//Move return how animal moves
func (a *Bird) Move() string {
	return a.locomotion
}

//Speak return what noise animal makes
func (a *Bird) Speak() string {
	return a.noise
}

//Snake struct
type Snake struct {
	food       string
	locomotion string
	noise      string
}

//Eat return what animal eats
func (a *Snake) Eat() string {
	return a.food
}

//Move return how animal moves
func (a *Snake) Move() string {
	return a.locomotion
}

//Speak return what noise animal makes
func (a *Snake) Speak() string {
	return a.noise
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
func hasName(name string) bool {
	return true
}

//method add animal
func (a *animalCollection) newAnimal(name string, animal Animal) {
	a.animals[name] = animal
}

//method return animal
func (a *animalCollection) getAnimalByName(name string) Animal {
	return a.animals[name]
}

func processRequest(userRequest userRequest, a animalCollection) string {

	if !validCmd(userRequest.cmd) {

		return "incorrect command: must be newanimal or query"
	}

	cmdType := commandType(userRequest.cmd)

	if cmdType == newAnimal {

		//valid animal is correct
		//get name
		//check name is not already in use
		//add animal to collecion, use name as key
		//add error type interface
	}

	if cmdType == query {
		//validate query command is correct
		//return animal by name from animalCollection
		//if not found throw errow
		//execute query
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

func validAnimal(a string) bool {

	return strings.ToLower(a) == "cow" || strings.ToLower(a) == "bird" || strings.ToLower(a) == "snake"
}

func validRequest(request string) bool {

	return strings.ToLower(request) == "eat" || strings.ToLower(request) == "move" || strings.ToLower(request) == "speak"
}
