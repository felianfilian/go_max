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
	//Maps()

	// functions advance
	FunctionAD()

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
	ratings := make(map[string]float64)
	ratings["go"] = 4.8
	ratings["aws"] = 4.5

	fmt.Println(ratings)

	names := []string{}
	names = append(names, "Mario")
	names = append(names, "Peach")
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])	
	}

}

func FunctionAD() {
	numbers := []int{1,2,3}
	doubled := transformNums(&numbers, double)
	tripled := transformNums(&numbers, triple)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformNums(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func double(number int) int {
	return number *2
}

func triple(number int) int {
	return number *3
}

