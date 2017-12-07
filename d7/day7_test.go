package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReadFile(t *testing.T) {
	result := readFile("test.txt")
	expected := program{name: "tknk", weight: 41, parent: "", children: []string{"ugml", "padx", "fwft"}}
	if !reflect.DeepEqual(result["tknk"], expected) {
		t.Fatalf("File read wrong: %v expected %v", result["tknk"], expected)
	}
}

func TestSumWeight(t *testing.T) {
	result := sumWeight(readFile("test.txt"), part1(readFile("test.txt")))
	expected := 778
	if result != expected {
		t.Fatalf("Sum weight wrong: %v expected %v", result, expected)
	}
	result = sumWeight(readFile("test.txt"), "ugml")
	expected = 251
	if result != expected {
		t.Fatalf("Sum weight wrong: %v expected %v", result, expected)
	}
}

func TestIsBalanced(t *testing.T) {
	result, resultDiff := isBalanced(readFile("test.txt"), part1(readFile("test.txt")))
	expected := "ugml"
	if result != expected {
		t.Fatalf("TestIsBalanced: %v expected %v", result, expected)
	}
	expectedDiff := -8
	if resultDiff != expectedDiff {
		t.Fatalf("TestIsBalanced diff1: %v expected %v", result, expected)
	}
}

func TestPart1(t *testing.T) {
	result := part1(readFile("test.txt"))
	expected := "tknk"
	if result != expected {
		t.Fatalf("Part 1 wrong: %v expected %v", result, expected)
	}
	result = part1(nil)
	expected = ""
	if result != expected {
		t.Fatalf("Part 1 wrong: %v expected %v", result, expected)
	}
}

func TestPart2(t *testing.T) {
	data := readFile("test.txt")
	result := part2(data, part1(readFile("test.txt")))
	expected := 60
	if result != expected {
		t.Fatalf("Part 2 ugml wrong: %v expected %v", result, expected)
	}
	fmt.Println()
	tmp := data["ugml"]
	tmp.weight = 60
	data["ugml"] = tmp
	tmp = data["ebii"]
	tmp.weight = 60
	data["ebii"] = tmp
	result = part2(data, part1(readFile("test.txt")))
	expected = 61
	if result != expected {
		t.Fatalf("Part 2 ebii wrong: %v expected %v", result, expected)
	}
}
