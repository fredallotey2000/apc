package pkg

import (
	"errors"
	"fmt"
	"testing"
)

//Test implementation for the main unit leastNumberOfHops
func TestLeastNumberOfHops(t *testing.T) {
	hopper := NewHopper()
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			testCase.ActualBool, testCase.ActualNumber, testCase.ActualError = hopper.LeastNumberOfHops(
				testCase.Blocked,
				testCase.Start,
				testCase.Finish,
				testCase.GridX,
				testCase.GridY)
			if testCase.ActualError != nil && testCase.ActualError.Error() != testCase.ExpectedError.Error() {
				t.Errorf("Got %v - Expected %v", testCase.ActualError, testCase.ExpectedError)
			}
			if testCase.ActualBool != testCase.ExpectedBool && testCase.ActualNumber != testCase.ExpectedNumber {
				t.Errorf("Got %t and %d - Expected %t and %d", testCase.ActualBool, testCase.ActualNumber, testCase.ExpectedBool, testCase.ExpectedNumber)
			}
		})
	}
}

//Example test
func ExampleLeastNumberOfHops() {
	hopper := NewHopper()
	var testCase = TestModel{
		Name:           "does not find target",
		Blocked:        [][]int{{1, 0}, {1, 2}, {0, 1}, {2, 1}},
		Start:          []int{0, 0},
		Finish:         []int{2, 2},
		GridX:          3,
		GridY:          3,
		ExpectedBool:   false,
		ExpectedNumber: 1,
		ExpectedError:  nil,
	}
	testCase.ActualBool, testCase.ActualNumber, testCase.ActualError = hopper.LeastNumberOfHops(
		testCase.Blocked,
		testCase.Start,
		testCase.Finish,
		testCase.GridX,
		testCase.GridY)
	fmt.Println(testCase.ActualBool, testCase.ActualNumber, testCase.ActualError)
	// Output: false 1 <nil>
}

// Table test for various test cases
var testCases = []TestModel{
	{
		Name:           "does not find target",
		Blocked:        [][]int{{1, 0}, {1, 2}, {0, 1}, {2, 1}},
		Start:          []int{0, 0},
		Finish:         []int{2, 2},
		GridX:          3,
		GridY:          3,
		ExpectedBool:   false,
		ExpectedNumber: 1,
		ExpectedError:  nil,
	},
	{
		Name:           "find target",
		Blocked:        [][]int{{1, 2}, {4, 3}},
		Start:          []int{4, 0},
		Finish:         []int{4, 4},
		GridX:          5,
		GridY:          5,
		ExpectedBool:   true,
		ExpectedNumber: 7,
		ExpectedError:  nil,
	},
	{
		Name:           "invalid entries",
		Blocked:        [][]int{{1}, {4, 3, 9}},
		Start:          []int{4, 0},
		Finish:         []int{4, 4},
		GridX:          50,
		GridY:          5,
		ExpectedBool:   false,
		ExpectedNumber: 0,
		ExpectedError:  errors.New("invalid entries"),
	},

	{Name: "invalid entries above grid X",
		Blocked:        [][]int{{1, 0}, {4, 3}},
		Start:          []int{4, 0},
		Finish:         []int{4, 4},
		GridX:          50,
		GridY:          5,
		ExpectedBool:   false,
		ExpectedNumber: 0,
		ExpectedError:  errors.New("invalid entries"),
	},
	{

		Name:           "invalid block entries",
		Blocked:        [][]int{{1}, {4, 3, 9}},
		Start:          []int{4, 0},
		Finish:         []int{4, 4},
		GridX:          5,
		GridY:          5,
		ExpectedBool:   false,
		ExpectedNumber: 0,
		ExpectedError:  errors.New("invalid entries"),
	},
	{

		Name:           "invalid entries start point",
		Blocked:        [][]int{{1}, {4, 3, 9}},
		Start:          []int{-4, 0},
		Finish:         []int{4, 4},
		GridX:          50,
		GridY:          5,
		ExpectedBool:   false,
		ExpectedNumber: 0,
		ExpectedError:  errors.New("invalid entries"),
	},

	{
		Name:           "invalid finish point ",
		Blocked:        [][]int{{1}, {4, 3, 9}},
		Start:          []int{4, 0},
		Finish:         []int{4, 49},
		GridX:          50,
		GridY:          5,
		ExpectedBool:   false,
		ExpectedNumber: 0,
		ExpectedError:  errors.New("invalid entries"),
	},
}

//benchmark tests to access performance of solution
func benchmarkLeastNumberOfHops(model TestModel, b *testing.B) {
	hopper := NewHopper()
	for i := 0; i < b.N; i++ {
		ok, _, _ := hopper.LeastNumberOfHops(model.Blocked,
			model.Start,
			model.Finish,
			model.GridX,
			model.GridY)
		if ok {
			fmt.Print("")
		}
	}
}

func BenchmarkLeastNumberOfHops_No_Find(b *testing.B) {
	benchmarkLeastNumberOfHops(TestModel{
		Name:           "does not find target",
		Blocked:        [][]int{{1, 0}, {1, 2}, {0, 1}, {2, 1}},
		Start:          []int{0, 0},
		Finish:         []int{2, 2},
		GridX:          3,
		GridY:          3,
		ExpectedBool:   false,
		ExpectedNumber: 1,
		ExpectedError:  nil,
	}, b)
}

func BenchmarkLeastNumberOfHops_Find(b *testing.B) {
	benchmarkLeastNumberOfHops(TestModel{
		Name:           "find target",
		Blocked:        [][]int{{1, 2}, {4, 3}},
		Start:          []int{4, 0},
		Finish:         []int{4, 4},
		GridX:          5,
		GridY:          5,
		ExpectedBool:   true,
		ExpectedNumber: 7,
		ExpectedError:  nil,
	}, b)
}

func BenchmarkLeastNumberOfHops_Invalid_Entries(b *testing.B) {
	benchmarkLeastNumberOfHops(TestModel{
		Name:           "invalid entries",
		Blocked:        [][]int{{1, 2}, {4, 3}},
		Start:          []int{4, 0},
		Finish:         []int{4, 4},
		GridX:          50,
		GridY:          5,
		ExpectedBool:   false,
		ExpectedNumber: 0,
		ExpectedError:  errors.New("invalid entries"),
	}, b)
}
