package main

import rl "github.com/gen2brain/raylib-go/raylib"

const asteroidSpeed = 80

func createBasicAsteroid(position rl.Vector2) *entity {
	ba := &entity{}

	ba.position = position
	ba.scale = 1.0
	ba.active = true

	r := createRenderer(ba, "sprites/Asteroid#01.png")
	ba.addComponent(r)

	// TODO: Create random rotation and direction values.
	dm := createDumbMovement(ba, asteroidSpeed, 0.5, rl.Vector2{X: 0.5, Y: 0.2})
	ba.addComponent(dm)

	return ba
}

// TODO: Create random rotation and direction values.
func createAdvanceAsteroid(position rl.Vector2) *entity {
	aa := &entity{}

	aa.position = position
	aa.scale = 1.0
	aa.active = true

	r := createRenderer(aa, "sprites/Asteroid#02.png")
	aa.addComponent(r)

	dm := createDumbMovement(aa, asteroidSpeed, -0.3, rl.Vector2{X: -0.2, Y: -0.15})
	aa.addComponent(dm)

	return aa
}
