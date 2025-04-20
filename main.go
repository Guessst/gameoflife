package main

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const SCREEN_HEIGHT = 450
const SCREEN_WIDTH = 800
const SCREEN_SQUARE = min(SCREEN_HEIGHT, SCREEN_WIDTH)
const BOARD_DIM = 40

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

var directions = []struct{ dir_i, dir_j int }{
	{-1, 0},  // Up
	{-1, -1}, // UpLeft
	{0, -1},  // Left
	{1, -1},  // DownLeft
	{1, 0},   // Down
	{1, 1},   // DownRight
	{0, 1},   // Right
	{-1, 1},  // UpRight
}

func is_in_bounds(i int, j int) bool {
	return 0 <= i && i < BOARD_DIM && 0 <= j && j < BOARD_DIM
}

func count_neighbours(i int, j int, b *Board) int {
	count := 0
	for dir := UP; dir < NULLDIRECTION; dir++ {
		dir_i := directions[dir].dir_i
		dir_j := directions[dir].dir_j

		cell_i := i + dir_i
		cell_j := j + dir_j
		ib := is_in_bounds(cell_i, cell_j)
		if ib && b[cell_i*BOARD_DIM+cell_j] {
			count += 1
		}
	}
	return count
}

func new_generation(b Board) Board {
	new_b := Board{}

	for i := 0; i < BOARD_DIM; i++ {
		for j := 0; j < BOARD_DIM; j++ {
			index := i*BOARD_DIM + j

			cell_alive := b[index]
			count := count_neighbours(i, j, &b)
			if cell_alive {
				if count < 2 {
					new_b[index] = false
				} else if count == 2 || count == 3 {
					new_b[index] = true
				} else {
					new_b[index] = false
				}
			} else if count == 3 {
				new_b[index] = true
			}
		}

		/*
			1. Any live cell with fewer than two live neighbours dies, as if by underpopulation.
			2. Any live cell with two or three live neighbours lives on to the next generation.
			3. Any live cell with more than three live neighbours dies, as if by overpopulation.
				4. Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
		*/
		// rule 1

	}

	return new_b
}

func initState() Board {
	b := Board{}

	top := (5 * BOARD_DIM) + BOARD_DIM/2
	b[(0*BOARD_DIM)+top] = true
	b[(1*BOARD_DIM)+top] = true
	b[(2*BOARD_DIM)+top] = true
	// b[(0*BOARD_DIM)+top] = true
	// b[(1*BOARD_DIM)+top] = true
	// b[(1*BOARD_DIM)+top-1] = true
	// b[(1*BOARD_DIM)+top+1] = true

	return b
}

func main() {

	board := initState()

	const timestep_in_sec float32 = 0.5
	reload := timestep_in_sec

	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "raylib [core] example - basic window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(120)

	for !rl.WindowShouldClose() {
		reload -= rl.GetFrameTime()
		fmt.Println("RELOAD::", reload)
		if reload <= 0 {
			reload = timestep_in_sec
			board = new_generation(board)
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		// rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		rect_w := int32(math.Round(float64(SCREEN_SQUARE/BOARD_DIM) * 0.9))
		rect_h := int32(math.Round(float64(SCREEN_SQUARE/BOARD_DIM) * 0.9))

		left_padding := (SCREEN_WIDTH - SCREEN_SQUARE) / 2
		top_padding := (SCREEN_HEIGHT - SCREEN_SQUARE) / 2

		for i := int32(0); i < BOARD_DIM; i++ {
			for j := int32(0); j < BOARD_DIM; j++ {
				x := j*(SCREEN_SQUARE/BOARD_DIM) + int32(left_padding)
				y := i*(SCREEN_SQUARE/BOARD_DIM) + int32(top_padding)

				index := i*BOARD_DIM + j
				cell := board[index]
				if cell {
					rl.DrawRectangle(x, y, rect_w, rect_h, rl.RayWhite)
				} else {
					rl.DrawRectangle(x, y, rect_w, rect_h, rl.Gray)
				}

			}
		}

		rl.DrawFPS(SCREEN_WIDTH-100, 50)
		rl.EndDrawing()
	}
}
