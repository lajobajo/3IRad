package main

import (
	"fmt"
)

func main() {
	var input string

	fmt.Scan(&input)

	var letter byte = input[0]
	var number byte = input[1]

	var x_pick int = int(number-'0') - 1
	var y_pick int = int(letter - 'a')

	for y := 2; y >= 0; y-- {
		for x := 0; x < 3; x++ {
			if y == y_pick && x == x_pick {
				fmt.Print("x ")
			} else {
				fmt.Print("o ")
			}

		}

		fmt.Print("\n")
	}
}
