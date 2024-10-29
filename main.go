package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"physx/math/vars"
	"physx/math/vector"
)

var position = vector.New(10, 20)
var velocity = vector.New(2, 1)

func main() {
	rl.InitWindow(vars.WindowWidth, vars.WindowHeight, "2D Physics engine")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		render()
		update()

		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func render() {
	position.Draw(vars.YELLOW)
}

func update() {
	position.Add(velocity)
}
