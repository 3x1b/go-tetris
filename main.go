package main

import (
	"fmt"

	piecelib "tetris/pkg/draw_tile"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var frames int = 0

const (
	width  = 2304
	height = 1536
)

// Sprites

var (
	gopher  rl.Texture2D
	tileset rl.Texture2D
)

func update() {} // Nothing right now

var fs_key bool = false // Fullscreen key

func init() {
	rl.InitWindow(width, height, "Tetris")
	rl.SetExitKey(rl.KeyLeftAlt + rl.KeyF4)

	rl.SetTargetFPS(60)

	gopher = rl.LoadTexture("gopher.png")
	tileset = rl.LoadTexture("tileset.png")
}

func input() {

	if rl.IsKeyDown(rl.KeyLeftAlt) && rl.IsKeyDown(rl.KeyEnter) && !fs_key {
		rl.ToggleFullscreen()
		fs_key = true
	} else if rl.IsKeyDown(rl.KeyF11) && !fs_key {
		rl.ToggleFullscreen()
		fs_key = true
	} else if !(rl.IsKeyDown(rl.KeyLeftAlt) && rl.IsKeyDown(rl.KeyEnter)) &&
		!(rl.IsKeyDown(rl.KeyF11)) && fs_key {
		fs_key = false
	}

}

func draw_piece() {
	dimensions := piecelib.DrawTile("T")
	var scale float32 = 8.0

	for i := 0; i < len(dimensions); i++ {
		srcRect := rl.NewRectangle(float32(dimensions[i][0]*16), float32(dimensions[i][1]*16), 16, 16)
		offset := [2]int{8, 8}
		destRect := rl.NewRectangle(float32(dimensions[i][2]*16*int(scale)+offset[0]), float32(dimensions[i][3]*16*int(scale)+offset[1]), 16*scale, 16*scale)
		origin := rl.NewVector2(0, 0)
		rl.DrawTexturePro(tileset, srcRect, destRect, origin, 0, rl.NewColor(255, 0, 255, 255))
	}
}

func draw() {

	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)
	rl.DrawText(fmt.Sprintf("Time elapsed: %d seconds", int(frames/60)), 0, 0, 40, rl.Black)

	draw_piece()

	rl.EndDrawing()
}

func logic() {
	frames++
}

func main() {
	rl.ToggleFullscreen()

	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		update()
		logic()
		input()
		draw()
	}

	rl.UnloadTexture(gopher)

	rl.CloseWindow()
}
