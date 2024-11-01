package src

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"physx/config/colors"
	"physx/config/properties"
	"physx/src/physics"
)

type Application struct {
	particle *physics.Particle
}

func NewApplication() Application {
	return Application{
		particle: physics.NewParticle(50, 100, 1.0),
	}
}

func (app *Application) Setup() {
	rl.InitWindow(properties.WindowWidth, properties.WindowHeight, "2D Physics engine")
	rl.SetTargetFPS(144)
}

func (app *Application) Input() {

}

func (app *Application) Update() {
	dt := float64(rl.GetFrameTime())
	app.particle.Velocity = physics.NewVec2(100.0*dt, 0.0*dt)
	app.particle.Position.Add(app.particle.Velocity)
}

func (app *Application) Render() {
	rl.BeginDrawing()
	rl.ClearBackground(colors.RichGray)

	rl.DrawCircle(int32(app.particle.Position.X), int32(app.particle.Position.Y), 5, colors.SunflowerYellow)

	rl.DrawFPS(10, 10)
	rl.EndDrawing()
}

func (app *Application) Destroy() {
	app.particle = nil
	rl.CloseWindow()
}
