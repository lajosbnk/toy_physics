package src

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"physx/config/colors"
	"physx/config/properties"
	"physx/src/physics"
	"physx/src/physics/force"
	"physx/src/physics/vector"
)

type Application struct {
	particles     []*physics.Particle
	PushForce     *vector.Vec2
	MousePosition *vector.Vec2
}

func NewApplication() Application {
	app := Application{
		particles:     make([]*physics.Particle, 0, 10),
		PushForce:     vector.NewVec2(0.0, 0.0),
		MousePosition: vector.NewVec2(0.0, 0.0),
	}
	particle1 := physics.NewParticle(200, 200, 1.0)
	particle1.Radius = 6
	app.particles = append(app.particles, particle1)

	particle2 := physics.NewParticle(500, 500, 20.0)
	particle2.Radius = 20
	app.particles = append(app.particles, particle2)

	return app
}

func (app *Application) Setup() {
	rl.InitWindow(properties.WindowWidth, properties.WindowHeight, "2D Physics engine")
	rl.SetTargetFPS(144)
}

func (app *Application) Input() {
	if rl.IsKeyPressed(rl.KeyUp) {
		app.PushForce.Y = -50 * properties.PixelsPerMeter
	} else if rl.IsKeyReleased(rl.KeyUp) {
		app.PushForce.Y = 0
	}
	if rl.IsKeyPressed(rl.KeyDown) {
		app.PushForce.Y = 50 * properties.PixelsPerMeter
	} else if rl.IsKeyReleased(rl.KeyDown) {
		app.PushForce.Y = 0
	}
	if rl.IsKeyPressed(rl.KeyLeft) {
		app.PushForce.X = -50 * properties.PixelsPerMeter
	} else if rl.IsKeyReleased(rl.KeyLeft) {
		app.PushForce.X = 0
	}
	if rl.IsKeyPressed(rl.KeyRight) {
		app.PushForce.X = 50 * properties.PixelsPerMeter
	} else if rl.IsKeyReleased(rl.KeyRight) {
		app.PushForce.X = 0
	}
}

func (app *Application) Update() {
	dt := rl.GetFrameTime()

	for _, p := range app.particles {
		p.AddForce(force.GenerateFrictionForce(p, 10.0))

		p.Integrate(float64(dt))

		if p.Position.X-p.Radius <= 0 {
			p.Position.X = p.Radius
			p.Velocity.X *= -0.9
		} else if p.Position.X+p.Radius >= properties.WindowWidth {
			p.Position.X = properties.WindowWidth - p.Radius
			p.Velocity.X *= -0.9
		}
		if p.Position.Y-p.Radius <= 0 {
			p.Position.Y = p.Radius
			p.Velocity.Y *= -0.9
		} else if p.Position.Y+p.Radius >= properties.WindowHeight {
			p.Position.Y = properties.WindowHeight - p.Radius
			p.Velocity.Y *= -0.9
		}
	}

	app.particles[1].AddForce(app.PushForce)
	attraction := force.GenerateGravitationalForce(app.particles[0], app.particles[1], 1000.0, 5, 100)
	app.particles[0].AddForce(attraction)
	attraction.Scale(-1.0)
	app.particles[1].AddForce(attraction)

}

func (app *Application) Render() {
	rl.BeginDrawing()
	rl.ClearBackground(colors.RichGray)

	rl.DrawCircle(int32(app.particles[0].Position.X), int32(app.particles[0].Position.Y), float32(app.particles[0].Radius), colors.AzureBlue)
	rl.DrawCircle(int32(app.particles[1].Position.X), int32(app.particles[1].Position.Y), float32(app.particles[1].Radius), colors.SunflowerYellow)

	rl.DrawFPS(10, 10)
	rl.EndDrawing()
}

func (app *Application) Destroy() {
	rl.CloseWindow()
}
