package part2

import (
	"testing"
)

func TestCheckSumV2(t *testing.T) {
	input1, input2, input3 := 2018, 1, 1
	expectedOutput := []interface{}{2018, 1, 1, 2018, nil}
	calculatedOutput1, calculatedOutput2, calculatedOutput3, calculatedOutput4, calculatedOutput5 := CheckSumV2(input1, input2, input3)
	calculatedOutput := []interface{}{calculatedOutput1, calculatedOutput2, calculatedOutput3, calculatedOutput4, calculatedOutput5}

	if len(expectedOutput) != len(calculatedOutput) {
		t.Errorf("Expected and Calculated have a different size! Expected: %d. Calculated: %d", len(expectedOutput), len(calculatedOutput))
	}

	for i := range calculatedOutput {
		if calculatedOutput[i] != expectedOutput[i] {
			t.Errorf("Output mismatch. Expected %d. Got %d", expectedOutput[i], calculatedOutput[i])
		}
	}
}

func TestMakeTriples(t *testing.T) {
	// function doesn't currently remove duplicates
	input := []int{1, 2, 3, 4}
	expectedOutput := [][]int{{1, 2, 3}, {1, 2, 4}, {1, 3, 2}, {1, 3, 4}, {1, 4, 2}, {1, 4, 3}, {2, 1, 3}, {2, 1, 4}, {2, 3, 1}, {2, 3, 4}, {2, 4, 1}, {2, 4, 3}, {3, 1, 2}, {3, 1, 4}, {3, 2, 1}, {3, 2, 4}, {3, 4, 1}, {3, 4, 2}, {4, 1, 2}, {4, 1, 3}, {4, 2, 1}, {4, 2, 3}, {4, 3, 1}, {4, 3, 2}}
	calculatedOutput := MakeTriples(input)

	if len(expectedOutput) != len(calculatedOutput) {
		t.Errorf("Expected and Calculated have a different size! Expected: %d. Calculated: %d", len(expectedOutput), len(calculatedOutput))
	}

	for i, val := range calculatedOutput {
		for j := range val {
			if calculatedOutput[i][j] != expectedOutput[i][j] {
				t.Error("Output mismatch", "Expected: ", expectedOutput, "Calculated: ", calculatedOutput)
			}
		}
	}
}

func TestMain(t *testing.T) {
	expectedOutput := []int{438, 360, 1222, 192684960}
	co1, co2, co3, co4, calculatedErr := Main("../input_test.txt")
	calculatedOutput := []int{co1, co2, co3, co4}

	if calculatedErr != nil {
		t.Errorf("%s", calculatedErr)
	}
	for i := range calculatedOutput {
		if calculatedOutput[i] != expectedOutput[i] {
			t.Errorf("Output mismatch. Expected %v. Got %v", expectedOutput, calculatedOutput)
			break
		}
	}
}