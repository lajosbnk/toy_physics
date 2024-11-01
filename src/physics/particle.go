package physics

type Particle struct {
	Position     *Vec2
	Velocity     *Vec2
	Acceleration *Vec2
	Mass         float64
}

func NewParticle(x, y float64, mass float64) *Particle {
	return &Particle{
		Position:     NewVec2(x, y),
		Velocity:     NewVec2(0.0, 0.0),
		Acceleration: NewVec2(0.0, 0.0),
		Mass:         mass,
	}
}
