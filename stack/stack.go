package stack

import (
	"errors"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jjjosephhh/solitaire-thirteens/card"
	"github.com/jjjosephhh/solitaire-thirteens/constants"
)

type Stack struct {
	Cards []*card.Card
	Pos   rl.Vector2
}

func NewStack(x, y float32) *Stack {
	s := &Stack{}
	s.Pos.X = x
	s.Pos.Y = y
	return s
}

func (s *Stack) Pop() (*card.Card, error) {
	if len(s.Cards) == 0 {
		return nil, errors.New("deck unassigned pile is empty")
	}
	c := s.Cards[len(s.Cards)-1]
	s.Cards = s.Cards[:len(s.Cards)-1]
	return c, nil
}
func (s *Stack) Push(c *card.Card) {
	if c == nil {
		return
	}
	s.Cards = append(s.Cards, c)
}

func (s *Stack) CalcOverlap(c *card.Card) float32 {
	numCards := len(s.Cards)
	if numCards == 0 {
		return 0
	}
	x1 := s.Pos.X
	x2 := x1 + c.Width
	y1 := s.Pos.Y
	y2 := y1 + float32(constants.STACK_OFFSET*(numCards-1)) + c.Height

	x3 := c.CurPos.X
	x4 := x3 + c.Width
	y3 := c.CurPos.Y
	y4 := y3 + c.Height

	area1 := (x2 - x1) * (y2 - y1)
	area2 := (x4 - x3) * (y4 - y3)

	max_x1 := x1
	if x3 > max_x1 {
		max_x1 = x3
	}
	max_y1 := y1
	if y3 > max_y1 {
		max_y1 = y3
	}
	min_x2 := x2
	if x4 < x2 {
		min_x2 = x4
	}
	min_y2 := y2
	if y4 < min_y2 {
		min_y2 = y4
	}

	if min_x2 < max_x1 || min_y2 < max_y1 {
		return 0
	}

	overlap_area := (min_x2 - max_x1) * (min_y2 - max_y1)
	percent_overlap := (overlap_area / (area1 + area2)) * 100

	return percent_overlap

}

func (s *Stack) IsEmpty() bool {
	return len(s.Cards) == 0
}
func (s *Stack) Peek() *card.Card {
	if s.IsEmpty() {
		return nil
	}
	return s.Cards[len(s.Cards)-1]
}
