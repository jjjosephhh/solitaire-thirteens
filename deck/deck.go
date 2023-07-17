package deck

import (
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jjjosephhh/solitaire-thirteens/card"
	"github.com/jjjosephhh/solitaire-thirteens/constants"
)

type Deck struct {
	InDeck  []*card.Card
	InPlay  []*card.Card
	Matched []*card.Card
	Pos     rl.Vector2
}

func NewDeck(cardWidth, cardHeight float32) *Deck {
	d := &Deck{
		Pos: rl.NewVector2(0, 0),
	}
	for i := 1; i <= 13; i++ {
		d.InDeck = append(d.InDeck, card.NewCard(card.Clubs, i, cardWidth, cardHeight))
		d.InDeck = append(d.InDeck, card.NewCard(card.Diamonds, i, cardWidth, cardHeight))
		d.InDeck = append(d.InDeck, card.NewCard(card.Hearts, i, cardWidth, cardHeight))
		d.InDeck = append(d.InDeck, card.NewCard(card.Spades, i, cardWidth, cardHeight))
	}
	for _, card := range d.InDeck {
		card.CurPos.X = d.Pos.X
		card.CurPos.Y = d.Pos.Y
	}
	d.Shuffle()
	return d
}

func (d *Deck) Shuffle() {
	// Fisher-Yates shuffle algorithm
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.InDeck), func(i, j int) {
		d.InDeck[i], d.InDeck[j] = d.InDeck[j], d.InDeck[i]
	})
}

func (d *Deck) Draw() (*card.Card, bool) {
	if len(d.InDeck) == 0 {
		return nil, false
	}
	c := d.InDeck[len(d.InDeck)-1]
	d.InDeck = d.InDeck[:len(d.InDeck)-1]
	return c, true
}

func (d *Deck) IsThirteen(c1, c2 *card.Card) []rl.Vector2 {
	var empty []rl.Vector2
	if c1 == nil && c2 == nil {
		return empty
	}
	if c1 == nil && c2.Num == 13 {
		return d.MoveToMatched(c2)
	}
	if c2 == nil && c1.Num == 13 {
		return d.MoveToMatched(c1)
	}
	if c1 != nil && c2 != nil && c1.Num+c2.Num == 13 {
		return d.MoveToMatched(c1, c2)
	}
	return empty
}

func (d *Deck) MoveToMatched(cardsToMove ...*card.Card) []rl.Vector2 {
	var empty []rl.Vector2
	if len(cardsToMove) == 0 {
		return empty
	}

	var i int
	for _, cardInPlay := range d.InPlay {
		skip := false
		for _, cardToMove := range cardsToMove {
			if cardInPlay == cardToMove {
				skip = true
				break
			}
		}
		if skip {
			continue
		}
		d.InPlay[i] = cardInPlay
		i++
	}
	d.InPlay = d.InPlay[:len(d.InPlay)-len(cardsToMove)]
	d.Matched = append(d.Matched, cardsToMove...)

	for _, c := range cardsToMove {
		c.NextPos.X = 0
		c.NextPos.Y = constants.TOP_OFFSET + 2*(c.Height+constants.SPACING_V)
		empty = append(empty, c.CurPos)
	}
	return empty
}
