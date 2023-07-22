package constants

import rl "github.com/gen2brain/raylib-go/raylib"

const TARGET_NUM = 13
const SPEED_FACTOR float32 = 2
const STACK_OFFSET int = 20
const TOP_OFFSET float32 = 5
const SPACING_H float32 = 5
const SPACING_V float32 = 5
const SPEED_CARD float32 = 25
const FRAMES_PER_SEC_EXPLOSION = 16
const FRAME_TIME_EXPLOSION = 0.5 / FRAMES_PER_SEC_EXPLOSION

const TEXT_WIN = "You win! Click to play again."
const TEXT_SIZE_WIN = 26

const TEXT_RESTART = "Impossible to make 13! Click to play again."
const TEXT_SIZE_RESTART = 22

var COLOR_RESTART = rl.NewColor(0, 0, 0, 180)
var POSITION_DECK rl.Vector2 = rl.NewVector2(0, TOP_OFFSET)
