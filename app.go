package main

import (
	"fmt"
	"math"
)

func main() {
	var investAmount float64 = 1000
	var returnRate = 3.7
	var years float64 = 5

	var futureValue = investAmount * math.Pow((1+returnRate/100), years)
	fmt.Print(futureValue)
}
