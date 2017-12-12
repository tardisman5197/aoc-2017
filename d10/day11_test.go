package main

import (
	"reflect"
	"testing"
)

func TestReadFile(t *testing.T) {
	result := readFile("test.txt")
	expected := []int{3, 4, 1, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Read File: %v != %v", result, expected)
	}
}

func TestPart1(t *testing.T) {
	result := part1(5, readFile("test.txt"))
	expected := 12
	if result != expected {
		t.Fatalf("Read File: %v != %v", result, expected)
	}
}
