package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sort"
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
	data := make([][]string, 0)
	s := bufio.NewScanner(f)
	for s.Scan() {
		row := strings.Fields(s.Text())
		data = append(data, row)
	}
	part1(data)
	part2(data)
}

func part1(data [][]string) {
	var numValid = 0
	for _, row := range data {
		var valid = true
		used := make(map[string]bool)
		for _, word := range row {
			if !used[word] {
				used[word] = true
			} else {
				valid = false
			}
		}
		if valid {
			numValid++
		}
	}
	fmt.Println(numValid)
}

func part2(data [][]string) {
	var numValid = 0
	for _, row := range data {
		var valid = true
		used := make(map[string]bool)
		for _, word := range row {
			for k := range used {
				keySplit := strings.Split(k, "")
				sort.Strings(keySplit)
				wordSplit := strings.Split(word, "")
				sort.Strings(wordSplit)
				if reflect.DeepEqual(keySplit, wordSplit) {
					valid = false
					break
				}
			}
			if valid {
				used[word] = true
			} else {
				break
			}
		}
		if valid {
			numValid++
		}
	}
	fmt.Println(numValid)
}
