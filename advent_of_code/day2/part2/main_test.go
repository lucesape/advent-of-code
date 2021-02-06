package part2

import (
	"advent-of-code/advent_of_code/day2/part1"
	"testing"
)

func TestIsPolicyValid(t *testing.T) {
	p1 := part1.Policy{1, 4, "a", "abcde"}
	if true != IsPolicyValid(p1) {
		t.Errorf("Mismatch! Expected %t, got %t", true, IsPolicyValid(p1))
	}
	p2 := part1.Policy{1, 4, "b", "cdefg"}
	if false != IsPolicyValid(p2) {
		t.Errorf("Mismatch! Expected %t, got %t", false, IsPolicyValid(p2))
	}
	p3 := part1.Policy{1, 4, "c", "ccccccccc"}
	if false != IsPolicyValid(p2) {
		t.Errorf("Mismatch! Expected %t, got %t", false, IsPolicyValid(p3))
	}
}

func TestMain(t *testing.T) {
	input_data := Main("../input_test.txt")
	if input_data != 1 {
		t.Errorf("Mismatch! Expected %d, got %d", 2, input_data)
	}
}
