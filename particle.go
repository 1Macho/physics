package physics

import "github.com/1Macho/geometry"

type Particle struct {
  Position geometry.Point
  Velocity geometry.Point
  Acceleration geometry.Point
  Mass float64
}

func (p *Particle) Tick () {
  //println(p.Position.X)
  p.Velocity.X += p.Acceleration.X
  p.Velocity.Y += p.Acceleration.Y
  p.Acceleration.X = 0
  p.Acceleration.Y = 0
  p.Position.X += p.Velocity.X
  p.Position.Y += p.Velocity.Y
}

func (p *Particle) ApplyForce (force geometry.Point) {
  newAccX := force.X * p.Mass
  newAccY := force.Y * p.Mass
  p.Acceleration.X += newAccX
  p.Acceleration.Y += newAccY
}
