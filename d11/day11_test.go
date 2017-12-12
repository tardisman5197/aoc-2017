package main

import (
	"reflect"
	"testing"
)

func TestReadFile(t *testing.T) {
	result := readFile("test.txt")
	expected := []int{0, 1, 2, 3, 4, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Read File: %v != %v", result, expected)
	}
}

func TestPart1(t *testing.T) {
	result := part1(readFile("test.txt"))
	expected := 0
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Part 1 : %v != %v", result, expected)
	}

	result = part1([]int{1, 1, 1})
	expected = 3
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Part 1 : %v != %v", result, expected)
	}

	result = part1([]int{1, 1, 3, 3})
	expected = 2
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Part 1 : %v != %v", result, expected)
	}

	result = part1([]int{2, 4, 2, 4, 4})
	expected = 3
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Part 1 : %v != %v", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := part2(readFile("test.txt"))
	expected := 2
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Part 2: %v != %v", result, expected)
	}
}
