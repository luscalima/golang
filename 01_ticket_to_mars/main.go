package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const distance = 62100000   // km
	const secondsPerDay = 86400 // seconds
	const minSpeed = 16         // km/s

	fmt.Printf("%-18v %v %-12v %v\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Println("==========================================")

	for i := 0; i < 10; i++ {
		var spaceline string
		switch id := rand.Intn(3); id {
		case 0:
			spaceline = "Space Adventures"
		case 1:
			spaceline = "SpaceX"
		case 2:
			spaceline = "Virgin Galactic"
		}

		var speed = minSpeed + rand.Intn(15)                  // km/s
		var tripDuration = distance / (speed * secondsPerDay) // days
		var tripType = rand.Intn(2) + 1                       // 1 One-way 2 Round-trip
		var tripPrice = (36 + speed - minSpeed) * tripType    // millions
		var typeLabel string

		if tripType == 1 {
			typeLabel = "One-way"
		} else {
			typeLabel = "Round-trip"
		}

		fmt.Printf("%-18v %4v %-12v $%4v\n", spaceline, tripDuration, typeLabel, tripPrice)
	}
}
