package board

import (
	"fmt"

	"github.com/sslampa/battleship-presentation/first/board/ship"
)

type Board [4][4]string

func Create() *Board {
	var b Board
	b[0] = [4]string{"-", "-", "-", "-"}
	b[1] = [4]string{"-", "-", "-", "-"}
	b[2] = [4]string{"-", "-", "-", "-"}
	b[3] = [4]string{"-", "-", "-", "-"}

	return &b
}

func (b *Board) Print() {
	fmt.Println("The current board is:")
	for _, row := range b {
		fmt.Println(row)
	}
}

func (b *Board) Check(r, c int, f ship.Fleet) bool {
	ok := b.checkValid(r, c)
	if !ok {
		return false
	}

	for _, s := range f {
		fmt.Println(s)
		hit := b.checkHit(r, c, s)
		if hit {
			fmt.Println("You got a hit!")
			b.placeMarker(r, c, hit)
			break
		}
		b.placeMarker(r, c, hit)
	}

	win := b.checkWin(f)
	if win {
		return true
	}

	return false
}

func (b *Board) checkValid(r, c int) bool {
	return b[r][c] == "-"
}

func (b *Board) checkHit(r, c int, s ship.Ship) bool {
	if s.Direction == "right" {
		cr := s.StartCol + s.Length - 1
		if c >= s.StartCol && c <= cr && r == s.StartRow {
			return true
		}
	} else if s.Direction == "left" {
		cr := s.StartCol - s.Length - 1
		if c >= cr && c <= s.StartCol && r == s.StartRow {
			return true
		}
	} else if s.Direction == "up" {
		rr := s.StartRow - s.Length - 1
		if r >= rr && r <= s.StartRow && c == s.StartCol {
			return true
		}
	} else if s.Direction == "down" {
		rr := s.StartRow + s.Length - 1
		if r >= s.StartRow && r <= rr && c == s.StartCol {
			return true
		}
	}

	return false
}

func (b *Board) checkWin(f ship.Fleet) bool {
	won := true
	for i, s := range f {
		if s.Sunk {
			continue
		}

		cr := s.StartRow
		cc := s.StartCol
		for j := 0; j < s.Length; j++ {
			if b[cr][cc] != "X" {
				won = false
				break
			}

			cr, cc = ship.AddInDirection(cr, cc, s.Direction)

			if j == s.Length {
				f[i].Sunk = true
			}
		}
	}

	return won
}

func (b *Board) checkLoss() bool {
	return false
}

func (b *Board) placeMarker(r, c int, hit bool) {
	fmt.Println(hit)
	if hit {
		b[r][c] = "X"
	} else {
		b[r][c] = "O"
	}
}
