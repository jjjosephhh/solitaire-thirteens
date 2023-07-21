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
	windowWidth := 440 + int32(6*constants.SPACING_H)
	windowHeight := int32(2*372/3 + 3*constants.SPACING_V)
	rl.InitWindow(windowWidth, windowHeight, "raylib [core] example - basic window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	tt := textures.LoadTextures()
	defer tt.UnloadTextures()

	d := deck.NewDeck(tt.CardWidth, tt.CardHeight, windowWidth, windowHeight)
	d.Draw10()
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
				canClick = false
				break
			}
		}
		if canClick {
			for _, c := range d.Matched {
				if c.InMotion() {
					canClick = false
					break
				}
			}
		}

		if canClick && rl.IsMouseButtonDown(rl.MouseLeftButton) {
			mousePos = rl.GetMousePosition()
			if !leftMouseDown {
				if len(d.InDeck) == 0 && len(d.InPlay) == 0 {
					d = deck.NewDeck(tt.CardWidth, tt.CardHeight, windowWidth, windowHeight)
					d.Draw10()
					continue
				}

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

		if clickedCard1 != nil && clickedCard2 != nil {
			clickedCard1.Exploading = true
			now := rl.GetTime()
			clickedCard1.ExploadingTimeCur = now
			clickedCard1.ExploadingTimeNext = now + constants.FRAME_TIME_EXPLOSION
			clickedCard1.ExploadingFrame = 0
			clickedCard1 = nil

			clickedCard2.Exploading = true
			clickedCard2.ExploadingTimeCur = now
			clickedCard2.ExploadingTimeNext = now + constants.FRAME_TIME_EXPLOSION
			clickedCard2.ExploadingFrame = 0
			clickedCard2 = nil
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

			if c.Exploading {
				now := rl.GetTime()
				c.ExploadingTimeCur = now
				if c.ExploadingTimeCur >= c.ExploadingTimeNext {
					c.ExploadingTimeNext = now + constants.FRAME_TIME_EXPLOSION
					c.ExploadingFrame++
					fmt.Println("-->", c.ExploadingTimeCur, c.ExploadingTimeNext, c.ExploadingFrame)
				}
				if c.ExploadingFrame >= 16 {
					c.Exploading = false
				} else {
					var rectExplosion rl.Rectangle
					switch c.ExploadingFrame {
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
					fmt.Println(rectExplosion)
					posExplosion := rl.NewVector2(
						c.CurPos.X+c.Width/2-tt.ExplosionWidth/2,
						c.CurPos.Y+c.Height/2-tt.ExplosionHeight/2,
					)
					rl.DrawTextureRec(tt.Explosion, rectExplosion, posExplosion, rl.White)
				}
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

		// rl.DrawTextureRec(tt.Back, tt.Rect01, constants.POSITION_DECK, rl.White)
		if len(d.InPlay) == 0 {
			textWidth := rl.MeasureText(constants.TEXT_WIN, constants.TEXT_SIZE_WIN)
			rl.DrawText(
				constants.TEXT_WIN,
				(windowWidth-textWidth)/2,
				(windowHeight-constants.TEXT_SIZE_WIN)/2,
				constants.TEXT_SIZE_WIN,
				rl.White,
			)
		}

		rl.EndDrawing()
	}
}
