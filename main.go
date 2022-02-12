package main

import (
	"fmt"

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

	ba := createBasicAsteroid(rl.Vector2{X: screenWidth / 2.0, Y: screenHeight / 2.0})
	entities = append(entities, ba)

	aa := createAdvanceAsteroid(rl.Vector2{X: screenWidth / 2.0, Y: screenHeight / 2.0})
	entities = append(entities, aa)

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		for _, entity := range entities {
			if entity.active {
				err := entity.update()
				if err != nil {
					fmt.Errorf("failed to update: %v", entity)
				}
				err = entity.draw()
				if err != nil {
					fmt.Errorf("failed to draw: %v", entity)
				}
			}
		}
		rl.EndDrawing()

	}
	rl.CloseWindow()
}
