package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	cubePos := rl.NewVector3(0.0, 0.0, 0.0)

	for !rl.WindowShouldClose() {

		if rl.IsKeyDown(rl.KeyUp) {
			cubePos.Z -= 0.3
		}
		if rl.IsKeyDown(rl.KeyLeft) {
			cubePos.X -= 0.3
		}
		if rl.IsKeyDown(rl.KeyDown) {
			cubePos.Z += 0.3
		}
		if rl.IsKeyDown(rl.KeyRight) {
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
		rl.Rotatef(45, 0, 0, 1)
		rl.DrawCubeWires(rl.Vector3{}, 2.0, 2.0, 2.0, rl.Red)
		rl.PopMatrix()

		rl.EndMode3D()
		rl.EndDrawing()
	}
}
