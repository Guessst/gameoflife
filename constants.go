package main

var MODE = "2D"

const SCREEN_H = 600
const SCREEN_W = 800
const SCREEN_SQUARE = min(SCREEN_H, SCREEN_W) // todo: pensar num nome melhor pra isso

const RECT_SCALE = 0.9
const SCREEN_SQUARE_LEFT_PADDING = (SCREEN_W - SCREEN_SQUARE) / 2 // todo: fazer esse + 5 ser metade de um retangulo, ou algo assim
const SCREEN_SQUARE_TOP_PADDING = (SCREEN_H - SCREEN_SQUARE) / 2

var FULL_RECT_W = float64(SCREEN_SQUARE / BOARD_DIM)
var FULL_RECT_H = float64(SCREEN_SQUARE / BOARD_DIM)

var RECT_W = int32(FULL_RECT_W * RECT_SCALE)
var RECT_H = int32(FULL_RECT_H * RECT_SCALE)

const BOARD_DIM = 40
const TIMESTEP_IN_SECS float32 = 0.5

type Board2D [BOARD_DIM * BOARD_DIM]bool
type Board3D [BOARD_DIM * BOARD_DIM * BOARD_DIM]bool
type Index2D struct {
	i int
	j int
}
type Index3D struct {
	i int
	j int
	k int
}

type Direction2D int

const (
	UP_2D Direction2D = iota
	UPLEFT_2D
	LEFT_2D
	DOWNLEFT_2D
	DOWN_2D
	DOWNRIGHT_2D
	RIGHT_2D
	UPRIGHT_2D
	NULLDIRECTION_2D
)

var DIRECTIONS_2D = []Index2D{
	{-1, 0},  // Up
	{-1, -1}, // UpLeft
	{0, -1},  // Left
	{1, -1},  // DownLeft
	{1, 0},   // Down
	{1, 1},   // DownRight
	{0, 1},   // Right
	{-1, 1},  // UpRight
}

type Direction3D int

const (
	UPLEFTFRONT_3D Direction3D = iota
	UPFRONT_3D
	UPRIGHTFRONT_3D
	LEFTFRONT_3D
	FRONT_3D
	RIGHTFRONT_3D
	DOWNLEFTFRONT_3D
	DOWNFRONT_3D
	DOWNRIGHTFRONT_3D

	UPLEFT_3D
	UP_3D
	UPRIGHT_3D
	LEFT_3D
	// CENTER is skipped (0,0,0)
	RIGHT_3D
	DOWNLEFT_3D
	DOWN_3D
	DOWNRIGHT_3D

	UPLEFTBACK_3D
	UPBACK_3D
	UPRIGHTBACK_3D
	LEFTBACK_3D
	BACK_3D
	RIGHTBACK_3D
	DOWNLEFTBACK_3D
	DOWNBACK_3D
	DOWNRIGHTBACK_3D
	NULLDIRECTION_3D
)

var DIRECTIONS_3D = []Index3D{
	{-1, -1, -1}, // UPLEFTFRONT
	{-1, 0, -1},  // UPFRONT
	{-1, 1, -1},  // UPRIGHTFRONT
	{0, -1, -1},  // LEFTFRONT
	{0, 0, -1},   // FRONT
	{0, 1, -1},   // RIGHTFRONT
	{1, -1, -1},  // DOWNLEFTFRONT
	{1, 0, -1},   // DOWNFRONT
	{1, 1, -1},   // DOWNRIGHTFRONT

	{-1, -1, 0}, // UPLEFT
	{-1, 0, 0},  // UP
	{-1, 1, 0},  // UPRIGHT
	{0, -1, 0},  // LEFT
	// {0,  0,  0}, // CENTER (skipped)
	{0, 1, 0},  // RIGHT
	{1, -1, 0}, // DOWNLEFT
	{1, 0, 0},  // DOWN
	{1, 1, 0},  // DOWNRIGHT

	{-1, -1, 1}, // UPLEFTBACK
	{-1, 0, 1},  // UPBACK
	{-1, 1, 1},  // UPRIGHTBACK
	{0, -1, 1},  // LEFTBACK
	{0, 0, 1},   // BACK
	{0, 1, 1},   // RIGHTBACK
	{1, -1, 1},  // DOWNLEFTBACK
	{1, 0, 1},   // DOWNBACK
	{1, 1, 1},   // DOWNRIGHTBACK
}
