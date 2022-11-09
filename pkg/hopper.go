package pkg

import (
	"errors"
	"fmt"
	"math"
)

// Hopper interface that can be implemented by other hopper struct
type Hopper interface {
	LeastNumberOfHops(blocked [][]int, start []int, finish []int, gridX, gridY int) (bool, int, error)
	DisplayOutput(found bool, numberOfHops int, err error)
}

type hopper struct{}

//Creates a new hopper
func NewHopper() Hopper {
	return &hopper{}
}

//Main function for calculating number of hops
func (h *hopper) LeastNumberOfHops(blocked [][]int, start []int, finish []int, gridX, gridY int) (bool, int, error) {
	if h.isValidEntries(blocked, start, finish, gridX, gridY) {
		blockedSet := h.formatBlocked(blocked)
		startSlice := h.formatSlice(start)
		finishSlice := h.formatSlice(finish)
		found, numberOfHops := h.breadthFirstSearch(startSlice, finishSlice, blockedSet, gridX, gridY)
		return found, numberOfHops, nil
	}
	return false, 0, errors.New("invalid entries")
}

//Helper method to convert slices to array
func (h *hopper) formatSlice(a []int) [2]int {
	return [2]int{a[0], a[1]}
}

//Output of the LeastNumberOfHops solution
func (h *hopper) DisplayOutput(found bool, numberOfHops int, err error) {
	if err != nil {
		fmt.Print(err.Error())
	} else if found {
		fmt.Printf("Optimal solution takes %d hops\n", numberOfHops)
	} else {
		fmt.Printf("No solution no way hopper can reach finish point\n")
	}
}

func isValidBlockedSet(blockedSet [][]int, gridX, gridY int) bool {
	for _, v := range blockedSet {
		if len(v) != 2 || v[0] < 0 || v[0] >= gridX || v[1] < 0 || v[1] >= gridY {
			return false
		}
	}
	return true
}

//validate entries passed to least number of hops method
func (h *hopper) isValidEntries(blockedSet [][]int, start, finish []int, gridX, gridY int) bool {
	if gridX >= 1 && gridX <= 30 && gridY >= 1 && gridY <= 30 {
		if len(start) == 2 && len(finish) == 2 {
			if start[0] >= 0 && start[1] < gridX && finish[0] >= 0 && finish[1] < gridY {
				return isValidBlockedSet(blockedSet, gridX, gridY)
			}
		}
	}
	return false
}

//Algorithm to find the fihish point BFS
func (h *hopper) breadthFirstSearch(start [2]int, finish [2]int, blockedSet map[[2]int]bool, gridX, gridY int) (bool, int) {
	//In other not to touch squares in between we do not add {1,1} to the moves
	moves := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var queue [][2]int
	queue = append(queue, start)
	visited := make(map[[2]int]bool)
	initX := start[0]
	initY := start[1]
	minMoves := 0
	for len(queue) > 0 {
		for range queue {
			//dequeue
			first := queue[0]
			queue = queue[1:]
			//Check if we have found the finish square
			if first[0] == finish[0] && first[1] == finish[1] {
				minMoves += 1
				return true, minMoves
			}
			//Apply all moves
			for i := range moves {
				xs := first[0] + moves[i][0]
				ys := first[1] + moves[i][1]
				point := [2]int{xs, ys}
				_, ok1 := visited[point]
				_, ok2 := blockedSet[point]
				//Extra condition to make sure velocity does not exceed 3 or -3
				if (xs >= 0 && xs < gridX && ys >= 0 && ys < gridY && !ok1 && !ok2) && (int(math.Abs(float64(xs-initX))) < 4 || int(math.Abs(float64(ys-initY))) < 4) {
					//add to visited points so we do not visit again
					visited[point] = true
					//enqueue
					queue = append(queue, point)
				}
			}
		}
		minMoves += 1
	}
	return false, minMoves
}

//convert 2D array to map
func (h *hopper) formatBlocked(blocked [][]int) map[[2]int]bool {
	blockedSet := make(map[[2]int]bool)
	for i := range blocked {
		key := h.formatSlice(blocked[i])
		blockedSet[key] = true
	}
	return blockedSet
}

//Extract obstacle indexes the right way
func FormatObstacles(arr [][]int) [][]int {
	var results [][]int
	for _, v := range arr {
		results = append(results, []int{v[0], v[2]})
		results = append(results, []int{v[1], v[3]})
	}
	return results
}
