package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// variables
var seatingCapacity = 10
var arrivalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {
	// seed our random number generator
	rand.Seed(time.Now().UnixNano())

	// print welcome message
	color.Yellow("The Sleeping Barber Problem")
	color.Yellow("--------------------------------")

	// create channels if we need any
	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	// create a barberShop
	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientsChan:     clientChan,
		BarbersDoneChan: doneChan,
		Open:            true,
	}

	color.Green("the shop is open for the day")

	// add barbers
	shop.addBarber("Bisi")

	// create the barbershop as goroutine
	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeShoForDay()
		closed <- true

	}()

	// add clients
	i := 1
	go func() {
		for {
			randomMilliseconds := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMilliseconds)):
				shop.addClient(fmt.Sprintf("client #%d", i))
				i++
			}
		}
	}()

	// block until the barbershop is closed
	<-closed
}
