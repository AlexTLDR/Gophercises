package main

import "fmt"

func main() {
	var revenue, expense, taxRate float64

	fmt.Print("Enter the revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter the expense: ")
	fmt.Scan(&expense)

	fmt.Print("Enter the tax rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expense
	eat := ebt * (1 - taxRate/100)
	ration := ebt / eat

	fmt.Printf("Earnings Before Tax: %.2f\n", ebt)
	fmt.Printf("Earnings After Tax: %.2f\n", eat)
	fmt.Printf("Earnings Before Tax to Earnings After Tax Ratio: %.2f\n", ration)
}
