package main

import (
	"apc/pkg"
	"fmt"
)

var hopper pkg.Hopper
var reader pkg.InputReader

//Initialize the Hopper and Reader to start
func setup() {
	hopper = pkg.NewHopper()
	reader = pkg.NewReader()
}

func main() {
	setup()
	fmt.Println("Enter number of test cases, followed by all inputs required")
	numberOfCases := reader.ReadInteger()
	userInputs := getInputs(numberOfCases)
	Run(userInputs)
}

//Gets all inputs to be used
func getInputs(numberOfCases int) []pkg.InputsModel {
	var userInputs []pkg.InputsModel
	for i := 0; i < numberOfCases; i++ {
		userInputs = append(userInputs, reader.Readinputs())
	}
	return userInputs
}

//Extracts parameters required for calculating the hops
func extractParameters(input pkg.InputsModel) ([]int, []int, [][]int, int, int) {
	start := []int{input.StartFinish[0], input.StartFinish[1]}
	finish := []int{input.StartFinish[2], input.StartFinish[3]}
	blockedSet := pkg.FormatObstacles(input.ObstacleSlice)
	gridX := input.GridSize[0]
	gridY := input.GridSize[1]
	return start, finish, blockedSet, gridX, gridY
}

//Runs the main solution
func Run(userInputs []pkg.InputsModel) {
	for _, input := range userInputs {
		start, finish, blockedSet, gridX, gridY := extractParameters(input)
		found, numberOfHops, err := hopper.LeastNumberOfHops(blockedSet, start, finish, gridX, gridY)
		hopper.DisplayOutput(found, numberOfHops, err)
	}
}
