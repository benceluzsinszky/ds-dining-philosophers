package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

func eat(id int, eatCounter int) {
	time.Sleep(2 * time.Second)
	fmt.Printf("Philosopher: %d is eating for the %d time\n", id, eatCounter)
}

func think(id int, forks chan [5]bool) {
	eatCounter := 0
	for {
		leftFork := isAvailable(id, forks)
		rightFork := isAvailable((id+1)%5, forks)
		if leftFork && rightFork {
			setAvailable(id, forks)
			setAvailable((id+1)%5, forks)
			eatCounter++
			eat(id, eatCounter)
			setAvailable(id, forks)
			setAvailable((id+1)%5, forks)
		}
		time.Sleep(1 * time.Second)
		fmt.Printf("Philosopher: %d is thinking...\n", id)
	}
}

func isAvailable(id int, forks chan [5]bool) bool {
	mutex.Lock()
	defer mutex.Unlock()
	myForks := <-forks
	availability := myForks[id]
	forks <- myForks
	return availability
}

func setAvailable(id int, forks chan [5]bool) {
	mutex.Lock()
	defer mutex.Unlock()
	myForks := <-forks
	myForks[id] = !myForks[id]
	forks <- myForks
}

func main() {
	forks := make(chan [5]bool, 2)
	forks <- [5]bool{true, true, true, true, true}

	for i := 0; i < 5; i++ {
		go think(i, forks)
	}

	time.Sleep(30 * time.Second)

}
