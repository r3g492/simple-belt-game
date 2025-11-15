package unit

import (
	"math"
	"simple-belt-game/movement"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Soldier struct {
	Direction       movement.Direction
	PrevPosition    rl.Vector3
	Position        rl.Vector3
	Model           rl.Model
	Selected        bool
	Speed           float32
	TargetPosition  rl.Vector3
	Reached         bool
	LastReachedTime time.Time
	Type            SoldierType
	Status          SoldierStatus
	Size            float32
}

type SoldierType int

const (
	Agent SoldierType = iota
	Shield
)

type SoldierStatus int

const (
	Idle SoldierStatus = iota
	Move
	Attack
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
	if s.Selected {
		var w float32 = 80
		var h float32 = 80
		p := rl.GetWorldToScreen(s.Position, cam)
		r := rl.Rectangle{X: p.X - w/2, Y: p.Y - h/2, Width: w, Height: h}
		rl.DrawRectangleLinesEx(r, 2, rl.Green)
	}
}

func (s *Soldier) Get2DControlRec(cam rl.Camera3D) rl.Rectangle {
	var w float32 = 80
	var h float32 = 80
	p := rl.GetWorldToScreen(s.Position, cam)
	r := rl.Rectangle{X: p.X - w/2, Y: p.Y - h/2, Width: w, Height: h}
	return r
}

func (s *Soldier) Act(
	dt float32,
	soldierList []Soldier,
	selfIdx int,
) {
	if s.Status == Move || s.Status == Attack {
		var moveDirection = rl.Vector3Subtract(s.TargetPosition, s.Position)
		var distance = rl.Vector3Length(moveDirection)

		var speedDelta = s.Speed * dt

		if distance <= speedDelta || distance < 0.01 {
			s.Position = s.TargetPosition
			s.Status = Idle
			return
		}

		var unitDirection = rl.Vector3Normalize(moveDirection)

		var moveAmount = rl.Vector3Scale(unitDirection, speedDelta)
		var nextPosition = rl.Vector3Add(s.Position, moveAmount)

		var angle = math.Atan2(float64(unitDirection.Z), float64(unitDirection.X))
		var angleDegrees = angle * (180 / math.Pi)
		if angleDegrees > -22.5 && angleDegrees <= 22.5 {
			s.Direction = movement.Right
		} else if angleDegrees > 22.5 && angleDegrees <= 67.5 {
			s.Direction = movement.RightDown
		} else if angleDegrees > 67.5 && angleDegrees <= 112.5 {
			s.Direction = movement.Down
		} else if angleDegrees > 112.5 && angleDegrees <= 157.5 {
			s.Direction = movement.DownLeft
		} else if angleDegrees > 157.5 || angleDegrees <= -157.5 {
			s.Direction = movement.Left
		} else if angleDegrees > -157.5 && angleDegrees <= -112.5 {
			s.Direction = movement.LeftUp
		} else if angleDegrees > -112.5 && angleDegrees <= -67.5 {
			s.Direction = movement.Up
		} else if angleDegrees > -67.5 && angleDegrees <= -22.5 {
			s.Direction = movement.UpRight
		}

		for i, other := range soldierList {
			if i == selfIdx {
				continue
			}

			minDistance := s.Size + other.Size

			distanceToOther := rl.Vector3Distance(nextPosition, other.Position)

			if distanceToOther < minDistance && distanceToOther > 0.001 {
				overlap := minDistance - distanceToOther
				var pushDirection = rl.Vector3Subtract(nextPosition, other.Position)
				pushDirection = rl.Vector3Normalize(pushDirection)
				var correction = rl.Vector3Scale(pushDirection, overlap)
				nextPosition = rl.Vector3Add(nextPosition, correction)
			} else if distanceToOther <= 0.001 {
				arbitraryPush := rl.NewVector3(minDistance, 0, 0)
				nextPosition = rl.Vector3Add(nextPosition, arbitraryPush)
			}
		}

		s.Position = nextPosition
	}
}
