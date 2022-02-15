package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 600
	screenHeight = 600
	title        = "Asteroids"
	fps          = 60
)

var entities []*entity

func main() {
	rl.InitWindow(screenWidth, screenHeight, title)
	rl.SetTargetFPS(fps)

	player := createPlayer()
	entities = append(entities, player)

	ba := createBasicAsteroid(rl.Vector2{X: screenWidth / 3.0, Y: screenHeight / 2.0})
	entities = append(entities, ba)

	aa := createAdvanceAsteroid(rl.Vector2{X: screenWidth / 1.5, Y: screenHeight / 2.0})
	entities = append(entities, aa)

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		for _, entity := range entities {
			if entity.active {
				entity.update()
				entity.draw()
				entity.checkCollision()
			}
		}
		rl.EndDrawing()

	}
	rl.CloseWindow()
}
