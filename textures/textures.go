package textures

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jjjosephhh/solitaire-thirteens/card"
)

type Textures struct {
	Spades          rl.Texture2D
	Clubs           rl.Texture2D
	Diamonds        rl.Texture2D
	Hearts          rl.Texture2D
	Back            rl.Texture2D
	Deck            rl.Texture2D
	Explosion       rl.Texture2D
	CardWidth       float32
	CardHeight      float32
	DeckWidth       float32
	DeckHeight      float32
	ExplosionWidth  float32
	ExplosionHeight float32

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

	RectExplosion01 rl.Rectangle
	RectExplosion02 rl.Rectangle
	RectExplosion03 rl.Rectangle
	RectExplosion04 rl.Rectangle
	RectExplosion05 rl.Rectangle
	RectExplosion06 rl.Rectangle
	RectExplosion07 rl.Rectangle
	RectExplosion08 rl.Rectangle
	RectExplosion09 rl.Rectangle
	RectExplosion10 rl.Rectangle
	RectExplosion11 rl.Rectangle
	RectExplosion12 rl.Rectangle
	RectExplosion13 rl.Rectangle
	RectExplosion14 rl.Rectangle
	RectExplosion15 rl.Rectangle
	RectExplosion16 rl.Rectangle
}

func LoadTextures() *Textures {
	var tt Textures
	tt.Spades = rl.LoadTexture("assets/spades.png")
	tt.Clubs = rl.LoadTexture("assets/clubs.png")
	tt.Diamonds = rl.LoadTexture("assets/diamonds.png")
	tt.Hearts = rl.LoadTexture("assets/hearts.png")
	tt.Back = rl.LoadTexture("assets/back.png")
	tt.Deck = rl.LoadTexture("assets/deck1.png")
	tt.Explosion = rl.LoadTexture("assets/explosion01.png")
	tt.CardWidth = float32(tt.Spades.Width) / 5
	tt.CardHeight = float32(tt.Spades.Height) / 3
	tt.Rect01 = rl.NewRectangle(0*tt.CardWidth, 0*tt.CardHeight, tt.CardWidth, tt.CardHeight)
	tt.Rect02 = rl.NewRectangle(1*tt.CardWidth, 0*tt.CardHeight, tt.CardWidth, tt.CardHeight)
	tt.Rect03 = rl.NewRectangle(2*tt.CardWidth, 0*tt.CardHeight, tt.CardWidth, tt.CardHeight)
	tt.Rect04 = rl.NewRectangle(3*tt.CardWidth, 0*tt.CardHeight, tt.CardWidth, tt.CardHeight)
	tt.Rect05 = rl.NewRectangle(4*tt.CardWidth, 0*tt.CardHeight, tt.CardWidth, tt.CardHeight)

	tt.Rect06 = rl.NewRectangle(0*tt.CardWidth, 1*tt.CardHeight, tt.CardWidth, tt.CardHeight)
	tt.Rect07 = rl.NewRectangle(1*tt.CardWidth, 1*tt.CardHeight, tt.CardWidth, tt.CardHeight)
	tt.Rect08 = rl.NewRectangle(2*tt.CardWidth, 1*tt.CardHeight, tt.CardWidth, tt.CardHeight)
	tt.Rect09 = rl.NewRectangle(3*tt.CardWidth, 1*tt.CardHeight, tt.CardWidth, tt.CardHeight)
	tt.Rect10 = rl.NewRectangle(4*tt.CardWidth, 1*tt.CardHeight, tt.CardWidth, tt.CardHeight)

	tt.Rect11 = rl.NewRectangle(0*tt.CardWidth, 2*tt.CardHeight, tt.CardWidth, tt.CardHeight)
	tt.Rect12 = rl.NewRectangle(1*tt.CardWidth, 2*tt.CardHeight, tt.CardWidth, tt.CardHeight)
	tt.Rect13 = rl.NewRectangle(2*tt.CardWidth, 2*tt.CardHeight, tt.CardWidth, tt.CardHeight)

	tt.DeckWidth = float32(tt.Deck.Width / 3)
	tt.DeckHeight = float32(tt.Deck.Height)
	tt.RectDeck = rl.NewRectangle(tt.DeckWidth, 0, tt.DeckWidth, tt.DeckHeight)

	tt.ExplosionWidth = float32(tt.Explosion.Width) / 4
	tt.ExplosionHeight = float32(tt.Explosion.Height) / 4
	fmt.Println("Explosion:", tt.ExplosionWidth, tt.ExplosionHeight)
	tt.RectExplosion01 = rl.NewRectangle(0*tt.ExplosionWidth, 0*tt.ExplosionHeight, tt.ExplosionWidth, tt.ExplosionHeight)
	tt.RectExplosion02 = rl.NewRectangle(1*tt.ExplosionWidth, 0*tt.ExplosionHeight, tt.ExplosionWidth, tt.ExplosionHeight)
	tt.RectExplosion03 = rl.NewRectangle(2*tt.ExplosionWidth, 0*tt.ExplosionHeight, tt.ExplosionWidth, tt.ExplosionHeight)
	tt.RectExplosion04 = rl.NewRectangle(3*tt.ExplosionWidth, 0*tt.ExplosionHeight, tt.ExplosionWidth, tt.ExplosionHeight)
	tt.RectExplosion05 = rl.NewRectangle(0*tt.ExplosionWidth, 1*tt.ExplosionHeight, tt.ExplosionWidth, tt.ExplosionHeight)
	tt.RectExplosion06 = rl.NewRectangle(1*tt.ExplosionWidth, 1*tt.ExplosionHeight, tt.ExplosionWidth, tt.ExplosionHeight)
	tt.RectExplosion07 = rl.NewRectangle(2*tt.ExplosionWidth, 1*tt.ExplosionHeight, tt.ExplosionWidth, tt.ExplosionHeight)
	tt.RectExplosion08 = rl.NewRectangle(3*tt.ExplosionWidth, 1*tt.ExplosionHeight, tt.ExplosionWidth, tt.ExplosionHeight)
	tt.RectExplosion09 = rl.NewRectangle(0*tt.ExplosionWidth, 2*tt.ExplosionHeight, tt.ExplosionWidth, tt.ExplosionHeight)
	tt.RectExplosion10 = rl.NewRectangle(1*tt.ExplosionWidth, 2*tt.ExplosionHeight, tt.ExplosionWidth, tt.ExplosionHeight)
	tt.RectExplosion11 = rl.NewRectangle(2*tt.ExplosionWidth, 2*tt.ExplosionHeight, tt.ExplosionWidth, tt.ExplosionHeight)
	tt.RectExplosion12 = rl.NewRectangle(3*tt.ExplosionWidth, 2*tt.ExplosionHeight, tt.ExplosionWidth, tt.ExplosionHeight)
	tt.RectExplosion13 = rl.NewRectangle(0*tt.ExplosionWidth, 3*tt.ExplosionHeight, tt.ExplosionWidth, tt.ExplosionHeight)
	tt.RectExplosion14 = rl.NewRectangle(1*tt.ExplosionWidth, 3*tt.ExplosionHeight, tt.ExplosionWidth, tt.ExplosionHeight)
	tt.RectExplosion15 = rl.NewRectangle(2*tt.ExplosionWidth, 3*tt.ExplosionHeight, tt.ExplosionWidth, tt.ExplosionHeight)
	tt.RectExplosion16 = rl.NewRectangle(3*tt.ExplosionWidth, 3*tt.ExplosionHeight, tt.ExplosionWidth, tt.ExplosionHeight)

	return &tt
}

func (tt *Textures) UnloadTextures() {
	rl.UnloadTexture(tt.Spades)
	rl.UnloadTexture(tt.Clubs)
	rl.UnloadTexture(tt.Diamonds)
	rl.UnloadTexture(tt.Hearts)
	rl.UnloadTexture(tt.Back)
	rl.UnloadTexture(tt.Deck)
	rl.UnloadTexture(tt.Explosion)
}

func (tt *Textures) FetchCardTexture(cur *card.Card) (texture rl.Texture2D, rect rl.Rectangle) {
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

func (tt *Textures) DrawExplosion(frame int, pos *rl.Vector2, dim *rl.Vector2) {
	if frame >= 16 {
		return
	}
	var rectExplosion rl.Rectangle
	switch frame {
	case 0:
		rectExplosion = tt.RectExplosion01
	case 1:
		rectExplosion = tt.RectExplosion02
	case 2:
		rectExplosion = tt.RectExplosion03
	case 3:
		rectExplosion = tt.RectExplosion04
	case 4:
		rectExplosion = tt.RectExplosion05
	case 5:
		rectExplosion = tt.RectExplosion06
	case 6:
		rectExplosion = tt.RectExplosion07
	case 7:
		rectExplosion = tt.RectExplosion08
	case 8:
		rectExplosion = tt.RectExplosion09
	case 9:
		rectExplosion = tt.RectExplosion10
	case 10:
		rectExplosion = tt.RectExplosion11
	case 11:
		rectExplosion = tt.RectExplosion12
	case 12:
		rectExplosion = tt.RectExplosion13
	case 13:
		rectExplosion = tt.RectExplosion14
	case 14:
		rectExplosion = tt.RectExplosion15
	case 15:
		rectExplosion = tt.RectExplosion16
	}
	posExplosion := rl.NewVector2(
		pos.X+dim.X/2-tt.ExplosionWidth/2,
		pos.Y+dim.Y/2-tt.ExplosionHeight/2,
	)
	rl.DrawTextureRec(tt.Explosion, rectExplosion, posExplosion, rl.White)
}
