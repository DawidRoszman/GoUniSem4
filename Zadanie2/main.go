package main

import (
	"fmt"
	"math/rand"
	"time"
)

var EMPTY, TREE, FIRE = "  ", "🌲", "🔥"

type Coordinates struct {
	x, y int
}

func (c Coordinates) String() string {
	return fmt.Sprintf("x: %d, y: %d", c.x, c.y)
}

type Forest [][]int

func (f *Forest) createForest(width, height int) {
	*f = make([][]int, height)
	for i := range *f {
		(*f)[i] = make([]int, width)
	}
}

func (f Forest) String() string {
	result := ""
	for i := range f {
		for j := range f[i] {
			switch f[i][j] {
			case 2:
				result += fmt.Sprint(FIRE)
			case 1:
				result += fmt.Sprint(TREE)
			default:
				result += fmt.Sprint(EMPTY)

			}
		}
		result += "\n"
	}
	return result
}

func (f *Forest) populateForest(probablityOfTree int) {
	for i := 0; i < len(*f); i++ {
		for j := 0; j < len((*f)[0]); j++ {
			if rand.Intn(100) < probablityOfTree {
				(*f)[i][j] = 1
			}
		}
	}
}

func (f *Forest) isTreeOnPosition(coordinates Coordinates) (bool, error) {
	if coordinates.x >= len(*f) || coordinates.x < 0 || coordinates.y >= len((*f)[0]) || coordinates.y < 0 {
		return false, fmt.Errorf("out of bounds")
	}
	return (*f)[coordinates.x][coordinates.y] == 1, nil
}

func (f *Forest) findAdjacentTrees(coordinates Coordinates) []Coordinates {
	allAdjacent := make([]Coordinates, 0)
	coordinatesToCheck := [...]Coordinates{
		{coordinates.x - 1, coordinates.y - 1},
		{coordinates.x - 1, coordinates.y},
		{coordinates.x - 1, coordinates.y + 1},
		{coordinates.x, coordinates.y - 1},
		{coordinates.x, coordinates.y + 1},
		{coordinates.x + 1, coordinates.y - 1},
		{coordinates.x + 1, coordinates.y},
		{coordinates.x + 1, coordinates.y + 1},
	}
	for current_coordinates := range coordinatesToCheck {
		isTreeOnPosition, err := f.isTreeOnPosition(coordinatesToCheck[current_coordinates])
		if err != nil {
			continue
		}
		if isTreeOnPosition {
			allAdjacent = append(allAdjacent, coordinatesToCheck[current_coordinates])
		}
	}

	return allAdjacent
}

// func (f *Forest) burnAdjacent(coordinates Coordinates) {
// 	adjacientCoordinates := f.findAdjacentTrees(coordinates)
//
// 	for id := range adjacientCoordinates {
// 		pos := adjacientCoordinates[id]
// 		x, y := pos.x, pos.y
// 		(*f)[x][y] = 2
// 		time.Sleep(time.Millisecond * 200)
// 		f.burnAdjacent(pos)
// 	}
// 	fmt.Println(f)
// }

func (f *Forest) burnAdjacent(coordinates Coordinates, t time.Duration, display bool) {
	queue := []Coordinates{coordinates}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		adjacents := f.findAdjacentTrees(current)
		for _, pos := range adjacents {
			x, y := pos.x, pos.y
			if (*f)[x][y] == 1 {
				(*f)[x][y] = 2
				time.Sleep(t)
				queue = append(queue, pos)
			}
		}
		if display {
			fmt.Println(f)
		}
	}
}

func (f *Forest) checkState() float32 {
	numOfFires, numOfTreesLeft := 0.0, 0.0
	for row := range *f {
		for column := range (*f)[row] {
			switch (*f)[row][column] {
			case 2:
				numOfFires += 1
			case 1:
				numOfTreesLeft += 1
			}
		}
	}
	percentage := numOfFires / (numOfFires + numOfTreesLeft)
	return float32(percentage)
}

func getThundarCoordinates(width, height int) Coordinates {
	return Coordinates{rand.Intn(width), rand.Intn(height)}
}

func testOptimalAfforestation() (int, float32) {
	forestWidth, forestHeight := 26, 16
	burntTreesPercentage := map[int]float32{}
	for j := 0; j < 100; j += 1 {
		for i := 30; i <= 100; i += 1 {
			var forestState float32 = 0.0
			for forestState == 0.0 {
				forest := Forest{}
				forest.createForest(forestWidth, forestHeight)
				forest.populateForest(i)
				forest.burnAdjacent(getThundarCoordinates(forestWidth, forestHeight), time.Millisecond*0, false)
				forestState = forest.checkState()
				if burntTreesPercentage[i] < forestState {
					burntTreesPercentage[i] = forestState
				}
			}
		}
	}
	fmt.Println(burntTreesPercentage)
	return findMinInMap(burntTreesPercentage)
}

func findMinInMap(mapToCheck map[int]float32) (int, float32) {
	minKey := -1
	var minVal float32 = -1.0
	for key, val := range mapToCheck {
		if minKey == -1 {
			minKey = key
			minVal = val
		}
		if val < minVal {
			minKey = key
			minVal = val
		}
	}
	return minKey, minVal
}

func main() {
	forest := Forest{}
	forest.createForest(26, 16)
	forest.populateForest(40)
	fmt.Println(forest)
	forest.burnAdjacent(Coordinates{1, 2}, time.Millisecond*0, false)
	fmt.Printf("Trees burnt: %.2f %%\n", forest.checkState()*100)

	minKey, minVal := testOptimalAfforestation()

	fmt.Printf("Optimal afforestation is %d %% with %.2f %% burnt trees\n", minKey, minVal*100)
}
