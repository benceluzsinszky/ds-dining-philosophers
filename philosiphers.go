package main

import (
	"fmt"
	"time"
)

func eat(id int, eatCounter int) {
	time.Sleep(2 * time.Second)
	fmt.Println("Philosopher: %d is eating for the %d time", id, eatCounter)
}

func thinking(id int) {
	time.Sleep(2 * time.Second)
	fmt.Println("Philosopher: %d is thinking", id)
}

func philisophise(id int, forks chan int){
	
}

func main() {
   // Creating a channel
   channel := make(chan int)

   // Creating 10.000 workers to execute the task
   for i := 0; i < 5; i++ {
      go philisophise(i, channel)
   }

   // Filling channel with 100.000 numbers to be executed
   for i := 0; i < 100000; i++ {
      channel <- i
   }

}