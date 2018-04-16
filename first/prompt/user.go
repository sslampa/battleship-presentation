package prompt

import (
	"bufio"
	"fmt"
	"os"
)

func Username() (name string, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your name: ")
	name, err = reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return name, nil
}

func Welcome(name string) {
	fmt.Println("Hello,", name)
	fmt.Println("This is Battleship!")
}

func Rules() {
	fmt.Println("You have 10 chances to sink all the battleships!")
	fmt.Println("There are 3 ")
}
