package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

var philosophers = []Philosopher{
	{name: "Plato", rightFork: 4, leftFork: 0},
	{name: "Socrates", rightFork: 0, leftFork: 1},
	{name: "Aristotle", rightFork: 1, leftFork: 2},
	{name: "Pascal", rightFork: 2, leftFork: 3},
	{name: "Locke", rightFork: 3, leftFork: 4},
}

var hunger = 3
var eatime = 1 * time.Second
var thinktime = 3 * time.Second
var sleeptime = 1 * time.Second

var orderMutex sync.Mutex
var orderFinished []string

func main() {

	fmt.Println("Dining Philosophers Problem")
	fmt.Println("--------------------------------")
	fmt.Println("The table is empty")

	time.Sleep(sleeptime)

	// start the meal
	dine()

	// print out the finished message
	fmt.Println("The table is empty!!")

}

func dine() {

	eatime = 0 * time.Second
	thinktime = 0 * time.Second
	sleeptime = 0 * time.Second

	// wg is the WaitGroup that keeps track of how many philosophers are still at the table. When
	// it reaches zero, everyone is finished eating and has left. We add 5 (the number of philosophers) to this
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	// We want everyone to be seated before they start eating, so create a WaitGroup for that, and set it to 5.
	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	// forks is a map of all 5 forks. Forks are assigned using the fields leftFork and rightFork in the Philosopher
	// type. Each fork, then, can be found using the index (an integer), and each fork has a unique mutex.
	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	// Start the meal by iterating through our slice of Philosophers.
	for i := 0; i < len(philosophers); i++ {
		// fire off a goroutine for the current philosopher
		go diningProblem(philosophers[i], wg, forks, seated)

	}
	// Wait for the philosophers to finish. This blocks until the wait group is 0.
	wg.Wait()
}

// diningProblem is the function fired off as a goroutine for each of our philosophers. It takes one
// philosopher, our WaitGroup to determine when everyone is done, a map containing the mutexes for every
// fork on the table, and a WaitGroup used to pause execution of every instance of this goroutine
// until everyone is seated at the table.
func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("%s is seated at the table.\n", philosopher.name)
	seated.Done()

	seated.Wait()

	for i := hunger; i > 0; i-- {

		if philosopher.leftFork > philosopher.rightFork {
			forks[philosopher.rightFork].Lock()
			fmt.Printf("%s takes the right fork", philosopher.name)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("%s takes the left fork", philosopher.name)
		} else {
			forks[philosopher.leftFork].Lock()
			fmt.Printf("%s takes the left fork", philosopher.name)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("%s takes the right fork", philosopher.name)
		}

		fmt.Printf("\t%s has both forks and is eating.\n", philosopher.name)
		time.Sleep(eatime)

		fmt.Printf("\t%s is thinking.\n", philosopher.name)
		time.Sleep(thinktime)

		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()

		fmt.Printf("\t%s put down forks.\n", philosopher.name)

	}

	fmt.Printf("\t%s is satisfied.\n", philosopher.name)
	fmt.Printf("\t%s left the table.\n", philosopher.name)

	orderMutex.Lock()
	orderFinished = append(orderFinished, philosopher.name)
	orderMutex.Unlock()

}
