package main

import (
	"reflect"
	"testing"
)

func TestReadFile(t *testing.T) {
	result := readFile1("test.txt")
	expected := []int{3, 4, 1, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Read File: %v != %v", result, expected)
	}

	resultByte := readFile2("test.txt")
	expectedByte := []byte{'3', ',', '4', ',', '1', ',', '5', 17, 31, 73, 47, 23}
	if !reflect.DeepEqual(resultByte, expectedByte) {
		t.Fatalf("Read File Byte: %v != %v", resultByte, expectedByte)
	}
}

func TestHash(t *testing.T) {
	result := hash([]int{65, 27, 9, 1, 4, 3, 40, 50, 91, 7, 6, 0, 2, 5, 68, 22})
	expected := []int{64}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Hash list: %v != %v", result, expected)
	}

	result = hash([]int{65, 27, 9, 1, 4, 3, 40, 50, 91, 7, 6, 0, 2, 5, 68, 22, 65, 27, 9, 1, 4, 3, 40, 50, 91, 7, 6, 0, 2, 5, 68, 22})
	expected = []int{64, 64}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Hash list: %v != %v", result, expected)
	}
}

func TestHex(t *testing.T) {
	result := encodeHex([]int{64, 7, 255})
	expected := "4007ff"
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Hex conv: %v != %v", result, expected)
	}

}
func TestPart1(t *testing.T) {
	result := part1(5, readFile1("test.txt"))
	expected := 12
	if result != expected {
		t.Fatalf("Read File: %v != %v", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := part2(5, readFile2("test.txt"))
	expected := ""
	if result != expected {
		t.Fatalf("Read File: %v != %v", result, expected)
	}

	result = part2(256, readFile2("input.txt"))
	expected = "a7af2706aa9a09cf5d848c1e6605dd2a"
	if result != expected {
		t.Fatalf("Part 2: %v != %v", result, expected)
	}
}
