package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

func eat(id int, eatCounter int) {
	fmt.Printf("Philosopher: %d is eating for the %d time\n", id, eatCounter)
	time.Sleep(1 * time.Second)
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
		fmt.Printf("Philosopher: %d is thinking...\n", id)
		time.Sleep(1 * time.Second)
	}
}

func isAvailable(id int, forks chan [5]bool) bool {
	mutex.Lock() // critical section is locked
	defer mutex.Unlock()
	myForks := <-forks
	availability := myForks[id]
	forks <- myForks
	return availability
}

func setAvailable(id int, forks chan [5]bool) {
	mutex.Lock() // critical section is locked
	defer mutex.Unlock()
	myForks := <-forks
	myForks[id] = !myForks[id]
	forks <- myForks
}

func main() {
	forks := make(chan [5]bool, 2)
	forks <- [5]bool{true, true, true, true, true}

	for i := 0; i < 5; i++ {
		go think(i, forks) // philosopher goroutines
	}

	time.Sleep(30 * time.Second) // main thread

}
