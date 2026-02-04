package main

import (
	"fmt"
)

func main() {
	var input string
	var counter int
	var cords []string

	for true {
		fmt.Scan(&input)

		cords = append(cords, input)

		for y := 0; y < 3; y++ {
			for x := 0; x < 3; x++ {
				var fill bool = false

				for _, cord := range cords {
					var letter byte = cord[0]
					var number byte = cord[1]

					var x_pick int = int(number-'0') - 1
					var y_pick int = int(letter - 'a')

					if y == y_pick && x == x_pick {
						fill = true
						break
					}
				}

				if fill {
					fmt.Print("x ")
				} else {
					fmt.Print("o ")
				}
			}

			fmt.Print("\n")
		}

		counter += 1
	}
}
