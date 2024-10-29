package vector

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

type Vec2 struct {
	X, Y float64
}

func New(x, y float64) Vec2 {
	return Vec2{x, y}
}

func (v *Vec2) Add(other Vec2) {
	v.X += other.X
	v.Y += other.Y
}

func (v *Vec2) Dot(other Vec2) float64 {
	return v.X*other.X + v.Y*other.Y
}

func (v *Vec2) Draw(color rl.Color) {
	rl.DrawCircle(int32(v.X), int32(v.Y), 10, color)
}

func (v *Vec2) Mag() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vec2) Normalize() {
	mag := v.Mag()
	v.X /= mag
	v.Y /= mag
}

func (v *Vec2) Perpendicular() Vec2 {
	return Vec2{X: v.Y, Y: -v.X}
}

func (v *Vec2) Scale(scale float64) {
	v.X = scale * v.X
	v.Y = scale * v.Y
}

func (v *Vec2) Sub(other Vec2) {
	v.X -= other.X
	v.Y -= other.Y
}

/////////////////////////////////////////////////
//// Static Methods
/////////////////////////////////////////////////

func Add(a, b Vec2) Vec2 {
	return Vec2{a.X + b.X, a.Y + b.Y}
}

func Dot(a, b Vec2) float64 {
	return a.X*b.X + a.Y*b.Y
}

func Sub(a, b Vec2) Vec2 {
	return Vec2{a.X - b.X, a.Y - b.Y}
}
