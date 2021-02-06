package part1

import "testing"

func TestMakePolicy(t *testing.T) {
	input := "1-4 q: abcqrs"
	expected := Policy{1, 4, "q", "abcqrs"}
	calculated := MakePolicy(input)

	if expected != *calculated {
		t.Errorf("Struct mismatch! Expected: %+v, got: %+v", expected, calculated)
	}
}

func TestIsPolicyValid(t *testing.T) {
	p1 := Policy{1, 4, "q", "q"}
	if true != IsPolicyValid(p1) {
		t.Errorf("Mismatch! Expected %t, got %t", true, IsPolicyValid(p1))
	}
	p2 := Policy{1, 4, "q", ""}
	if false != IsPolicyValid(p2) {
		t.Errorf("Mismatch! Expected %t, got %t", false, IsPolicyValid(p2))
	}
	p3 := Policy{1, 4, "q", "qqqqq"}
	if false != IsPolicyValid(p2) {
		t.Errorf("Mismatch! Expected %t, got %t", false, IsPolicyValid(p3))
	}
}

func TestMain(t *testing.T) {
	input_data := Main("../input_test.txt")
	if input_data != 2 {
		t.Errorf("Mismatch! Expected %d, got %d", 2, input_data)
	}
}
