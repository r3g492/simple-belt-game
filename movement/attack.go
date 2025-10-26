package movement

import rl "github.com/gen2brain/raylib-go/raylib"

func Punch(
	punch bool,
) {

}

func FrontAttackCube(direction Direction) rl.Vector3 {
	if direction == None {
		return rl.Vector3{X: 0, Y: 0, Z: 0}
	}

	if direction == Left {
		return rl.Vector3{X: -2, Y: 0, Z: 0}
	}

	if direction == LeftUp {
		return rl.Vector3{X: -1, Y: 0, Z: -1}
	}

	if direction == Up {
		return rl.Vector3{X: 0, Y: 0, Z: -2}
	}

	if direction == UpRight {
		return rl.Vector3{X: 1, Y: 0, Z: -1}
	}

	if direction == Right {
		return rl.Vector3{X: 2, Y: 0, Z: 0}
	}

	if direction == RightDown {
		return rl.Vector3{X: 1, Y: 0, Z: 1}
	}

	if direction == Down {
		return rl.Vector3{X: 0, Y: 0, Z: 2}
	}

	if direction == DownLeft {
		return rl.Vector3{X: -1, Y: 0, Z: 1}
	}

	return rl.Vector3{X: 0, Y: 0, Z: 0}
}
