package main

import (
	"fmt"
	// "math/rand"
	// "sync"
	"time"
)

func greet(c chan string) {
	name := <-c // get from channel
	fmt.Println("Hello ", name)
}

func greetUntilQuit(c chan string, quit chan int) {
	for {
		// select case only for channel
		select {
		case name := <-c:
			fmt.Println("Hello ", name)
		case <-quit:
			fmt.Println("quitting greeting")
			return
		}
	}
}

func nameReciever(c chan string, quit chan int) {
	for {
		select {
		case name, more := <-c:
			if more {
				fmt.Println("Hello ", name)
			} else {
				fmt.Println("recieved all data ")
				quit <- 0
			}
		}
	}
}

func nameProducer(c chan string) {
	c <- "World" // input to channel
	c <- "Banana"
	c <- "Apel"
}

func main() {
	c := make(chan string)
	quit := make(chan int)

	// go greet(c)
	// go greetUntilQuit(c, quit)

	// c <- "World" // input to channel
	// c <- "Banana"
	// c <- "Apel"

	// quit <- 0 // quit secara grasful

	go nameReciever(c, quit)
	nameProducer(c)

	close(c)
	<-quit

	time.Sleep(1 * time.Second)
}
