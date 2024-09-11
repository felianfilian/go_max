package main

import (
	"fmt"
	"math"
)

func main() {
	// Basic tests
	//BasicTest()

	// Pointers
	PointerTest()
}

func BasicTest() {
	const infaltionRate = 2.5
	var investAmount float64 = 1000
	returnRate := 3.7
	var years float64 = 5

	fmt.Print("Invest Amount: ")
	fmt.Scan(&investAmount)

	futureValue := investAmount * math.Pow((1+returnRate/100), years)
	realFutureValue := futureValue / math.Pow(1+infaltionRate/100, years)
	fmt.Print(realFutureValue)
}

func PointerTest() {
	var health = 100
	var pHealth = &health
	fmt.Println(pHealth)
	var iHealth = *pHealth
	fmt.Println(iHealth)
	fmt.Println(Calc(&health))
}

func Calc(heal *int) int {
	return *heal + 10
}
