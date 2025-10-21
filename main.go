package main

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(2000, 1200, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	cubePos := rl.NewVector3(0.0, 0.0, 0.0)

	now := time.Now()

	logEvery := 1000 * time.Millisecond
	lastLog := now

	rl.SetTraceLogLevel(rl.LogAll)

	model := rl.LoadModel("resources/robot.glb")
	defer rl.UnloadModel(model)
	anim := rl.LoadModelAnimations("resources/robot.glb")
	defer rl.UnloadModelAnimations(anim)

	animIdx := 0
	frame := int32(0)

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

		isUp := false
		_ = false
		isLeft := false
		_ = false
		if rl.IsKeyDown(rl.KeyW) {
			isUp = true
		}
		if rl.IsKeyDown(rl.KeyA) {
			isLeft = true
		}
		if rl.IsKeyDown(rl.KeyS) && !isUp {
			// isDown = true
		}
		if rl.IsKeyDown(rl.KeyD) && !isLeft {
			// isRight = true
		}

		// 12369874

		if rl.IsKeyDown(rl.KeyW) {
			cubePos.Z -= 0.3
		}
		if rl.IsKeyDown(rl.KeyA) {
			cubePos.X -= 0.3
		}
		if rl.IsKeyDown(rl.KeyS) {
			cubePos.Z += 0.3
		}
		if rl.IsKeyDown(rl.KeyD) {
			cubePos.X += 0.3
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		camera3d := rl.Camera3D{
			Position:   rl.NewVector3(0.0, 10.0, 0.0),
			Target:     rl.NewVector3(0.0, 0.0, 0.0),
			Up:         rl.NewVector3(0.0, 0.0, -1.0),
			Fovy:       30.0,
			Projection: rl.CameraOrthographic,
		}
		rl.BeginMode3D(camera3d)
		rl.DrawGrid(1000, 1.0)

		rl.PushMatrix()
		rl.Translatef(cubePos.X, cubePos.Y, cubePos.Z)
		rl.Rotatef(270, 1, 0, 0)
		rl.Rotatef(90, 0, 1, 0)
		rl.DrawCubeWires(rl.Vector3{}, 2.0, 2.0, 2.0, rl.Red)
		rl.DrawModel(model, rl.NewVector3(0, -1, 0), 0.7, rl.White)
		rl.PopMatrix()

		rl.EndMode3D()
		rl.EndDrawing()
	}
}
