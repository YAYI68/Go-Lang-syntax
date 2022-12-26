package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	name      string
	leftFork  int
	rightFork int
}

var philosophers = []Philosopher{
	{name: "Yayi", leftFork: 4, rightFork: 0},
	{name: "Dupe", leftFork: 0, rightFork: 1},
	{name: "Shola", leftFork: 1, rightFork: 2},
	{name: "Kunle", leftFork: 2, rightFork: 3},
	{name: "Chidi", leftFork: 3, rightFork: 4},
}

var hunger = 3
var eatTime = 1 * time.Second
var thinkTime = 3 * time.Second
var sleepTime = 1 * time.Second

var orderMutex sync.Mutex
var orderFinished []string

func main() {

	fmt.Println("Dining Philosophers Problem")
	fmt.Println("--------------------------------")
	fmt.Println("The table is empty")

	time.Sleep(sleepTime)

	dine()

	fmt.Println("heelo")
}

func dine() {

	eatTime = 0 * time.Second
	thinkTime = 0 * time.Second
	sleepTime = 0 * time.Second

	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	for i := 0; i < len(philosophers); i++ {

		go diningProblem(philosophers[i], wg, forks, seated)
	}

	wg.Wait()

}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("%s is seated at the table.\n", philosopher.name)
	seated.Done()

	seated.Wait()

	for i := hunger; i > 0; i-- {

		if philosopher.leftFork > philosopher.rightFork {
			forks[philosopher.rightFork].Lock()
			fmt.Printf("%s takes the right fork \n", philosopher.name)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("%s takes the left fork \n", philosopher.name)
		} else {
			forks[philosopher.leftFork].Lock()
			fmt.Printf("%s takes the left fork \n", philosopher.name)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("%s takes the right fork \n", philosopher.name)
		}

		fmt.Printf("\t%s has both forks and is eating.\n", philosopher.name)
		time.Sleep(eatTime)

		fmt.Printf("\t%s is thinking.\n", philosopher.name)
		time.Sleep(thinkTime)

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
