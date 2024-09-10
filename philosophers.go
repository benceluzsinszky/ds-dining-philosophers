package main

import (
	"fmt"
	"time"
)

func eat(id int, eatCounter int) {
	fmt.Printf("Philosopher: %d is eating for the %d time\n", id, eatCounter)
	time.Sleep(1 * time.Second)
	fmt.Printf("Philosopher: %d is done eating\n", id)
}

func think(id int, leftFork, rightFork chan bool) {
	eatCounter := 0
	for {
		select {
		case <-leftFork:
			select {
			case <-rightFork:
				eatCounter++
				eat(id, eatCounter)
				leftFork <- true
			default:
				leftFork <- true
			}
		default:
		}
		rightFork <- true

		fmt.Printf("Philosopher: %d is thinking...\n", id)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	forks := make([]chan bool, 5) // forks are channels
	for i := 0; i < 5; i++ {
		forks[i] = make(chan bool, 1)
		forks[i] <- true
	}

	for i := 0; i < 5; i++ {
		go think(i, forks[i], forks[(i+1)%5]) // philosopher goroutines
	}

	time.Sleep(30 * time.Second) // main thread

}
