package main

import (
	"fmt"
	"time"
)

/*
race condition is defined as the condition of an electronics, software,
or other systems where the system’s substantive behavior is dependent on the sequence or timing of other uncontrollable events
In below code, race condition is produced where two go routines access by incrementing and printing the count variable concurrently.
If you running software multiple times one can observe result printed:

scenario one:
one, 1
two, 2
scenario two:
two, 2
one, 1
scenario three:
two, 2
one, 2
scenario four:
one, 2
two, 2

in addition you can run go run -race filename.go
*/

var count int

func race(str string) {
	count++
	fmt.Printf("%s, %d\n", str, count)
}

func main() {
	go race("one")
	go race("two")
	time.Sleep(1 * time.Second)
}
