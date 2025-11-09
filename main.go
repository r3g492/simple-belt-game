package main

import (
	"fmt"
	"math"
	"simple-belt-game/side"
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

		/*clicked := rl.IsMouseButtonDown(rl.MouseButtonLeft)
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
		}*/

		// mouse down edge -> begin drag
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			dragging = true
			dragStart = rl.GetMousePosition()
		}

		// while held -> draw preview
		if dragging && rl.IsMouseButtonDown(rl.MouseButtonLeft) {
			cur := rl.GetMousePosition()
			r := rectFromPoints(dragStart, cur)
			rl.DrawRectangleLines(int32(r.X), int32(r.Y), int32(r.Width), int32(r.Height), rl.Green)
		}

		// mouse up edge -> finalize
		if dragging && rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
			dragging = false
			cur := rl.GetMousePosition() // capture "cur" only when released
			selectionRect = rectFromPoints(dragStart, cur)

			// clear previous selection (or skip if you want additive with Shift)
			for i := range side.PlayerSoldiers {
				side.PlayerSoldiers[i].Selected = false
			}

			// select on intersection
			for i := range side.PlayerSoldiers {
				r := side.PlayerSoldiers[i].Get2DControlRec(camera3d) // rl.Rectangle
				if rl.CheckCollisionRecs(selectionRect, r) {
					side.PlayerSoldiers[i].Selected = true
				}
			}
		}

		for _, p := range side.PlayerSoldiers {
			p.Draw2D(camera3d)
		}

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
