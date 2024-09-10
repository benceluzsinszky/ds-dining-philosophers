package main

import (
	"fmt"
	"time"
)

// Fork represents a single fork, handled by a separate goroutine
func fork(signal chan bool) {
	for {
		// Wait for a signal from a philosopher to request or release the fork
		signal <- true // Fork is available (this is the signal that a philosopher can take the fork)
		<-signal       // Wait for a signal from a philosopher to release the fork
	}
}

// Philosopher's behavior: thinking and attempting to eat
func think(id int, leftFork, rightFork chan bool) {
	eatCounter := 0
	fmt.Printf("Philosopher: %d is thinking...\n", id)
	for {
		select {
		case <-leftFork: // Acquire the left fork
			select {
			case <-rightFork: // Acquire the right fork
				eatCounter++
				eat(id, eatCounter)
				leftFork <- true                                   // Release the left fork
				rightFork <- true                                  // Release the right fork
				fmt.Printf("Philosopher: %d is thinking...\n", id) // Philosopher is thinking after eating
				time.Sleep(100 * time.Millisecond)

			default:
				// Release the left fork if the right fork is unavailable to prevent deadlock
				leftFork <- true
			}
		default:
		}
	}
}

// Simulates a philosopher eating
func eat(id int, eatCounter int) {
	fmt.Printf("Philosopher: %d is eating for the %d time\n", id, eatCounter)
	time.Sleep(1 * time.Second)
	//fmt.Printf("Philosopher: %d is done eating\n", id)
}

func main() {
	numPhilosophers := 5

	// Create fork channels and run each fork as a goroutine
	forkChannels := make([]chan bool, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		forkChannels[i] = make(chan bool, 1)
		go fork(forkChannels[i])
	}

	// Start philosopher goroutines
	for i := 0; i < numPhilosophers; i++ {
		go think(i, forkChannels[i], forkChannels[(i+1)%numPhilosophers])
	}

	// Let the philosophers eat and think for 30 seconds
	time.Sleep(30 * time.Second)
}
