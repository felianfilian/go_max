package main

import (
	"fmt"
	"math"
)

func main() {

	// var name customType = "Mario"
	// name.log()

	// Basic tests
	//BasicTest()

	// Pointers
	//PointerTest()

	// Generics
	//Generics()

	// ARrays
	//Arrays()

	// Maps
	Maps()

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

func Generics() {
	myNum := add(5, 3)
	fmt.Println(myNum)
}

func add[T int | float64](a, b T) T {
	return a + b
}

func Arrays() {
	myArray := []string{"teenis", "golf", "soccer"}
	fmt.Println(myArray)
	fmt.Println(myArray[0])
	fmt.Println(myArray[1:])
	newArray := []string{}
	newArray = append(newArray, myArray...)
	fmt.Println(newArray)
}

func Maps() {
	websites := map[string]string{"Google": "www.google.com", "Amazon": "www.amazon.com"}

	fmt.Println(websites)
	websites["facebook"] = "www.facebook.com"
	fmt.Println(websites)
}
