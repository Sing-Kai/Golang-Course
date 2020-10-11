package main

import (
	"fmt"
	"sync"
)

/*
Dining Philosopher Problem
*/

type chopStick struct{ sync.Mutex }

type philo struct {
	leftCS, rightCS *chopStick
}

func (p philo) eat(philo int, wg *sync.WaitGroup) {

	for i := 0; i < 3; i++ {

		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("starting to eat %v \n", philo)

		p.rightCS.Unlock()
		p.leftCS.Unlock()

		fmt.Printf("finishing eating %v \n ", philo)

		wg.Done()
	}
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
		//change to not being the lowest number first
		philos[i] = &philo{cSticks[i], cSticks[(i+1)%5]}
	}

	//need to add some sort of wait group for syncronisation
	//all a maximum of two
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(2)
		go philos[i].eat(i, &wg)
	}

	wg.Wait()

}
