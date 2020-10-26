package main

// ---------------------------------------------------------
// EXERCISE: Housing Prices
//
//  We have received housing prices. Your task is to load the data into
//  appropriately typed slices then print them.
//
//  1. Check out the expected output
//
//
//  2. Check out the code below
//
//     1. header   : stores the column headers
//     2. data     : stores the real data related to each column
//     3. separator: you will use it to separate the data
//
//
//  3. Parse the data
//
//     1. Separate it into rows by using the newline character ("\n")
//
//     2. For each row, separate it by using the separator (",")
//
//
//  4. Load the data into distinct slices
//
//     1. Load the locations into a []string
//     2. Load the sizes into []int
//     3. Load the beds into []int
//     4. Load the baths into []int
//     5. Load the prices into []int
//
//
//  5. Print the header
//
//     1. Separate it by using the separator
//
//     2. Print each column 15 character wide ("%-15s")
//
//
//  6. Print the rows from the slices that you've created, line by line
//
//
// EXPECTED OUTPUT
//
//  Location       Size           Beds           Baths          Price
//  ===========================================================================
//  New York       100            2              1              100000
//  New York       150            3              2              200000
//  Paris          200            4              3              400000
//  Istanbul       500            10             5              1000000
//
// ---------------------------------------------------------
//Bonus:
// ---------------------------------------------------------
// Print the averages of the sizes, beds, baths, and prices.
//
//
// EXPECTED OUTPUT
//
//  Location       Size           Beds           Baths          Price
//  ===========================================================================
//  New York       100            2              1              100000
//  New York       150            3              2              200000
//  Paris          200            4              3              400000
//  Istanbul       500            10             5              1000000
//
//  ===========================================================================
//                 237.50         4.75           2.75           425000.00
//
// ---------------------------------------------------------
/*
Start with:
func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)
}
*/

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)
	var (
		location []string
		size     []int
		beds     []int
		baths    []int
		price    []int
	)

	rows := strings.Split(data, "\n")

	for _, row := range rows {
		line := strings.Split(row, separator)

		s, err := strconv.Atoi(line[1])
		be, err2 := strconv.Atoi(line[2])
		ba, err3 := strconv.Atoi(line[3])
		p, err4 := strconv.Atoi(line[4])
		if err != nil || err2 != nil || err3 != nil || err4 != nil {
			fmt.Println("Size/Beds/Baths/Price is not a numeric string")
		}

		location = append(location, line[0])
		size = append(size, s)
		beds = append(beds, be)
		baths = append(baths, ba)
		price = append(price, p)
	}

	for _, h := range strings.Split(header, separator) {
		fmt.Printf("%-15s", h)
	}
	fmt.Printf("\n%s\n", strings.Repeat("=", 75))
	var (
		sumSize, sumBeds, sumBaths, sumPrice int
	)
	for i := range rows {
		fmt.Printf("%-15v", location[i])
		fmt.Printf("%-15v", size[i])
		fmt.Printf("%-15v", beds[i])
		fmt.Printf("%-15v", baths[i])
		fmt.Printf("%-15v", price[i])
		fmt.Println()
		sumSize += size[i]
		sumBeds += beds[i]
		sumBaths += baths[i]
		sumPrice += price[i]
	}
	fmt.Println()
	fmt.Printf("\n%s\n", strings.Repeat("=", 75))

	fmt.Print(strings.Repeat(" ", 15))
	fmt.Printf("%-15.2f", float64(sumSize)/float64(len(rows)))
	fmt.Printf("%-15.2f", float64(sumBeds)/float64(len(rows)))
	fmt.Printf("%-15.2f", float64(sumBaths)/float64(len(rows)))
	fmt.Printf("%-15.2f", float64(sumPrice)/float64(len(rows)))
	println()
}
