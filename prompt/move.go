package prompt

import (
	"bufio"
	"fmt"
	"os"
)

func Move() (command string, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nPlease enter your move (or type 'help' for more information): ")

	command, err = reader.ReadString('\n')
	if err != nil {
		return command, err
	}

	return command, err
}

func Help() {
	fmt.Println("\n##### RULES #####")
	fmt.Println("\nMoves should be in the form 'x, y' (without the quotes), where x is the row and y is the column.")
	fmt.Println("\nThe values should be 0-index (meaning the first row/column starts at 0).")
	fmt.Println("\nAn example move would be '2, 1'.")
	fmt.Println("If there was a ship there it would look like this:")
	fmt.Println("----")
	fmt.Println("----")
	fmt.Println("-X--")
	fmt.Println("----")
	fmt.Println("\nIf you made the same move '2, 1', but there was not ship, it would look like this:")
	fmt.Println("----")
	fmt.Println("----")
	fmt.Println("-O--")
	fmt.Println("----")
}
