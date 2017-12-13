package main

import (
	"reflect"
	"testing"
)

func TestReadFile(t *testing.T) {
	result := readFile("test.txt")
	// 0 <-> 2
	expected := program{id: 0, pipes: []int{2}}
	if !reflect.DeepEqual(result[0], expected) {
		t.Fatalf("Read File: %v != %v", result[0], expected)
	}
	// 2 <-> 0, 3, 4
	expected = program{id: 2, pipes: []int{0, 3, 4}}
	if !reflect.DeepEqual(result[2], expected) {
		t.Fatalf("Read File: %v != %v", result[2], expected)
	}
}

func TestPart1(t *testing.T) {
	result := part1(readFile("test.txt"))
	expected := 6
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Part 1: %v != %v", result, expected)
	}
}
