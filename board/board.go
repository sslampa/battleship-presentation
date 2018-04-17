package board

import (
	"fmt"

	"github.com/sslampa/battleship-presentation/board/ship"
)

type BoardGame struct {
	Board    Board
	Fleet    ship.Fleet
	Missiles int
}

type Board [4][4]string

func Create() *BoardGame {
	var b Board
	b[0] = [4]string{"-", "-", "-", "-"}
	b[1] = [4]string{"-", "-", "-", "-"}
	b[2] = [4]string{"-", "-", "-", "-"}
	b[3] = [4]string{"-", "-", "-", "-"}

	fleet := ship.CreateFleet([]int{3, 2, 1}, len(b))

	boardGame := BoardGame{
		Board:    b,
		Fleet:    fleet,
		Missiles: 10,
	}

	return &boardGame
}

func (b *BoardGame) Print() {
	fmt.Println("The current board is:")
	for _, row := range b.Board {
		fmt.Println(row)
	}

	fmt.Printf("You have %d missiles left!\n", b.Missiles)
}
