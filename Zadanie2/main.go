package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var EMPTY, TREE, OLD_TREE, FIRE = "  ", "🌲", "🌳", "🔥"

type Coordinates struct {
	x, y int
}

func (c Coordinates) String() string {
	return fmt.Sprintf("x: %d, y: %d", c.x, c.y)
}

type Forest [][]Spot

type Spot struct {
	symbol int
	old    bool
}

func (f *Forest) createForest(width, height int) {
	*f = make([][]Spot, height)
	for i := range *f {
		(*f)[i] = make([]Spot, width)
	}
}

func (f Forest) String() string {
	result := ""
	for i := range f {
		for j := range f[i] {
			switch f[i][j].symbol {
			case 2:
				result += fmt.Sprint(FIRE)
			case 1:
				if f[i][j].old {
					result += fmt.Sprint(OLD_TREE)
				} else {
					result += fmt.Sprint(TREE)
				}
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
				(*f)[i][j].symbol = 1
				isOld := rand.Intn(100)
				if isOld > 55 {
					(*f)[i][j].old = true
				}
			}
		}
	}
}

func (f *Forest) isTreeOnPosition(coordinates Coordinates) (bool, error) {
	if coordinates.x >= len(*f) || coordinates.x < 0 || coordinates.y >= len((*f)[0]) || coordinates.y < 0 {
		return false, fmt.Errorf("out of bounds")
	}
	return (*f)[coordinates.x][coordinates.y].symbol == 1, nil
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
			if (*f)[x][y].symbol == 1 {
				if (*f)[x][y].old {
					(*f)[x][y].symbol = 2
				} else {
					willBurn := rand.Intn(100)
					if willBurn > 60 {
						(*f)[x][y].symbol = 2
					} else {
						continue
					}
				}
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
			switch (*f)[row][column].symbol {
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
		forestPopulation := rand.Intn(100)
		var forestState float32 = 0.0
		for forestState == 0.0 {
			forest := Forest{}
			forest.createForest(forestWidth, forestHeight)
			forest.populateForest(forestPopulation)
			forest.burnAdjacent(getThundarCoordinates(forestWidth, forestHeight), time.Millisecond*0, false)
			forestState = forest.checkState()
			if burntTreesPercentage[forestPopulation] < forestState {
				burntTreesPercentage[forestPopulation] = forestState
			}
		}
	}
	fmt.Println(burntTreesPercentage)

	bar := charts.NewBar()

	// Set global options
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Forest",
		Subtitle: "Optimal afforestation test",
	}))

	keys := make([]int, len(burntTreesPercentage))

	vals := make([]opts.BarData, len(burntTreesPercentage))

	i := 0
	for key, val := range burntTreesPercentage {
		keys[i] = key
		vals[i] = opts.BarData{
			Value: val * 100,
		}
		i += 1
	}
	sort.Ints(keys)
	// Put data into instance
	bar.SetXAxis(keys).
		AddSeries("Burn percentage", vals)
	f, _ := os.Create("bar.html")
	_ = bar.Render(f)

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
	forest.burnAdjacent(getThundarCoordinates(26, 16), time.Millisecond*200, true)
	fmt.Printf("Trees burnt: %.2f %%\n", forest.checkState()*100)

	minKey, minVal := testOptimalAfforestation()

	fmt.Printf("Optimal afforestation is %d %% with %.2f %% burnt trees\n", minKey, minVal*100)
}
