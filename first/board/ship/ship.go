package ship

import (
	"math/rand"
	"time"
)

type Fleet []Ship

type Ship struct {
	Length    int
	StartRow  int
	StartCol  int
	Direction string
	Sunk      bool
}

func CreateFleet(ships []int, dimension int) Fleet {
	var fleet Fleet
	for _, ship := range ships {
		fleet = append(fleet, createShip(ship, dimension))
	}

	return fleet
}

func AddInDirection(row, column int, direction string) (int, int) {
	switch direction {
	case "right":
		return row, column + 1
	case "left":
		return row, column - 1
	case "up":
		return row - 1, column
	case "down":
		return row + 1, column
	}

	return row, column
}

func createShip(length int, dimension int) Ship {
	var ship Ship
	row, col, direction := createValid(length, dimension)

	ship = Ship{
		Length:    length,
		StartRow:  row,
		StartCol:  col,
		Direction: direction,
		Sunk:      false,
	}

	return ship
}

func createValid(length, dimension int) (row int, col int, direction string) {
	for {
		row, col = getCoordinates(dimension)
		direction = getDirection()

		switch direction {
		case "right":
			if (col + length - 1) >= dimension {
				continue
			}
		case "left":
			if (col - length - 1) < 0 {
				continue
			}
		case "up":
			if (row - length - 1) < 0 {
				continue
			}
		case "down":
			if (row + length - 1) >= dimension {
				continue
			}
		}

		break
	}

	return row, col, direction
}

func getCoordinates(dimension int) (int, int) {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(dimension), rand.Intn(dimension)
}

func getDirection() string {
	rand.Seed(time.Now().UnixNano())
	directions := []string{"left", "right", "up", "down"}
	randNum := rand.Intn(len(directions))
	return directions[randNum]
}
