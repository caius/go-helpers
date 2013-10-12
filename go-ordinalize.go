package main

import (
	"fmt"
)

func Ordinalize(num int) string {
	by_100 := num % 100
	fmt_string := "%dth" // Default case

	// See if it should be anything else
	if 11 == by_100 || 12 == by_100 || 13 == by_100 {
		fmt_string = "%dth"
	} else {
		by_10 := num % 10
		if by_10 == 1 {
			fmt_string = "%dst"
		} else if by_10 == 2 {
			fmt_string = "%dnd"
		} else if by_10 == 3 {
			fmt_string = "%drd"
		}
	}

	return fmt.Sprintf(fmt_string, num)
}

func main() {
	fmt.Printf("%2d: %s\n", 1, Ordinalize(1))
	fmt.Printf("%2d: %s\n", 2, Ordinalize(2))
	fmt.Printf("%2d: %s\n", 3, Ordinalize(3))
	fmt.Printf("%2d: %s\n", 5, Ordinalize(5))
	fmt.Printf("%2d: %s\n", 10, Ordinalize(10))
	fmt.Printf("%2d: %s\n", 50, Ordinalize(50))
	fmt.Printf("%2d: %s\n", 51, Ordinalize(51))
	fmt.Printf("%2d: %s\n", 52, Ordinalize(52))
	fmt.Printf("%2d: %s\n", 53, Ordinalize(53))
	fmt.Printf("%2d: %s\n", 54, Ordinalize(54))
	fmt.Printf("%2d: %s\n", 100, Ordinalize(100))
}
