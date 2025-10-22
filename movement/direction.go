package movement

import rl "github.com/gen2brain/raylib-go/raylib"

// Define a custom type for the enum
type Direction int

// Declare constants for the enum values using iota
const (
	Left Direction = iota
	LeftUp
	Up
	UpRight
	Right
	RightDown
	Down
	DownLeft
	None
)

func GetViewDirection(
	isLeftKeyDown bool,
	isUpKeyDown bool,
	isRightKeyDown bool,
	isDownKeyDown bool,
	prev Direction,
) (Direction, bool) {
	if isLeftKeyDown && isRightKeyDown {
		isLeftKeyDown = false
		isRightKeyDown = false
	}
	if isUpKeyDown && isDownKeyDown {
		isUpKeyDown = false
		isDownKeyDown = false
	}

	// combination
	if isLeftKeyDown && isUpKeyDown {
		return LeftUp, true
	}
	if isUpKeyDown && isRightKeyDown {
		return UpRight, true
	}
	if isRightKeyDown && isDownKeyDown {
		return RightDown, true
	}
	if isDownKeyDown && isLeftKeyDown {
		return DownLeft, true
	}

	// single
	if isLeftKeyDown {
		return Left, true
	}
	if isUpKeyDown {
		return Up, true
	}
	if isRightKeyDown {
		return Right, true
	}
	if isDownKeyDown {
		return Down, true
	}

	return prev, false
}

func RotateByDirection(direction Direction) {
	if direction == None {
		return
	}

	if direction == Left {
		rl.Rotatef(-90, 0, 1, 0)
	}

	if direction == LeftUp {
		rl.Rotatef(-135, 0, 1, 0)
	}

	if direction == Up {
		rl.Rotatef(180, 0, 1, 0)
	}

	if direction == UpRight {
		rl.Rotatef(135, 0, 1, 0)
	}

	if direction == Right {
		rl.Rotatef(90, 0, 1, 0)
	}

	if direction == RightDown {
		rl.Rotatef(45, 0, 1, 0)
	}

	if direction == Down {
		rl.Rotatef(0, 0, 1, 0)
	}

	if direction == DownLeft {
		rl.Rotatef(-45, 0, 1, 0)
	}
}
