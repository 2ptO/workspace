package main

import "fmt"
import "sync"

/*
 * Example showing race condition within go routines
 */

var (
	counter int
	wg      sync.WaitGroup
	mu      sync.Mutex
)

func main() {
	const numRoutines = 100

	wg.Add(numRoutines)
	for id := 0; id < numRoutines; id++ {
		go incrementWithRace(id)
	}

	wg.Add(numRoutines)
	for id := 0; id < numRoutines; id++ {
		go incrementWithoutRace(id)
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
}

func incrementWithRace(id int) {
	counter++
	fmt.Printf("In routine%d counter = %d\n", id, counter)
	wg.Done()
}

func incrementWithoutRace(id int) {
	mu.Lock()
	counter++
	fmt.Printf("In routine%d counter = %d\n", id, counter)
	mu.Unlock()
	wg.Done()
}
