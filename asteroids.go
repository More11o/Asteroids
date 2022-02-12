package main

import rl "github.com/gen2brain/raylib-go/raylib"

const asteroidSpeed = 80

func createBasicAsteroid(position rl.Vector2) *entity {
	e := &entity{}

	e.position = position
	e.scale = 1.0
	e.active = true

	r := createRenderer(e, "sprites/Asteroid#01.png")
	e.addComponent(r)

	// TODO: Create random rotation and direction values.
	dm := createDumbMovement(e, asteroidSpeed, 0.5, rl.Vector2{X: 0.5, Y: 0.2})
	e.addComponent(dm)

	sw := createScreenWrap(e)
	e.addComponent(sw)

	return e
}

// TODO: Create random rotation and direction values.
func createAdvanceAsteroid(position rl.Vector2) *entity {
	e := &entity{}

	e.position = position
	e.scale = 1.0
	e.active = true

	r := createRenderer(e, "sprites/Asteroid#02.png")
	e.addComponent(r)

	dm := createDumbMovement(e, asteroidSpeed, -0.3, rl.Vector2{X: -0.2, Y: -0.15})
	e.addComponent(dm)

	sw := createScreenWrap(e)
	e.addComponent(sw)

	return e
}
