package physics

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

type Vec2 struct {
	X, Y float64
}

func NewVec2(x, y float64) *Vec2 {
	return &Vec2{x, y}
}

func (v *Vec2) Add(other *Vec2) {
	v.X += other.X
	v.Y += other.Y
}

func (v *Vec2) Cross(other *Vec2) float64 {
	return v.X*other.Y - v.Y*other.X
}

func (v *Vec2) Dot(other *Vec2) float64 {
	return v.X*other.X + v.Y*other.Y
}

func (v *Vec2) Draw(color rl.Color) {
	rl.DrawCircle(int32(v.X), int32(v.Y), 10, color)
}

func (v *Vec2) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vec2) MagnitudeSquared() float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v *Vec2) Normalize() *Vec2 {
	mag := v.Magnitude()
	if mag != 0.0 {
		v.X /= mag
		v.Y /= mag
	}
	return v
}

func (v *Vec2) Normal() Vec2 {
	result := Vec2{X: v.Y, Y: -v.X}

	return *result.Normalize()
}

func (v *Vec2) Perpendicular() Vec2 {
	return Vec2{X: v.Y, Y: -v.X}
}

func (v *Vec2) Rotate(angle float64) Vec2 {
	result := Vec2{X: 0, Y: 0}

	result.X = v.X*math.Cos(angle) - v.Y*math.Sin(angle)
	result.Y = v.X*math.Sin(angle) + v.Y*math.Cos(angle)

	return result
}

func (v *Vec2) Scale(scale float64) {
	v.X = scale * v.X
	v.Y = scale * v.Y
}

func (v *Vec2) Sub(other *Vec2) {
	v.X -= other.X
	v.Y -= other.Y
}

func (v *Vec2) UnitVector() *Vec2 {
	result := Vec2{X: 0, Y: 0}
	mag := v.Magnitude()

	if mag != 0.0 {
		result.X /= mag
		result.Y /= mag
	}

	return &result
}

/////////////////////////////////////////////////
//// Static Methods
/////////////////////////////////////////////////

func Add(a, b *Vec2) Vec2 {
	return Vec2{a.X + b.X, a.Y + b.Y}
}

func Dot(a, b *Vec2) float64 {
	return a.X*b.X + a.Y*b.Y
}

func Sub(a, b *Vec2) Vec2 {
	return Vec2{a.X - b.X, a.Y - b.Y}
}
