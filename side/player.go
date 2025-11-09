package side

import (
	"simple-belt-game/movement"
	"simple-belt-game/unit"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	PlayerSoldiers []unit.Soldier
)

func InitPlayerSoldiers(
	model rl.Model,
) {
	PlayerSoldiers = append(PlayerSoldiers,
		unit.Soldier{
			Direction: movement.Down,
			Position: rl.Vector3{
				X: 0,
				Y: 0,
				Z: 0,
			},
			Model: model,
		})
}
