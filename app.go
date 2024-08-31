package main

import "fmt"

func main() {
	var investAmount = 1000
	var returnRate = 3.7
	var years = 5

	var futureValue = float64(investAmount)*(1 + returnRate / 100)
	fmt.Print("future value: " + futureValue)
}