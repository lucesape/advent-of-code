package main

import "testing"

func TestReadFile(t *testing.T) {
	input := "helpers.txt"
	expectedOutput := "this\nis\na\ntest\n"
	calculatedOutput, err := ReadFile(input)
	if err != nil {
		t.Errorf("%s", err)
	}
	if expectedOutput != calculatedOutput {
		t.Errorf("Output mismatch. Expected %s. Got %s", expectedOutput, calculatedOutput)
	}
}

func TestReadFileLineByLine(t *testing.T) {
	input := "helpers.txt"
	expectedOutput := []string{"this", "is", "a", "test"}
	calculatedOutput, err := ReadFileLineByLine(input)
	if err != nil {
		t.Errorf("%s", err)
	}
	for index := range expectedOutput {
		if expectedOutput[index] != calculatedOutput[index] {
			t.Errorf("Output mismatch. Expected %s. Got %s", expectedOutput[index], calculatedOutput[index])
		}
	}

}
