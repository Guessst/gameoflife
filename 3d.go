package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// TODO
// - refatorar e organizar
// - opção de cor

func get_index_3d(i int, j int, k int) int {
	return (i * BOARD_DIM * BOARD_DIM) + (j * BOARD_DIM) + k
}

func is_in_bounds_3d(i int, j int, k int) bool {
	return 0 <= i && i < BOARD_DIM && 0 <= j && j < BOARD_DIM && 0 <= k && k < BOARD_DIM
}

func count_neighbours_3d(i int, j int, k int, b *Board3D) int {
	count := 0
	for dir := UPLEFTFRONT_3D; dir < NULLDIRECTION_3D; dir++ {
		dir_i := DIRECTIONS_3D[dir].i
		dir_j := DIRECTIONS_3D[dir].j
		dir_k := DIRECTIONS_3D[dir].k

		cell_i := i + dir_i
		cell_j := j + dir_j
		cell_k := k + dir_k
		ib := is_in_bounds_3d(cell_i, cell_j, cell_k)
		index := get_index_3d(cell_i, cell_j, cell_k)
		if ib && b[index] {
			count += 1
		}
	}
	return count
}

func new_generation_3d(b Board3D) Board3D {
	new_b := Board3D{}

	for i := 0; i < BOARD_DIM; i++ {
		for j := 0; j < BOARD_DIM; j++ {
			for k := 0; k < BOARD_DIM; k++ {
				index := get_index_3d(i, j, k)

				cell_alive := b[index]
				count := count_neighbours_3d(i, j, k, &b)
				/*
					Rules used: 4555
					1. A living cell remains alive only when surrounded by 4 or 5 living neighbors.
					2. A dead cell comes to life when it has exactly 5 living neighbors.
				*/
				if cell_alive {
					if count < 4 {
						new_b[index] = false
					} else if count == 4 || count == 5 {
						new_b[index] = true
					} else {
						new_b[index] = false
					}
				} else if count == 5 {
					new_b[index] = true
				}
			}
		}
	}

	return new_b
}

func chance(n int) bool {
	return rand.Intn(100) < n
}

func main_3d() {
	// initializations
	camera := rl.Camera3D{}
	offset := float32(BOARD_DIM * 0.5)
	camera.Position = rl.Vector3{X: BOARD_DIM + offset, Y: BOARD_DIM + offset, Z: BOARD_DIM + offset}
	target_center := rl.Vector3{X: BOARD_DIM / 2, Y: BOARD_DIM / 2, Z: BOARD_DIM / 2}
	camera.Target = target_center
	camera.Up = rl.Vector3{X: 0, Y: 1, Z: 0}
	camera.Fovy = 45
	camera.Projection = rl.CameraPerspective

	board := Board3D{}
	for i := 0; i < BOARD_DIM; i++ {
		for j := 0; j < BOARD_DIM; j++ {
			for k := 0; k < BOARD_DIM; k++ {
				if chance(30) {
					index := get_index_3d(i, j, k)
					board[index] = true
				}
			}
		}
	}
	reload := TIMESTEP_IN_SECS

	// raylib init window
	rl.InitWindow(SCREEN_W, SCREEN_H, "Game of Life 3D")
	defer rl.CloseWindow()
	rl.SetTargetFPS(0)
	rl.DisableCursor()

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera, rl.CameraFree)
		if rl.IsKeyPressed('Z') {
			camera.Target = target_center
		}

		// simulate
		reload -= rl.GetFrameTime()
		if reload <= 0 {
			reload = TIMESTEP_IN_SECS + reload
			board = new_generation_3d(board)
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.SkyBlue)
		rl.BeginMode3D(camera)
		{
			for i := 0; i < BOARD_DIM; i++ {
				for j := 0; j < BOARD_DIM; j++ {
					for k := 0; k < BOARD_DIM; k++ {
						index := get_index_3d(i, j, k)
						if board[index] {
							hue := float32(index) / 1000.0 * 360.0 // 0 to 360
							color := rl.ColorFromHSV(hue, 0.8, 0.9)

							pos := rl.NewVector3(float32(i), float32(j), float32(k))

							// pos := rl.Vector3{X: i, Y: j, Z: k}
							rl.DrawCube(pos, 0.75, 0.75, 0.75, color)
						}
					}
				}
			}

		}
		rl.EndMode3D()

		top := int32(20)
		left := int32(20)
		font_size := int32(20)
		rl.DrawText("Free camera default controls:", left, top, font_size, rl.Black)
		rl.DrawText("- Move with W, A, S, D", left+20, top+20, font_size, rl.DarkGray)
		rl.DrawText("- CTRL to move down", left+20, top+40, font_size, rl.DarkGray)
		rl.DrawText("- Mouse Wheel to Zoom in-out", left+20, top+60, font_size, rl.DarkGray)
		rl.DrawText("- Z to target the center", left+20, top+80, font_size, rl.DarkGray)
		rl.DrawFPS(SCREEN_W-100, 20)
		rl.EndDrawing()
	}
}
