package movement

import rl "github.com/gen2brain/raylib-go/raylib"

func GetNextLocation(
	direction Direction,
	curPos rl.Vector3,
	speed float32,
	dt float32,
) rl.Vector3 {

	speedDelta := speed * dt

	if direction == Left {
		return rl.Vector3{X: curPos.X - speedDelta, Y: curPos.Y, Z: curPos.Z}
	}

	if direction == LeftUp {
		return rl.Vector3{X: curPos.X - speedDelta/2, Y: curPos.Y, Z: curPos.Z - speedDelta/2}
	}

	if direction == Up {
		return rl.Vector3{X: curPos.X, Y: curPos.Y, Z: curPos.Z - speedDelta}
	}

	if direction == UpRight {
		return rl.Vector3{X: curPos.X + speedDelta/2, Y: curPos.Y, Z: curPos.Z - speedDelta/2}
	}

	if direction == Right {
		return rl.Vector3{X: curPos.X + speedDelta, Y: curPos.Y, Z: curPos.Z}
	}

	if direction == RightDown {
		return rl.Vector3{X: curPos.X + speedDelta/2, Y: curPos.Y, Z: curPos.Z + speedDelta/2}
	}

	if direction == Down {
		return rl.Vector3{X: curPos.X, Y: curPos.Y, Z: curPos.Z + speedDelta}
	}

	if direction == DownLeft {
		return rl.Vector3{X: curPos.X - speedDelta/2, Y: curPos.Y, Z: curPos.Z + speedDelta/2}
	}

	return rl.Vector3{X: curPos.X, Y: curPos.Y, Z: curPos.Z}
}
