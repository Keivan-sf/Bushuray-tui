package list

import "testing"

func TestResetPendingTestResults(t *testing.T) {
	model := Model{
		Items: []ListItem{
			{ProfileId: 1, TestResult: -2},
			{ProfileId: 2, TestResult: 0},
			{ProfileId: 3, TestResult: -1},
			{ProfileId: 4, TestResult: 120},
		},
	}

	model.ResetPendingTestResults()

	want := []int{0, 0, -1, 120}
	for i, item := range model.Items {
		if item.TestResult != want[i] {
			t.Fatalf("item %d test result = %d, want %d", i, item.TestResult, want[i])
		}
	}
}
