package card

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type CardSuit int

const (
	Clubs CardSuit = iota
	Diamonds
	Hearts
	Spades
)

type Card struct {
	CurPos  rl.Vector2
	NextPos rl.Vector2
	Suit    CardSuit
	Num     int
	Show    bool
	Width   float32
	Height  float32
}

func NewCard(suit CardSuit, num int, cardWidth, cardHeight float32) *Card {
	return &Card{
		Suit:   suit,
		Num:    num,
		Show:   false,
		Width:  cardWidth,
		Height: cardHeight,
	}
}

func (c *Card) InMotion() bool {
	if c.NextPos.X == -10000 && c.NextPos.Y == -10000 {
		return false
	}
	return c.CurPos.X != c.NextPos.X || c.CurPos.Y != c.NextPos.Y
}
