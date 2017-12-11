package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Open File
	f, err := os.Open(os.Args[1])
	check(err)
	defer f.Close()

	// Parse file
	data := make([][]int, 0)
	s := bufio.NewScanner(f)
	for s.Scan() {
		row := strings.Fields(s.Text())
		nums := make([]int, 0)
		for _, num := range row {
			val, _ := strconv.Atoi(num)
			nums = append(nums, val)
		}
		data = append(data, nums)
	}

	part1(data)
	part2(data)
}

func part1(data [][]int) {
	var total int = 0

	for _, row := range data {
		min := math.MaxInt32
		max := -min
		for _, num := range row {
			if num > max {
				max = num
			}
			if num < min {
				min = num
			}
		}
		total += (max - min)
	}
	fmt.Println(total)
}

func part2(data [][]int) {
	var total int = 0
	for _, row := range data {
		for i, num := range row {
			for j := i + 1; j < len(row); j++ {
				if num > row[j] {
					if num%row[j] == 0 {
						total += num / row[j]
						break
					}
				} else {
					if row[j]%num == 0 {
						total += row[j] / num
						break
					}
				}
			}
		}
	}
	fmt.Println(total)
	return
}

func part1Conccurent(data [][]int) {
	checksum := make(chan int, len(data))

	// Calculate checksums
	go func() {
		for _, row := range data {
			go func(row []int) {
				min := math.MaxInt32
				max := -min
				for _, num := range row {
					if num > max {
						max = num
					}
					if num < min {
						min = num
					}
				}
				checksum <- (max - min)
			}(row)
		}
	}()

	// Calculate total
	var total, done int = 0, 0
	for {
		select {
		case i := <-checksum:
			done++
			total += i
			if done >= len(data) {
				fmt.Println(total)
				return
			}
		}
	}
}

func part2Conccurent(data [][]int) {
	checksum := make(chan int, len(data))

	// Calculate checksums
	go func() {
		for _, row := range data {
			go func(row []int) {
				for i, num := range row {
					for j := i + 1; j < len(row); j++ {
						if num > row[j] {
							if num%row[j] == 0 {
								checksum <- num / row[j]
								return
							}
						} else {
							if row[j]%num == 0 {
								checksum <- row[j] / num
								return
							}
						}
					}
				}

			}(row)
		}
	}()

	// Calculate total
	var total, done int = 0, 0
	for {
		select {
		case i := <-checksum:
			done++
			total += i
			if done >= len(data) {
				fmt.Println(total)
				return
			}
		}
	}
}
