package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"physx/src"
)

func main() {
	app := src.NewApplication()

	app.Setup()

	for !rl.WindowShouldClose() {
		app.Input()
		app.Update()
		app.Render()
	}

	app.Destroy()
}
