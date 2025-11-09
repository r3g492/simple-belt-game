package unit

import (
	"simple-belt-game/movement"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Soldier struct {
	Direction movement.Direction
	Position  rl.Vector3
	Model     rl.Model
}

func (s *Soldier) DrawSoldier() {
	rl.PushMatrix()
	rl.Translatef(s.Position.X, s.Position.Y, s.Position.Z)
	// player cube
	rl.DrawCubeWires(rl.Vector3{X: 0, Y: 0, Z: 0}, 2.0, 2.0, 2.0, rl.Green)
	movement.RotateByDirection(s.Direction)
	rl.DrawModel(s.Model, rl.NewVector3(0, -1, 0), 0.45, rl.White)
	rl.PopMatrix()
}
