package main

import (
	"reflect"
	"testing"
)

func TestReadFile(t *testing.T) {
	result := readFile("test.txt")
	expected := scanner{depth: 0, area: 3}
	if !reflect.DeepEqual(result[0], expected) {
		t.Errorf("Read File: %v != %v", result[0], expected)
	}

	expected = scanner{depth: 4, area: 4}
	if !reflect.DeepEqual(result[2], expected) {
		t.Errorf("Read File: %v != %v", result[2], expected)
	}
}

func TestPart1(t *testing.T) {
	input := readFile("test.txt")
	result := part1(&input)
	expected := 24
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Part 1: %v != %v", result, expected)
	}
}

func TestPart2(t *testing.T) {
	input := readFile("test.txt")
	result := part2(&input)
	expected := 10
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Part 1: %v != %v", result, expected)
	}
}
