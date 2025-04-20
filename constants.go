package main

import "math"

const SCREEN_H = 600
const SCREEN_W = 800
const SCREEN_SQUARE = min(SCREEN_H, SCREEN_W)

const SCREEN_SQUARE_LEFT_PADDING = (SCREEN_W-SCREEN_SQUARE)/2 + 5 // todo: fazer esse + 5 ser metade de um retangulo, ou algo assim
const SCREEN_SQUARE_TOP_PADDING = (SCREEN_H-SCREEN_SQUARE)/2 + 5

var RECT_W = int32(math.Round(float64(SCREEN_SQUARE/BOARD_DIM) * 0.9))
var RECT_H = int32(math.Round(float64(SCREEN_SQUARE/BOARD_DIM) * 0.9))

const BOARD_DIM = 40
const TIMESTEP_IN_SECS float32 = 0.5

type Board [BOARD_DIM * BOARD_DIM]bool

// Direction is a custom type for direction indices
type Direction int

const (
	UP Direction = iota
	UPLEFT
	LEFT
	DOWNLEFT
	DOWN
	DOWNRIGHT
	RIGHT
	UPRIGHT
	NULLDIRECTION
)

var DIRECTIONS = []struct{ dir_i, dir_j int }{
	{-1, 0},  // Up
	{-1, -1}, // UpLeft
	{0, -1},  // Left
	{1, -1},  // DownLeft
	{1, 0},   // Down
	{1, 1},   // DownRight
	{0, 1},   // Right
	{-1, 1},  // UpRight
}
