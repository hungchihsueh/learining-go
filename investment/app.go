package main

import (
	"fmt"
	"math"
)
func main() {
	var investmentAmount float64 =1000
	var interestRate float64=6.5
	var years float64=10
	var futureValue= investmentAmount *math.Pow(1+interestRate/100,years)
	fmt.Print(futureValue)
}