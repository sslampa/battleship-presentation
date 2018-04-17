package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sslampa/battleship-presentation/first/board"
	"github.com/sslampa/battleship-presentation/first/board/ship"
	"github.com/sslampa/battleship-presentation/first/prompt"
)

func main() {
	name, err := prompt.Username()
	if err != nil {
		panic(err)
	}

	prompt.Welcome(name)
	board := board.Create()
	board.Print()
	fleet := ship.CreateFleet([]int{3, 2, 1}, len(board))
	fmt.Println("Fleet", fleet)

	missiles := 10
	for {
		fmt.Printf("You have %d missiles left!\n", missiles)
		command, err := prompt.Move()
		if err != nil {
			panic(err)
		}
		command = strings.TrimSpace(command)

		split := strings.Split(command, ", ")
		if len(split) != 2 && command != "help" {
			fmt.Println("Unknown command. Please try again.")
			continue
		}

		if command == "help" {
			prompt.Help()
		} else {
			row, err := strconv.Atoi(split[0])
			if err != nil {
				fmt.Println("Please insert a valid value for row")
				continue
			}

			column, err := strconv.Atoi(split[1])
			if err != nil {
				fmt.Println("Please insert a valid value for column")
				continue
			}

			win := board.Check(row, column, fleet)
			board.Print()

			if win {
				fmt.Println("******* WINNER ********")
				break
			}

			missiles--
			if missiles == 0 {
				fmt.Println("///// LOSER //////")
				break
			}
		}
	}
}
