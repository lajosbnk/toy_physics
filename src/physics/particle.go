package physics

import (
	"physx/src/physics/vector"
)

type Particle struct {
	Position     *vector.Vec2
	Velocity     *vector.Vec2
	Acceleration *vector.Vec2
	SumForces    *vector.Vec2
	Mass         float64
	MassInv      float64
	Radius       float64
}

func NewParticle(x, y, mass float64) *Particle {
	p := &Particle{
		Position:     vector.NewVec2(x, y),
		Velocity:     vector.NewVec2(0.0, 0.0),
		Acceleration: vector.NewVec2(0.0, 0.0),
		SumForces:    vector.NewVec2(0.0, 0.0),
		Mass:         mass,
	}

	if mass > 0.0 {
		p.MassInv = 1.0 / mass
	} else {
		p.MassInv = 0.0
	}

	return p
}

func (p *Particle) Integrate(dt float64) {
	//p.Acceleration = vector.Mul(p.SumForces, p.MassInv)
	p.Acceleration.X = p.SumForces.X * p.MassInv
	p.Acceleration.Y = p.SumForces.Y * p.MassInv

	//dtAcc := vector.NewVec2(p.Acceleration.X*dt, p.Acceleration.Y*dt)
	//p.Velocity.Add(dtAcc)
	p.Velocity.X += p.Acceleration.X * dt
	p.Velocity.Y += p.Acceleration.Y * dt

	//dtVel := vector.NewVec2(p.Velocity.X*dt, p.Velocity.Y*dt)
	//p.Position.Add(dtVel)
	p.Position.X += p.Velocity.X * dt
	p.Position.Y += p.Velocity.Y * dt

	p.ClearForces()
}

func (p *Particle) AddForce(force *vector.Vec2) {
	p.SumForces.Add(force)
}

func (p *Particle) ClearForces() {
	p.SumForces.X = 0.0
	p.SumForces.Y = 0.0
}
