package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sslampa/battleship-presentation/board"
	"github.com/sslampa/battleship-presentation/prompt"
)

func main() {
	// Get player's username
	name, err := prompt.Username()
	if err != nil {
		panic(err)
	}

	// Welcome the player
	prompt.Welcome(name)

	// Create the initial game
	boardGame := board.Create()
	boardGame.Print()

	// Loop through until player loses or wins
	for {
		// Ask the player for their move
		command, err := prompt.Move()
		if err != nil {
			panic(err)
		}

		// Clean the player input
		command = strings.TrimSpace(command)
		split := strings.Split(command, ", ")

		// Check if player input is valid
		if len(split) != 2 && command != "help" {
			fmt.Println("Unknown command. Please try again.")
			continue
		}

		// Print help
		if command == "help" {
			prompt.Help()
		} else {
			// Get the row input
			row, err := strconv.Atoi(split[0])
			if err != nil {
				fmt.Println("Please insert a valid value for row")
				continue
			}

			// Get the column input
			column, err := strconv.Atoi(split[1])
			if err != nil {
				fmt.Println("Please insert a valid value for column")
				continue
			}

			// Check if player hit or missed the ship
			ok := boardGame.CheckMove(row, column)
			if !ok {
				fmt.Println("You've already sent a missile here!")
				continue
			}

			// Check if player won
			win := boardGame.CheckWin()
			if win {
				boardGame.Print()
				fmt.Println("******* WINNER ********")
				break
			}

			// Check if player lost
			loss := boardGame.CheckLoss()
			if loss {
				boardGame.Print()
				fmt.Println("///// LOSER //////")
				break
			}

			// Print current board state
			boardGame.Print()
		}
	}
}
