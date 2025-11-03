package main

import (
	"fmt"
	"simple-belt-game/movement"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(2000, 1200, "hello game")
	defer rl.CloseWindow()

	rl.SetTargetFPS(144)
	cubePos := rl.NewVector3(0.0, 0.0, 0.0)
	now := time.Now()
	logEvery := 1000 * time.Millisecond
	lastLog := now

	rl.SetTraceLogLevel(rl.LogAll)
	bg := rl.LoadTexture("resources/background/cyberpunk_street_background.png")
	model := rl.LoadModel("resources/player/my_robot_v3.glb")
	defer rl.UnloadModel(model)
	anim := rl.LoadModelAnimations("resources/player/my_robot_v3.glb")
	defer rl.UnloadModelAnimations(anim)

	animIdx := 0
	frame := int32(0)
	prevDirection := movement.Right
	for !rl.WindowShouldClose() {
		now = time.Now()
		dt := rl.GetFrameTime()
		mouseLocation := rl.GetMousePosition()

		if len(anim) > 0 && rl.IsModelAnimationValid(model, anim[animIdx]) {
			rl.UpdateModelAnimation(model, anim[animIdx], frame)
			frame++
			if frame >= anim[animIdx].FrameCount {
				frame = 0
			}
		}

		// log
		if time.Since(lastLog) >= logEvery {
			fmt.Println("mouseLocation:", mouseLocation)
			fmt.Println("dt:", dt)
			lastLog = time.Now()
		}

		var direction, move = movement.GetViewDirection(
			rl.IsKeyDown(rl.KeyLeft),
			rl.IsKeyDown(rl.KeyUp),
			rl.IsKeyDown(rl.KeyRight),
			rl.IsKeyDown(rl.KeyDown),
			prevDirection,
		)
		prevDirection = direction
		if move {
			cubePos = movement.GetNextLocation(direction, cubePos, 20, dt)
		}
		movement.Punch(rl.IsKeyDown(rl.KeyQ))
		isPunch := rl.IsKeyDown(rl.KeyQ)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.DrawTextureEx(bg, rl.Vector2{X: 0, Y: 0}, 0.0, 4.0, rl.RayWhite)

		camera3d := rl.Camera3D{
			Position:   rl.NewVector3(0.0, 10.0, 0.0),
			Target:     rl.NewVector3(0.0, 0.0, 0.0),
			Up:         rl.NewVector3(0.0, 0.0, -1.0),
			Fovy:       30.0,
			Projection: rl.CameraOrthographic,
		}
		rl.BeginMode3D(camera3d)

		// rl.DrawGrid(1000, 1.0)
		rl.PushMatrix()
		rl.Translatef(cubePos.X, cubePos.Y, cubePos.Z)
		// player cube
		rl.DrawCubeWires(rl.Vector3{X: 0, Y: 0, Z: 0}, 2.0, 2.0, 2.0, rl.Green)

		if isPunch {
			// TODO: attack range
			var attackLen float32 = 2.0
			attackVector3 := movement.FrontAttackCube(direction, attackLen)
			rl.DrawCubeWires(attackVector3, attackLen, attackLen, attackLen, rl.Red)
		}

		movement.RotateByDirection(direction)
		rl.DrawModel(model, rl.NewVector3(0, -1, 0), 0.45, rl.White)
		rl.PopMatrix()

		rl.EndMode3D()
		rl.EndDrawing()
	}
}
