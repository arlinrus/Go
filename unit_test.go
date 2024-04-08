package main

import "testing"

func TestMax(t *testing.T) {
	//Arrange
	numbers := []int{2, 3, 85, 93, 7, -93}
	expected := 93

	//Act
	result := Max(numbers)

	//Assert
	if result != expected {
		t.Errorf("Incorrest result. Expected %d , got %d",
			expected, result)
	}
}

func TestIndexMax(t *testing.T) {
	//Arrange
	testTable := []struct {
		numbers  []int
		expected int
	}{
		{
			numbers:  []int{4, 45, 48, 98, 5},
			expected: 0,
		},
		{
			numbers:  []int{2, 3, 85, 93, 9, 15},
			expected: -18,
		},
		{
			numbers:  []int{},
			expected: 0,
		},
	}
	//Act
	for _, testCase := range testTable {
		result := Max(testCase.numbers)

		t.Logf("Calling Max(%v), result %d\n", testCase.numbers, result)

		//Assert
		if result != testCase.expected {
			t.Errorf("Incorrest result. Expected %d , got %d", testCase.expected, result)
		}

	}
}
