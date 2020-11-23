// ---------------------------------------------------------
// EXERCISE: Swap
//
//  Using funcs, swap values through pointers, and swap
//  the addresses of the pointers.
//
//
//  1- Swap the values through a func
//
//     a- Declare two float variables
//
//     b- Declare a func that can swap the variables' values
//        through their memory addresses
//
//     c- Pass the variables' addresses to the func
//
//     d- Print the variables
//
//
//  2- Swap the addresses of the float pointers through a func
//
//     a- Declare two float pointer variables and,
//        assign them the addresses of float variables
//
//     b- Declare a func that can swap the addresses
//        of two float pointers
//
//     c- Pass the pointer variables to the func
//
//     d- Print the addresses, and values that are
//        pointed by the pointer variables
//
// ---------------------------------------------------------

package main

import "fmt"

func main() {
	a, b := 1.1, 2.2
	swap(&a, &b)
	fmt.Printf("a : %g - b : %g\n", a, b)

	ptra, ptrb := &a, &b
	ptra, ptrb = swapAddr(ptra, ptrb)
	fmt.Printf("ptra: %p - ptrb: %p\n", ptra, ptrb)
	fmt.Printf("ptra: %g - ptrb: %g\n", *ptra, *ptrb)

}

func swap(a, b *float64) {
	*a, *b = *b, *a
}

func swapAddr(a, b *float64) (*float64, *float64) {
	return b, a
}
