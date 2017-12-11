package main

import "fmt"

func main() {
	part1(361527)
	part2(361527)
}

func getNum(n, d int) int {
	return 4*n*n - (11-d)*n + (8 - d)
}

func part1(input int) {
	var n = 1
	for {
		var x, y int

		for d := 0; d < 7; d++ {
			if input >= getNum(n, d) && input < getNum(n, d+1) {
				switch d {
				case 0, 4:
					if input == getNum(n, d) {
						x, y = n-1, 0
					} else {
						x, y = n-1, input-getNum(n, d)
					}
				case 2, 6:
					if input == getNum(n, d) {
						x, y = 0, n-1
					} else {
						x, y = input-getNum(n, d), n-1
					}
				case 1, 3, 5:
					if input == getNum(n, d) {
						x, y = n-1, n-1
					} else {
						x, y = input-getNum(n, d), n-1
					}
				}
				fmt.Println(x, y)
				fmt.Println(x + y)
				return
			}
		}
		if input >= getNum(n, 7) && input < getNum(n+1, 0) {
			if input == getNum(n, 7) {
				x = n
				y = n
			} else {
				if i := input - getNum(n, 7); i > 1 {
					x, y = n, i-1

				} else {
					x, y = n, n-1
				}
			}
			fmt.Println(x, y)
			fmt.Println(x + y)
			return
		}
		n++
	}
}

func part2(input int) {

}
