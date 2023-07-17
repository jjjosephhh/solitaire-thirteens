package textures

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jjjosephhh/solitaire-thirteens/card"
)

type Textures struct {
	Spades     rl.Texture2D
	Clubs      rl.Texture2D
	Diamonds   rl.Texture2D
	Hearts     rl.Texture2D
	Back       rl.Texture2D
	Deck       rl.Texture2D
	CardWidth  float32
	CardHeight float32
	DeckWidth  float32
	DeckHeight float32

	Rect01 rl.Rectangle
	Rect02 rl.Rectangle
	Rect03 rl.Rectangle
	Rect04 rl.Rectangle
	Rect05 rl.Rectangle

	Rect06 rl.Rectangle
	Rect07 rl.Rectangle
	Rect08 rl.Rectangle
	Rect09 rl.Rectangle
	Rect10 rl.Rectangle

	Rect11 rl.Rectangle
	Rect12 rl.Rectangle
	Rect13 rl.Rectangle

	RectDeck rl.Rectangle
}

func LoadTextures() *Textures {
	var tt Textures
	tt.Spades = rl.LoadTexture("assets/spades.png")
	tt.Clubs = rl.LoadTexture("assets/clubs.png")
	tt.Diamonds = rl.LoadTexture("assets/diamonds.png")
	tt.Hearts = rl.LoadTexture("assets/hearts.png")
	tt.Back = rl.LoadTexture("assets/back.png")
	tt.Deck = rl.LoadTexture("assets/deck1.png")
	width := float32(tt.Spades.Width) / 5
	height := float32(tt.Spades.Height) / 3
	tt.CardWidth = width
	tt.CardHeight = height
	tt.Rect01 = rl.NewRectangle(0*width, 0*height, width, height)
	tt.Rect02 = rl.NewRectangle(1*width, 0*height, width, height)
	tt.Rect03 = rl.NewRectangle(2*width, 0*height, width, height)
	tt.Rect04 = rl.NewRectangle(3*width, 0*height, width, height)
	tt.Rect05 = rl.NewRectangle(4*width, 0*height, width, height)

	tt.Rect06 = rl.NewRectangle(0*width, 1*height, width, height)
	tt.Rect07 = rl.NewRectangle(1*width, 1*height, width, height)
	tt.Rect08 = rl.NewRectangle(2*width, 1*height, width, height)
	tt.Rect09 = rl.NewRectangle(3*width, 1*height, width, height)
	tt.Rect10 = rl.NewRectangle(4*width, 1*height, width, height)

	tt.Rect11 = rl.NewRectangle(0*width, 2*height, width, height)
	tt.Rect12 = rl.NewRectangle(1*width, 2*height, width, height)
	tt.Rect13 = rl.NewRectangle(2*width, 2*height, width, height)

	deckWidth := float32(tt.Deck.Width / 3)
	deckHeight := float32(tt.Deck.Height)
	tt.DeckWidth = deckWidth
	tt.DeckHeight = deckHeight
	tt.RectDeck = rl.NewRectangle(deckWidth, 0, deckWidth, deckHeight)

	return &tt
}

func (tt *Textures) UnloadTextures() {
	rl.UnloadTexture(tt.Spades)
	rl.UnloadTexture(tt.Clubs)
	rl.UnloadTexture(tt.Diamonds)
	rl.UnloadTexture(tt.Hearts)
	rl.UnloadTexture(tt.Back)
}

func (tt *Textures) FetchTexture(cur *card.Card) (texture rl.Texture2D, rect rl.Rectangle) {
	if cur.Show {
		var texture rl.Texture2D
		switch cur.Suit {
		case card.Clubs:
			texture = tt.Clubs
		case card.Diamonds:
			texture = tt.Diamonds
		case card.Hearts:
			texture = tt.Hearts
		case card.Spades:
			texture = tt.Spades
		default:
			texture = tt.Back
		}
		var rect rl.Rectangle
		switch cur.Num {
		case 1:
			rect = tt.Rect01
		case 2:
			rect = tt.Rect02
		case 3:
			rect = tt.Rect03
		case 4:
			rect = tt.Rect04
		case 5:
			rect = tt.Rect05
		case 6:
			rect = tt.Rect06
		case 7:
			rect = tt.Rect07
		case 8:
			rect = tt.Rect08
		case 9:
			rect = tt.Rect09
		case 10:
			rect = tt.Rect10
		case 11:
			rect = tt.Rect11
		case 12:
			rect = tt.Rect12
		case 13:
			rect = tt.Rect13
		default:
			rect = tt.Rect01
		}
		return texture, rect
	}
	return tt.Back, tt.Rect01
}
