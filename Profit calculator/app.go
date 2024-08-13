package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64
	fmt.Print("Enter the revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter the expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter the tax rate: ")
	fmt.Scan(&taxRate)
	fmt.Println("earnings before tax(EBT) ", revenue-expenses)
	fmt.Println("earnings after tax(Profit) ", (revenue-expenses)*(100-taxRate)/100)
	fmt.Println("EBT/profit ratio ", (revenue-expenses)/((revenue-expenses)*(100-taxRate)/100))
}
