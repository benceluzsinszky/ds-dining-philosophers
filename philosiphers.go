package main

import (
	"fmt"
	"time"
)

func eat(id int, eatCounter int) {
	time.Sleep(2 * time.Second)
	fmt.Printf("Philosopher: %d is eating for the %d time", id, eatCounter)
}

func think(id int, forks chan [5]bool) {
	eatCounter := 0
	for {
		time.Sleep(2 * time.Second)
		fmt.Printf("Philosopher: %d is thinking...", id)
		leftFork := isAvailable(id, forks)
		rightFork := isAvailable((id+1)%5, forks)
		if leftFork && rightFork {
			setAvailable(id, forks)
			setAvailable((id+1)%5, forks)
			eat(id, eatCounter)
			eatCounter++
			setAvailable(id, forks)
			setAvailable((id+1)%5, forks)
		}
	}
}

func isAvailable(id int, forks chan [5]bool) bool {
	myForks := <-forks
	availability := myForks[id]
	forks <- myForks
	return availability
}

func setAvailable(id int, forks chan [5]bool) {
	myForks := <-forks
	myForks[id] = !myForks[id]
	forks <- myForks
}

func main() {
	// Creating a channel
	channel := make(chan [5]bool)
	channel <- [5]bool{true, true, true, true, true}

	// Creating 10.000 workers to execute the task
	for i := 0; i < 5; i++ {
		go think(i, channel)
	}

}
