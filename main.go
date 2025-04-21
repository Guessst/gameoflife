package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

/*
TODO
- Pausar direito
- Permitir voltar ou avançar passo enquanto pausado
- Carregar estado inicial baseado em arquivo
- GUI (Reload, Dimensões)
- Salvar histórico para arquivo
*/

func is_in_bounds(i int, j int) bool {
	return 0 <= i && i < BOARD_DIM && 0 <= j && j < BOARD_DIM
}

func count_neighbours(i int, j int, b *Board) int {
	count := 0
	for dir := UP; dir < NULLDIRECTION; dir++ {
		dir_i := DIRECTIONS[dir].dir_i
		dir_j := DIRECTIONS[dir].dir_j

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
			/*
				1. Any live cell with fewer than two live neighbours dies, as if by underpopulation.
				2. Any live cell with two or three live neighbours lives on to the next generation.
				3. Any live cell with more than three live neighbours dies, as if by overpopulation.
				4. Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
			*/
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
	}

	return new_b
}

func render(board *Board, paused bool, hovering Index) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)
	for i := int32(0); i < BOARD_DIM; i++ {
		for j := int32(0); j < BOARD_DIM; j++ {
			x := j*(SCREEN_SQUARE/BOARD_DIM) + int32(SCREEN_SQUARE_LEFT_PADDING)
			y := i*(SCREEN_SQUARE/BOARD_DIM) + int32(SCREEN_SQUARE_TOP_PADDING)

			index := i*BOARD_DIM + j
			cell := board[index]
			if cell {
				rl.DrawRectangle(x, y, RECT_W, RECT_H, rl.RayWhite)
			} else {
				rl.DrawRectangle(x, y, RECT_W, RECT_H, rl.Gray)
			}
		}
	}
	rl.DrawFPS(SCREEN_W-90, 50)
	if paused {
		rl.DrawText("PAUSED", SCREEN_W-90, 100, 20, rl.Purple)
	}
	if hovering.i > -1 {
		i := int32(hovering.i)
		j := int32(hovering.j)

		x := j*(SCREEN_SQUARE/BOARD_DIM) + int32(SCREEN_SQUARE_LEFT_PADDING)
		y := i*(SCREEN_SQUARE/BOARD_DIM) + int32(SCREEN_SQUARE_TOP_PADDING)
		purple_minus_opacity := rl.Purple
		purple_minus_opacity.A /= 2
		rl.DrawRectangle(x, y, RECT_W, RECT_H, purple_minus_opacity)

	}
	rl.EndDrawing()
}

func process_input(board *Board, hovering *Index) {
	pos := rl.GetMousePosition()
	start_x := SCREEN_SQUARE_LEFT_PADDING
	start_y := SCREEN_SQUARE_TOP_PADDING
	is_in_grid_area_x := start_x <= int(pos.X) && int(pos.X) <= start_x+SCREEN_SQUARE
	is_in_grid_area_y := start_y <= int(pos.Y) && int(pos.Y) <= start_y+SCREEN_SQUARE
	is_in_grid_area := is_in_grid_area_x && is_in_grid_area_y

	if is_in_grid_area {
		nearest_j := int((pos.X - float32(start_x)) / float32(FULL_RECT_W))
		nearest_i := int((pos.Y - float32(start_y)) / float32(FULL_RECT_H))
		//fmt.Println(nearest_i, nearest_j)
		index := nearest_i*BOARD_DIM + nearest_j

		pressed_left := rl.IsMouseButtonPressed(rl.MouseButtonLeft)
		pressed_right := rl.IsMouseButtonPressed(rl.MouseButtonRight)
		if pressed_left || pressed_right {
			if pressed_left {
				board[index] = true
			} else { // pressed_right
				board[index] = false
			}
		} else { // hack
			hovering.i = nearest_i
			hovering.j = nearest_j
		}
	} else {
		hovering.i = -1
		hovering.j = -1
	}
}

func main() {
	// raylib init window
	rl.InitWindow(SCREEN_W, SCREEN_H, "Game of Life")
	defer rl.CloseWindow()
	rl.SetTargetFPS(0)

	// initial state
	board := Board{}
	reload := TIMESTEP_IN_SECS
	paused := false
	hovering := Index{-1, -1}

	for !rl.WindowShouldClose() {
		if !paused {
			if rl.IsKeyDown(rl.KeySpace) {
				paused = true
				continue
			}

			// simulate
			reload -= rl.GetFrameTime()
			if reload <= 0 {
				reload = TIMESTEP_IN_SECS + reload
				board = new_generation(board)
			}
		} else {
			if rl.IsKeyReleased(rl.KeySpace) {
				paused = false
				hovering.i = -1
				hovering.j = -1
				continue
			}

			process_input(&board, &hovering)
		}

		// render
		render(&board, paused, hovering)
	}
}
