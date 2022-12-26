package main

import (
	"time"

	"github.com/fatih/color"
)

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientsChan     chan string
	Open            bool
}

func (shop *BarberShop) addBarber(barber string) {
	shop.NumberOfBarbers++

	go func() {
		isSleeping := false
		color.Yellow("%s goes to the waiting room to check for clients", barber)

		for {
			//  if there are no clients, the barber goes to sleep
			if len(shop.ClientsChan) == 0 {
				color.Yellow("there is nothing to do, so %s takes a nap", barber)
				isSleeping = true
			}

			client, shopOpen := <-shop.ClientsChan

			if shopOpen {
				if isSleeping {
					color.Yellow("%s wakes %s up.", client, barber)
					isSleeping = false
				}
				shop.cutHair(barber, client)
			} else {
				shop.sendBaberHome(barber)
				return
			}

		}

	}()

}

func (shop *BarberShop) cutHair(barber, client string) {
	color.Green("%s is cuting %s hair.", client, barber)
	time.Sleep(shop.HairCutDuration)
	color.Green("%s is finish cuting %s hair.", client, barber)

}

func (shop *BarberShop) sendBaberHome(barber string) {
	color.Cyan("%s is going  home.", barber)
	shop.BarbersDoneChan <- true
}

func (shop *BarberShop) closeShoForDay() {
	color.Cyan("closing sho for the day.")
	close(shop.ClientsChan)
	shop.Open = false

	for a := 0; a < shop.NumberOfBarbers; a++ {
		<-shop.BarbersDoneChan
	}

	close(shop.BarbersDoneChan)
	color.Green("------------------------------------------------------")
	color.Green("Barbers are done for the day ")
}

func (shop *BarberShop) addClient(client string) {
	color.Green(" **** %s arrives!", client)

	if shop.Open {
		select {
		case shop.ClientsChan <- client:
			color.Yellow("%s takes a seat in the waiting room ", client)
		default:
			color.Red("The waiting room is full, so %s leave ", client)
		}
	} else {
		color.Red("The shop is already closed, so %s leave ", client)
	}
}
