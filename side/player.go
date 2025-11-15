package side

import (
	"simple-belt-game/movement"
	"simple-belt-game/unit"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	PlayerSoldiers []unit.Soldier
)

func InitPlayerSoldiers(
	model rl.Model,
) {
	PlayerSoldiers = append(
		PlayerSoldiers,
		unit.Soldier{
			Direction: movement.Down,
			Position: rl.Vector3{
				X: 0,
				Y: 0,
				Z: 0,
			},
			PrevPosition: rl.Vector3{
				X: 0,
				Y: 0,
				Z: 0,
			},
			Model:    model,
			Selected: false,
			Speed:    5,
			TargetPosition: rl.Vector3{
				X: 0,
				Y: 0,
				Z: 0,
			},
			Reached:         false,
			LastReachedTime: time.Now(),
			Type:            unit.Agent,
			Status:          unit.Idle,
			Size:            0.8,
		},
		unit.Soldier{
			Direction: movement.Down,
			Position: rl.Vector3{
				X: 10,
				Y: 0,
				Z: 0,
			},
			PrevPosition: rl.Vector3{
				X: 0,
				Y: 0,
				Z: 0,
			},
			Model:    model,
			Selected: false,
			Speed:    5,
			TargetPosition: rl.Vector3{
				X: 0,
				Y: 0,
				Z: 0,
			},
			Reached:         false,
			LastReachedTime: time.Now(),
			Type:            unit.Agent,
			Status:          unit.Idle,
			Size:            0.8,
		},
	)
}
