package main

import (
	"reflect"
	"testing"
)

func TestReadFile(t *testing.T) {
	result := readFile("test.txt")
	expected := []int{0, 2, 7, 0}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Read File wrong: %d expeceted %d", result, expected)
	}
}

func TestPart1(t *testing.T) {
	data := readFile("test.txt")
	result, _, _ := part1(data)
	if result != 5 {
		t.Fatalf("Part 1 wrong: %d expected 5", result)
	}
}

func TestPart2(t *testing.T) {
	data := readFile("test.txt")
	result := part2(data)
	if result != 4 {
		t.Fatalf("Part 2 wrong: %d expected 4", result)
	}
}
