package main

import (
	"fmt"
	"simple-belt-game/side"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var model rl.Model

func main() {
	rl.InitWindow(2000, 1200, "hello game")
	defer rl.CloseWindow()

	rl.SetTargetFPS(144)
	now := time.Now()
	logEvery := 1000 * time.Millisecond
	lastLog := now

	rl.SetTraceLogLevel(rl.LogAll)
	bg := rl.LoadTexture("resources/background/cyberpunk_street_background.png")
	model = rl.LoadModel("resources/player/robot.glb")
	defer rl.UnloadModel(model)
	anim := rl.LoadModelAnimations("resources/player/robot.glb")
	defer rl.UnloadModelAnimations(anim)

	animIdx := 0
	frame := int32(0)
	dragStart := rl.Vector2{}
	side.InitPlayerSoldiers(model)
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

		clicked := rl.IsMouseButtonDown(rl.MouseButtonLeft)

		for _, p := range side.PlayerSoldiers {
			p.Move(dt)
		}

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
		for _, p := range side.PlayerSoldiers {
			p.Draw3D()
		}
		rl.EndMode3D()
		for _, p := range side.PlayerSoldiers {
			p.Draw2D(camera3d)
		}
		if !clicked {
			dragStart = mouseLocation
		}
		if clicked {
			rl.DrawRectangleLines(
				int32(dragStart.X),
				int32(dragStart.Y),
				int32(mouseLocation.X-dragStart.X),
				int32(mouseLocation.Y-dragStart.Y),
				rl.Green,
			)
		}
		rl.EndDrawing()
	}
}
