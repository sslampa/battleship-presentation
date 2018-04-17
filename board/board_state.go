package board

import (
	"fmt"

	"github.com/sslampa/battleship-presentation/board/ship"
)

func (b *BoardGame) CheckMove(r, c int) bool {
	ok := b.checkValid(r, c)
	if !ok {
		return false
	}

	for _, s := range b.Fleet {
		hit := b.checkHit(r, c, s)
		if hit {
			fmt.Println("You got a hit!")
			b.placeMarker(r, c, hit)
			break
		}

		b.placeMarker(r, c, hit)
	}

	return true
}

func (b *BoardGame) CheckWin() bool {
	won := true
	for i, s := range b.Fleet {
		if s.Sunk {
			continue
		}

		cr := s.StartRow
		cc := s.StartCol
		for j := 0; j < s.Length; j++ {
			if b.Board[cr][cc] != "X" {
				won = false
				break
			}

			cr, cc = ship.AddInDirection(cr, cc, s.Direction)

			if j == s.Length {
				b.Fleet[i].Sunk = true
			}
		}
	}

	return won
}

func (b *BoardGame) CheckLoss() bool {
	b.Missiles -= 1
	if b.Missiles == 0 {
		return true
	}

	return false
}

func (b *BoardGame) checkValid(r, c int) bool {
	return b.Board[r][c] == "-"
}

func (b *BoardGame) checkHit(r, c int, s ship.Ship) bool {
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

func (b *BoardGame) placeMarker(r, c int, hit bool) {
	if hit {
		b.Board[r][c] = "X"
	} else {
		b.Board[r][c] = "O"
	}
}
