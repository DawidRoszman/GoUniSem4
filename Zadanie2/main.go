package main

import (
	"fmt"
	"math/rand"
)

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
			result += fmt.Sprint(f[i][j])
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
	if coordinates.x >= len(*f) || coordinates.x < 0 {
		return false, fmt.Errorf("out of bounds")
	}
	return (*f)[coordinates.y][coordinates.x] == 1, nil
}

func (f *Forest) findAdjacientTrees(coordinates Coordinates) []Coordinates {
	allAdjacient := make([]Coordinates, 0)
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
			allAdjacient = append(allAdjacient, coordinatesToCheck[current_coordinates])
		}
	}

	return allAdjacient
}

func main() {
	forest := Forest{}
	forest.createForest(26, 16)
	forest.populateForest(40)
	fmt.Println(forest)
	adjacient := forest.findAdjacientTrees(Coordinates{1, 1})
	fmt.Println(adjacient)
}
