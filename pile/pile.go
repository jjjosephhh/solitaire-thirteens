package pile

import (
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jjjosephhh/solitaire-thirteens/card"
	"github.com/jjjosephhh/solitaire-thirteens/constants"
)

type Pile struct {
	Cards    []*card.Card
	Position *rl.Vector2
}

func NewDeck(position *rl.Vector2, cardDim *rl.Vector2) *Pile {
	p := &Pile{
		Cards:    make([]*card.Card, 0),
		Position: position,
	}
	for i := 1; i <= 13; i++ {
		p.Cards = append(
			p.Cards,
			card.NewCard(card.Clubs, i, cardDim.X, cardDim.Y),
			card.NewCard(card.Diamonds, i, cardDim.X, cardDim.Y),
			card.NewCard(card.Hearts, i, cardDim.X, cardDim.Y),
			card.NewCard(card.Spades, i, cardDim.X, cardDim.Y),
		)
	}
	for _, card := range p.Cards {
		card.CurPos = *p.Position
	}
	p.Shuffle()
	return p
}

func (p *Pile) Size() int {
	return len(p.Cards)
}

func (p *Pile) IsEmpty() bool {
	return p.Size() == 0
}

func (p *Pile) InitializeInPlay() []*card.Card {
	var drawn []*card.Card
	for i := 0; i < 10; i++ {
		c, ok := p.Draw()
		if !ok {
			continue
		}
		c.Show = true
		c.NextPos = rl.NewVector2(
			float32(i%5)*c.Width+float32(1+i%5)*constants.SPACING_H,
			float32(i/5)*(c.Height+constants.SPACING_V)+constants.TOP_OFFSET,
		)
		drawn = append(drawn, c)
	}
	return drawn
}

func (p *Pile) Shuffle() {
	// Fisher-Yates shuffle algorithm
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(p.Cards), func(i, j int) {
		p.Cards[i], p.Cards[j] = p.Cards[j], p.Cards[i]
	})
}

func (p *Pile) Draw() (*card.Card, bool) {
	if len(p.Cards) == 0 {
		return nil, false
	}
	end := len(p.Cards) - 1
	c := p.Cards[end]
	p.Cards = p.Cards[:end]
	return c, true
}

func (p *Pile) MoveTo(targetPile *Pile, cards []*card.Card) []*card.Card {
	var moved []*card.Card
	if len(cards) == 0 {
		return moved
	}

	var i int
	for _, cardInPlay := range p.Cards {
		skip := false
		for _, cardToMove := range cards {
			if cardInPlay == cardToMove {
				skip = true
				break
			}
		}
		if skip {
			continue
		}
		p.Cards[i] = cardInPlay
		i++
	}
	p.Cards = p.Cards[:len(p.Cards)-len(cards)]
	targetPile.Cards = append(targetPile.Cards, cards...)

	for _, c := range cards {
		c.NextPos = *targetPile.Position
		moved = append(moved, c)
	}
	return moved
}
