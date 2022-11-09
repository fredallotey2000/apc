package pkg

//input model to help collect multiple cases from consold
type InputsModel struct {
	GridSize      []int
	StartFinish   []int
	Obstacles     int
	ObstacleSlice [][]int
}

//Model for testing the LeastNumberOfHops unit
type TestModel struct {
	Name           string
	Blocked        [][]int
	Start          []int
	Finish         []int
	GridX          int
	GridY          int
	ExpectedBool   bool
	ExpectedNumber int
	ActualBool     bool
	ActualNumber   int
	ExpectedError  error
	ActualError    error
}
