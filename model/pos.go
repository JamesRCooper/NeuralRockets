package model

import "math"

//Pos is a container for an object's position
type Pos struct {
	X, Y, A float64
}

//Move moves the pos in the direction speicifed
func (p *Pos) Move(x float64, y float64) {
	p.X += x
	p.Y += y
}

//Rotate augments the pointing vector of the pos by the amount in PI radians
func (p *Pos) Rotate(angle float64) {
	p.A += angle
}

//Vec is a container for a pointing vector
type Vec struct {
	X, Y float64
}

//Angle returns the angle up from X of the vector
func (v Vec) Angle() float64 {
	return math.Atan2(v.Y, v.X)
}

//Mult changes the length of a vector
func (v *Vec) Mult(mag float64) {
	v.X *= mag
	v.Y *= mag
}

//Add combines this vector with a new supplied vector
func (v *Vec) Add(augmentingVec Vec) {
	v.X += augmentingVec.X
	v.Y += augmentingVec.Y
}

//Copy returns a deep copy of the vector
func (v Vec) Copy() Vec {
	return Vec{v.X, v.Y}
}
