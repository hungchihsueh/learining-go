package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5
	var investmentAmount float64
	var years float64
	var interestRate float64
	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)
	fmt.Print("Enter the interest rate: ")
	fmt.Scan(&interestRate)

	futureValue := investmentAmount * math.Pow(1+interestRate/100, years)
	realFututreValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println(realFututreValue)
}
