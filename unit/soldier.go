package unit

import (
	"simple-belt-game/movement"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Soldier struct {
	Direction       movement.Direction
	Position        rl.Vector3
	Model           rl.Model
	Selected        bool
	Speed           float32
	TargetPosition  rl.Vector3
	Reached         bool
	LastReachedTime time.Time
	Type            SoldierType
}

type SoldierType int

const (
	Agent SoldierType = iota
	Shield
)

func (s *Soldier) Draw3D() {
	rl.PushMatrix()
	rl.Translatef(s.Position.X, s.Position.Y, s.Position.Z)
	if s.Selected {
		rl.DrawCubeWires(rl.Vector3{X: 0, Y: 0, Z: 0}, 2.0, 2.0, 2.0, rl.Green)
	}
	movement.RotateByDirection(s.Direction)
	rl.DrawModel(s.Model, rl.NewVector3(0, -1, 0), 0.45, rl.White)
	rl.PopMatrix()
}

func (s *Soldier) Draw2D(cam rl.Camera3D) {
	r := ScreenRectAt(s.Position, cam, 80, 80)
	rl.DrawRectangleLinesEx(r, 2, rl.Yellow)
}

func ScreenRectAt(pos rl.Vector3, cam rl.Camera3D, w, h float32) rl.Rectangle {
	p := rl.GetWorldToScreen(pos, cam)
	return rl.Rectangle{X: p.X - w/2, Y: p.Y - h/2, Width: w, Height: h}
}

func (s *Soldier) Move(
	dt float32,
) rl.Vector3 {

	speedDelta := s.Speed * dt

	if s.Direction == movement.Left {
		return rl.Vector3{X: s.Position.X - speedDelta, Y: s.Position.Y, Z: s.Position.Z}
	}

	if s.Direction == movement.LeftUp {
		return rl.Vector3{X: s.Position.X - speedDelta/2, Y: s.Position.Y, Z: s.Position.Z - speedDelta/2}
	}

	if s.Direction == movement.Up {
		return rl.Vector3{X: s.Position.X, Y: s.Position.Y, Z: s.Position.Z - speedDelta}
	}

	if s.Direction == movement.UpRight {
		return rl.Vector3{X: s.Position.X + speedDelta/2, Y: s.Position.Y, Z: s.Position.Z - speedDelta/2}
	}

	if s.Direction == movement.Right {
		return rl.Vector3{X: s.Position.X + speedDelta, Y: s.Position.Y, Z: s.Position.Z}
	}

	if s.Direction == movement.RightDown {
		return rl.Vector3{X: s.Position.X + speedDelta/2, Y: s.Position.Y, Z: s.Position.Z + speedDelta/2}
	}

	if s.Direction == movement.Down {
		return rl.Vector3{X: s.Position.X, Y: s.Position.Y, Z: s.Position.Z + speedDelta}
	}

	if s.Direction == movement.DownLeft {
		return rl.Vector3{X: s.Position.X - speedDelta/2, Y: s.Position.Y, Z: s.Position.Z + speedDelta/2}
	}

	return rl.Vector3{X: s.Position.X, Y: s.Position.Y, Z: s.Position.Z}
}
