package main

import (
	"reflect"
	"testing"
)

func TestReadFile(t *testing.T) {
	result := readFile("test.txt")
	expected := instruction{register: "b", operation: "inc", amount: 5,
		condition: condition{register: "a", logic: ">", amount: 1}}
	if !reflect.DeepEqual(result[0], expected) {
		t.Fatalf("Read File: %v != %v", result[0], expected)
	}

	expected = instruction{register: "c", operation: "inc", amount: -20,
		condition: condition{register: "c", logic: "==", amount: 10}}
	if !reflect.DeepEqual(result[3], expected) {
		t.Fatalf("Read File: %v != %v", result[3], expected)
	}
}

func TestConditionMet(t *testing.T) {
	registers := make(map[string]int)
	result := conditionMet(&registers, condition{"a", "==", 0})
	if !result {
		t.Fatalf("conditionMet: %v != true", result)
	}
}

func TestPart1(t *testing.T) {
	result := part1(readFile("test.txt"))
	expected := 1
	if result != expected {
		t.Fatalf("Part 1: %v != %v", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := part2(readFile("test.txt"))
	expected := 10
	if result != expected {
		t.Fatalf("Part 2: %v != %v", result, expected)
	}
}
