package force

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"physx/src/physics"
	"physx/src/physics/vector"
)

func GenerateDragForce(particle *physics.Particle, k float64) *vector.Vec2 {
	dragForce := vector.NewVec2(0, 0)

	magnitudeSquared := particle.Velocity.MagnitudeSquared()
	if magnitudeSquared > 0 {
		dragDirection := particle.Velocity.UnitVector()
		dragDirection.Scale(-1.0)

		dragMagnitude := k * magnitudeSquared
		dragDirection.Scale(dragMagnitude)

		dragForce = dragDirection
	}

	return dragForce
}

func GenerateFrictionForce(particle *physics.Particle, k float64) *vector.Vec2 {
	frictionForce := particle.Velocity.UnitVector()
	frictionForce.Scale(-1.0)

	frictionMagnitude := k

	frictionForce.Scale(frictionMagnitude)

	return frictionForce
}

func GenerateGravitationalForce(a *physics.Particle, b *physics.Particle, G, minDistance, maxDistance float64) *vector.Vec2 {
	d := vector.Sub(b.Position, a.Position)

	distanceSquared := d.MagnitudeSquared()
	distanceSquared = float64(rl.Clamp(float32(distanceSquared), float32(minDistance), float32(maxDistance)))

	attractionForce := d.UnitVector()
	attractionMagnitude := G * (a.Mass * b.Mass) / distanceSquared

	attractionForce.Scale(attractionMagnitude)

	return attractionForce
}
