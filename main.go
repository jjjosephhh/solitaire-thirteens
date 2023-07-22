package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jjjosephhh/solitaire-thirteens/card"
	"github.com/jjjosephhh/solitaire-thirteens/constants"
	"github.com/jjjosephhh/solitaire-thirteens/pile"
	"github.com/jjjosephhh/solitaire-thirteens/textures"
	"github.com/jjjosephhh/solitaire-thirteens/utils"
)

func main() {
	screenWidth := 440 + int32(6*constants.SPACING_H)
	screenHeight := int32(2*372/3 + 3*constants.SPACING_V)
	rl.InitWindow(screenWidth, screenHeight, "Solitaire - Thirteens")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	tt := textures.LoadTextures()
	defer tt.UnloadTextures()

	cardDim := rl.NewVector2(tt.CardWidth, tt.CardHeight)
	posUnplayed := rl.NewVector2(float32(screenWidth)/2-cardDim.X/2, -cardDim.Y)
	posMatched := rl.NewVector2(float32(screenWidth)/2-cardDim.X/2, float32(screenHeight))
	unplayed := pile.NewDeck(&posUnplayed, &cardDim)
	inPlay := &pile.Pile{Cards: unplayed.InitializeInPlay()}
	matched := &pile.Pile{
		Cards:    make([]*card.Card, 0),
		Position: &posMatched,
	}

	var t float32 // time parameter

	var selected1 *card.Card
	var selected2 *card.Card

	var posMouse rl.Vector2
	var leftMouseDown bool
	var canSelect bool
	canPlay := true

	fmt.Println(cardDim, posUnplayed, posMatched, unplayed, inPlay, matched, selected1, selected2, posMouse, leftMouseDown, canSelect, canPlay)
	for !rl.WindowShouldClose() {
		t += rl.GetFrameTime()
		canSelect = true

		for _, c := range inPlay.Cards {
			if c.InMotion() {
				canSelect = false
				break
			}
		}

		if canSelect {
			for _, c := range matched.Cards {
				if c.InMotion() {
					canSelect = false
					break
				}
			}
		}

		if canSelect && rl.IsMouseButtonDown(rl.MouseLeftButton) {
			posMouse = rl.GetMousePosition()
			if !leftMouseDown {
				if unplayed.IsEmpty() && inPlay.IsEmpty() || !canPlay {
					canPlay = true
					unplayed = pile.NewDeck(&posUnplayed, &cardDim)
					inPlay.Cards = unplayed.InitializeInPlay()
					continue
				}

				for _, c := range inPlay.Cards {
					if c.CurPos.X <= posMouse.X && posMouse.X <= c.CurPos.X+tt.CardWidth {
						if c.CurPos.Y <= posMouse.Y && posMouse.Y <= c.CurPos.Y+tt.CardHeight {
							if selected1 == nil {
								selected1 = c
								if matches := utils.IsMatch(selected1, selected2); len(matches) > 0 {
									inPlay.MoveTo(matched, matches)
									selected1 = nil
									selected2 = nil
									for _, matchedCard := range matches {
										c, ok := unplayed.Draw()
										if !ok {
											continue
										}
										c.Show = true
										c.NextPos = matchedCard.CurPos
										inPlay.Cards = append(inPlay.Cards, c)
									}
								}
								break
							}

							if c == selected1 {
								selected1 = nil
								break
							}

							if selected2 == nil {
								selected2 = c
								if matches := utils.IsMatch(selected1, selected2); len(matches) > 0 {
									inPlay.MoveTo(matched, matches)
									selected1 = nil
									selected2 = nil
									for _, matchedCard := range matches {
										c, ok := unplayed.Draw()
										if !ok {
											continue
										}
										c.Show = true
										c.NextPos = matchedCard.CurPos
										inPlay.Cards = append(inPlay.Cards, c)
									}
								}
								break
							}

							if c == selected2 {
								selected2 = nil
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

		if selected1 != nil && selected2 != nil {
			selected1.Exploading = true
			now := rl.GetTime()
			selected1.ExploadingTimeCur = now
			selected1.ExploadingTimeNext = now + constants.FRAME_TIME_EXPLOSION
			selected1.ExploadingFrame = 0
			selected1 = nil

			selected2.Exploading = true
			selected2.ExploadingTimeCur = now
			selected2.ExploadingTimeNext = now + constants.FRAME_TIME_EXPLOSION
			selected2.ExploadingFrame = 0
			selected2 = nil
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.NewColor(0, 128, 128, 255))

		for _, c := range inPlay.Cards {
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
			if c == selected1 || c == selected2 {
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

		for _, c := range matched.Cards {
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

		if len(inPlay.Cards) == 0 {
			textWidth := rl.MeasureText(constants.TEXT_WIN, constants.TEXT_SIZE_WIN)
			rl.DrawText(
				constants.TEXT_WIN,
				(screenWidth-textWidth)/2,
				(screenHeight-constants.TEXT_SIZE_WIN)/2,
				constants.TEXT_SIZE_WIN,
				rl.White,
			)
		}

		if canPlay {
			canPlay = utils.MatchExists(inPlay.Cards)
		}
		if !canPlay {
			rl.DrawRectangle(0, 0, screenWidth, screenHeight, constants.COLOR_RESTART)
			textWidth := rl.MeasureText(constants.TEXT_RESTART, constants.TEXT_SIZE_RESTART)
			rl.DrawText(
				constants.TEXT_RESTART,
				(screenWidth-textWidth)/2,
				(screenHeight-constants.TEXT_SIZE_RESTART)/2,
				constants.TEXT_SIZE_RESTART,
				rl.White,
			)
		}

		rl.EndDrawing()
	}
}
