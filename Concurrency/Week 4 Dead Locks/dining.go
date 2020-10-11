package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Dining Philosopher Problem
*/
type chopStick struct{ sync.Mutex }

type philo struct {
	leftCS, rightCS *chopStick
}

// host to limit two philosophers to eat at the same time
var host = make(chan bool, 2)

func (p philo) eat(philo int, wg *sync.WaitGroup) {

	for i := 0; i < 3; i++ {

		//get host permission to eat, wait if full
		host <- true

		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("starting to eat %v \n", philo)
		time.Sleep(1 * time.Second)
		fmt.Printf("finishing eating %v \n", philo)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
		//tell host eating has done frees channel
		<-host

	}

	wg.Done()
}

func main() {

	cSticks := make([]*chopStick, 5)

	//create chopsticks
	for i := 0; i < 5; i++ {
		cSticks[i] = new(chopStick)
	}

	philos := make([]*philo, 5)

	//create 5 philosophers with chopsticks
	for i := 0; i < 5; i++ {
		philos[i] = &philo{cSticks[i], cSticks[(i+1)%5]}
	}

	//start eating
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].eat(i+1, &wg)
	}

	wg.Wait()

}
