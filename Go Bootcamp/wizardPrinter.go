package main

// ---------------------------------------------------------
// EXERCISE: Wizard Printer
//
//   In this exercise, your goal is to display a few famous scientists
//   in a pretty table.
//
//   1. Create a multi-dimensional array
//   2. In each inner array, store the scientist's name, lastname and his/her
//      nickname
//   3. Print their information in a pretty table using a loop.
//
// EXPECTED OUTPUT
//   First Name      Last Name       Nickname
//   ==================================================
//   Albert          Einstein        time
//   Isaac           Newton          apple
//   Stephen         Hawking         blackhole
//   Marie           Curie           radium
//   Charles         Darwin          fittest
// ---------------------------------------------------------

import (
	"fmt"
	"strings"
)

func main() {

	//This was my first approach in which I could not arrange the table properly:

	//	wizards := [3][5]string{
	//		{"Albert", "Isaac", "Stephen", "Marie", "Charles"},
	//		{"Einstein", "Newton", "Hawking", "Curie", "Darwin"},
	//		{"time", "apple", "blackhole", "radium", "fittest"},
	//	}
	//	fmt.Println(`First Name      Last Name       Nickname
	//==================================================`)
	//	for i := 0; i <= 4; i++ {
	//
	//		fmt.Printf("%s\t\t %s\t\t %s\n",wizards[0][i], wizards[1][i],wizards[2][i])
	//	}
	//	fmt.Println(" ---------------------------------------------------------")

	wizards := [...][3]string{
		{"First Name", "Last Name", "Nickname"},
		{"Albert", "Einstein", "time"},
		{"Isaac", "Newton", "apple"},
		{"Stephen", "Hawking", "blackhole"},
		{"Marie", "Curie", "radium"},
		{"Charles", "Darwin", "fittest"},
	}

	for i := 0; i <= len(wizards)-1; i++ {
		n := wizards[i]
		fmt.Printf("%-15s %-15s %-15s\n", n[0], n[1], n[2])
		if i == 0 {
			fmt.Println(strings.Repeat("=", 50))
		}
	}
	fmt.Println(strings.Repeat("-", 50))
}
