package main

import (
	"fmt"
	"sync"
)

type Income struct {
	Source string
	Amount int
}

func main() {
	var bankBalance int
	var balance sync.Mutex
	var wg sync.WaitGroup

	incomes := []Income{
		{Source: "Main Job", Amount: 200},
		{Source: "Part time Job", Amount: 50},
		{Source: "gift", Amount: 10},
		{Source: "Investment", Amount: 30},
	}

	wg.Add(len(incomes))
	for _, income := range incomes {
		go func(income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()
				fmt.Printf("On week %d, you earn  %d.00 from %s\n", week, income.Amount, income.Source)
			}
		}(income)
	}
	wg.Wait()

	fmt.Println(bankBalance)

}

// func updateMessage(s string, mutex *sync.Mutex) {
// 	defer wg.Done()
// 	mutex.Lock()
// 	msg = s
// 	mutex.Unlock()
// }

// func main() {

// 	var mutex sync.Mutex

// 	msg = "Hello, world!"
// 	wg.Add(2)
// 	go updateMessage("Hello world", &mutex)
// 	go updateMessage("Hello, Yayi", &mutex)
// 	wg.Wait()
// 	fmt.Println(msg)
// }
