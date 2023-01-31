package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func withSleep(id int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("[%d] counting %d\n", id, i)
	}
}

func testPrintWaitGroup(id int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("[%d] counting %d\n", id, i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(amt * time.Millisecond)
	}
}

func main() {
	//	goroutine with time sleep
	// go withSleep(0)

	// time.Sleep(1 * time.Second)
	//	goroutine with time sleep

	// goroutine with waitgroup
	// var wg sync.WaitGroup
	// for i := 0; i < 3; i++ {
	// 	wg.Add(1)

	// 	go func(i int) {
	// 		defer wg.Done()
	// 		testPrintWaitGroup(i)
	// 	}(i)
	// }
	// wg.Wait()
	// goroutine with waitgroup

	// goroutine share resources include mutex
	var sharedResource string
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			mutex.Lock()
			sharedResource = fmt.Sprintf("key owned by: [%d]", id)
			fmt.Println("Previous value: ", sharedResource)
			fmt.Println("Current value: ", sharedResource)
			mutex.Unlock()
		}(i)
	}
	wg.Wait()

	fmt.Println("final resource: ", sharedResource)
	// goroutine share resources
}
