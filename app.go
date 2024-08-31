package main

import (
	"fmt"
	"math"
)

func main() {
	var investAmount = 1000
	var returnRate = 3.7
	var years = 5

	var futureValue = float64(investAmount)* math.Pow((1 + returnRate / 100), float64(years)) 
	fmt.Print(futureValue)
}