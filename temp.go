package main

import (
	"fmt"
	"strconv"
)

var numCels float64
var convert float64
var numFahren string

func main() {

	fmt.Print("Please enter the current temperature in Fahrenheit:") //ask user for value in Fahrenheit
	fmt.Scan(&numFahren)
	a, _ := strconv.ParseFloat(numFahren, 2)

	convert = userInput(a)

	temp()
	userOutput()

}

func userInput(a float64) float64 {

	if a != 0 {
		b := strconv.FormatFloat(a, 'f', 2, 64)
		c, _ := strconv.ParseFloat(b, 2)
		return c
	}
	return 0
}

func temp() {
	temp := strconv.FormatFloat(((convert - 32) * 0.5556), 'f', 2, 64) //process conversion
	numCels, _ = strconv.ParseFloat(temp, 2)
}

func userOutput() {

	if squintForecast(convert) {
		fmt.Printf("The temperature is %vC. It's too hot. You're squinting, right?", numCels)
	} else {
		fmt.Printf("The temperature is %vC. It's  too cold. Brr!", numCels)
	}
	//end result
}

func squintForecast(f float64) bool {
	return f >= 87.5
}
