package main

import (
	"fmt"

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

func draw() {

	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)
	rl.DrawText(fmt.Sprintf("Time elapsed: %d seconds", int(frames/60)), 0, 0, 40, rl.Black)

	rl.DrawTexture(gopher, 0, 0, rl.White)

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
