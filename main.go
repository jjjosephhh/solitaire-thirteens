package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jjjosephhh/solitaire-thirteens/card"
	"github.com/jjjosephhh/solitaire-thirteens/constants"
	"github.com/jjjosephhh/solitaire-thirteens/deck"
	"github.com/jjjosephhh/solitaire-thirteens/textures"
)

func main() {
	rl.InitWindow(int32((440/5+constants.SPACING_H)*6), int32((372/3+constants.SPACING_V)*5), "raylib [core] example - basic window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	tt := textures.LoadTextures()
	defer tt.UnloadTextures()

	d := deck.NewDeck(tt.CardWidth, tt.CardHeight)
	for i := 0; i < 10; i++ {
		c, ok := d.Draw()
		if !ok {
			continue
		}
		c.Show = true
		c.NextPos.X = (1 + float32(i%5)) * (c.Width + constants.SPACING_H)
		c.NextPos.Y = float32(i/5)*(c.Height+constants.SPACING_V) + constants.TOP_OFFSET
		d.InPlay = append(d.InPlay, c)
	}

	var t float32 // time parameter
	var clickedCard1 *card.Card
	var clickedCard2 *card.Card
	var mousePos rl.Vector2
	var leftMouseDown bool
	var canClick bool
	for !rl.WindowShouldClose() {
		t += rl.GetFrameTime()

		canClick = true
		for _, c := range d.InPlay {
			if c.InMotion() {
				fmt.Println("cannot click", c.CurPos, c.NextPos)
				canClick = false
				break
			}
		}
		if canClick {
			for _, c := range d.Matched {
				if c.InMotion() {
					fmt.Println("cannot click", c.CurPos, c.NextPos)
					canClick = false
					break
				}
			}
		}

		if canClick && rl.IsMouseButtonDown(rl.MouseLeftButton) {
			mousePos = rl.GetMousePosition()
			if !leftMouseDown {
				for _, c := range d.InPlay {
					if c.CurPos.X <= mousePos.X && mousePos.X <= c.CurPos.X+tt.CardWidth {
						if c.CurPos.Y <= mousePos.Y && mousePos.Y <= c.CurPos.Y+tt.CardHeight {
							if clickedCard1 == nil {
								clickedCard1 = c
								if emptySpots := d.IsThirteen(clickedCard1, clickedCard2); len(emptySpots) > 0 {
									clickedCard1 = nil
									clickedCard2 = nil
									for _, emptySpot := range emptySpots {
										c, ok := d.Draw()
										if !ok {
											continue
										}
										c.Show = true
										c.NextPos = emptySpot
										d.InPlay = append(d.InPlay, c)
									}
								}
								break
							}

							if c == clickedCard1 {
								clickedCard1 = nil
								break
							}

							if clickedCard2 == nil {
								clickedCard2 = c
								if emptySpots := d.IsThirteen(clickedCard1, clickedCard2); len(emptySpots) > 0 {
									clickedCard1 = nil
									clickedCard2 = nil
									for _, emptySpot := range emptySpots {
										c, ok := d.Draw()
										if !ok {
											continue
										}
										c.Show = true
										c.NextPos = emptySpot
										d.InPlay = append(d.InPlay, c)
									}
								}
								break
							}

							if c == clickedCard2 {
								clickedCard2 = nil
								break
							}
						}
					}
				}
			}
			leftMouseDown = true
		} else {
			leftMouseDown = false
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.NewColor(0, 128, 128, 255))

		for _, c := range d.InPlay {
			if c.CurPos.X == c.NextPos.X && c.CurPos.Y == c.NextPos.Y {
				c.NextPos.X = -10000
				c.NextPos.Y = -10000
			}
			if c.NextPos.X != -10000 {
				direction := rl.Vector2Subtract(c.NextPos, c.CurPos)
				direction = rl.Vector2Normalize(direction)
				amount := rl.NewVector2(direction.X*constants.SPEED_CARD, direction.Y*constants.SPEED_CARD)
				c.CurPos = rl.Vector2Add(c.CurPos, amount)
				distance := rl.Vector2Distance(c.CurPos, c.NextPos)
				if distance <= constants.SPEED_CARD {
					c.CurPos.X = c.NextPos.X
					c.CurPos.Y = c.NextPos.Y
				}
			}
			texture, rect := tt.FetchTexture(c)
			rl.DrawTextureRec(texture, rect, c.CurPos, rl.White)
			if c == clickedCard1 || c == clickedCard2 {
				rl.DrawRectangleLinesEx(
					rl.NewRectangle(
						c.CurPos.X-2,
						c.CurPos.Y-2,
						c.Width+2*2,
						c.Height+2*2,
					),
					4,
					rl.Red,
				)
			}
		}

		for _, c := range d.Matched {
			if !c.Show {
				continue
			}
			if c.CurPos.X == c.NextPos.X && c.CurPos.Y == c.NextPos.Y {
				c.NextPos.X = -10000
				c.NextPos.Y = -10000
			}
			if c.NextPos.X != -10000 {
				direction := rl.Vector2Subtract(c.NextPos, c.CurPos)
				direction = rl.Vector2Normalize(direction)
				amount := rl.NewVector2(direction.X*constants.SPEED_CARD, direction.Y*constants.SPEED_CARD)
				c.CurPos = rl.Vector2Add(c.CurPos, amount)
				distance := rl.Vector2Distance(c.CurPos, c.NextPos)
				if distance <= constants.SPEED_CARD {
					c.CurPos.X = c.NextPos.X
					c.CurPos.Y = c.NextPos.Y
				}
			}
			texture, rect := tt.FetchTexture(c)
			rl.DrawTextureRec(texture, rect, c.CurPos, rl.White)
		}

		rl.DrawTextureRec(tt.Back, tt.Rect01, constants.POSITION_DECK, rl.White)

		rl.EndDrawing()
	}
}
