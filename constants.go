package main

const SCREEN_HEIGHT = 450
const SCREEN_WIDTH = 800
const SCREEN_SQUARE = min(SCREEN_HEIGHT, SCREEN_WIDTH)
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
