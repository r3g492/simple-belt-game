package main

import (
	"fmt"
	"math"
	"simple-belt-game/side"
	"simple-belt-game/unit"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	model         rl.Model
	dragging      bool
	dragStart     rl.Vector2
	selectionRect rl.Rectangle
)

func main() {
	rl.InitWindow(2000, 1200, "Kill Squad")
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
	rl.DisableCursor()

	animIdx := 0
	frame := int32(0)
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

		var clickBegin = rl.IsMouseButtonPressed(rl.MouseButtonLeft)
		var clickHold = dragging && rl.IsMouseButtonDown(rl.MouseButtonLeft)
		var clickRelease = dragging && rl.IsMouseButtonReleased(rl.MouseButtonLeft)
		var actionClick = !dragging && rl.IsMouseButtonPressed(rl.MouseButtonRight)

		for i := range side.PlayerSoldiers {
			side.PlayerSoldiers[i].Act(dt)
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
		// rl.DrawGrid(1000, 1.0)

		if actionClick {
			var ray = rl.GetScreenToWorldRay(mouseLocation, camera3d)
			target3D := rl.Vector3{
				X: ray.Position.X,
				Y: 0,
				Z: ray.Position.Z,
			}
			rl.DrawSphere(target3D, 0.1, rl.Red)

			for i := range side.PlayerSoldiers {
				if side.PlayerSoldiers[i].Selected == true {
					side.PlayerSoldiers[i].Status = unit.Move
					side.PlayerSoldiers[i].TargetPosition = target3D
					fmt.Println(side.PlayerSoldiers[i].Status)
					fmt.Println(side.PlayerSoldiers[i].TargetPosition)
				}
			}
		}
		rl.EndMode3D()

		if clickBegin {
			dragging = true
			dragStart = mouseLocation
		}

		if clickHold {
			cur := mouseLocation
			r := rectFromPoints(dragStart, cur)
			rl.DrawRectangleLines(int32(r.X), int32(r.Y), int32(r.Width), int32(r.Height), rl.Green)
		}

		if clickRelease {
			dragging = false
			cur := mouseLocation
			selectionRect = rectFromPoints(dragStart, cur)

			for i := range side.PlayerSoldiers {
				side.PlayerSoldiers[i].Selected = false
			}

			for i := range side.PlayerSoldiers {
				r := side.PlayerSoldiers[i].Get2DControlRec(camera3d)
				if rl.CheckCollisionRecs(selectionRect, r) {
					side.PlayerSoldiers[i].Selected = true
				}
			}
		}

		for _, p := range side.PlayerSoldiers {
			p.Draw2D(camera3d)
		}

		// draw cursor
		rl.DrawRectangle(
			int32(mouseLocation.X),
			int32(mouseLocation.Y),
			10,
			10,
			rl.Blue,
		)

		rl.EndDrawing()
	}
}

func rectFromPoints(a, b rl.Vector2) rl.Rectangle {
	x := float32(math.Min(float64(a.X), float64(b.X)))
	y := float32(math.Min(float64(a.Y), float64(b.Y)))
	w := float32(math.Abs(float64(b.X - a.X)))
	h := float32(math.Abs(float64(b.Y - a.Y)))
	return rl.Rectangle{X: x, Y: y, Width: w, Height: h}
}
