package mathUtils

import "testing"

func TestSquare(t *testing.T) {
	if Square(2) != 4 {
		t.Error("Expected 2 square to equal 4")
	}
}

func TestSquareFalse(t *testing.T) {
	if Square(12) == 143 {
		t.Error("Expected 12 square not equal 143")
	}
}

func TestTableSquare(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{2, 4},
		{3, 9},
		{10, 100},
		{25, 625},
	}

	for _, test := range tests {
		if output := Square(test.input); output != test.expected {
			t.Error("test failed input: {}, expected: {}, output: {}", test.input, test.expected, output)
		}
	}
}
