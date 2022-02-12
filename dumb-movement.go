package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type dumbMovement struct {
	speed     float32
	rotation  float32
	direction rl.Vector2
	parent    *entity
}

func createDumbMovement(parent *entity, speed float32, rotation float32, direction rl.Vector2) *dumbMovement {
	dm := dumbMovement{}
	dm.speed = speed
	dm.rotation = rotation
	dm.direction = direction
	dm.parent = parent

	return &dm

}

func (dm *dumbMovement) update() error {
	dm.parent.position.X += dm.direction.X * dm.speed * rl.GetFrameTime()
	dm.parent.position.Y += dm.direction.Y * dm.speed * rl.GetFrameTime()
	dm.parent.rotation += dm.rotation * dm.speed * rl.GetFrameTime()

	return nil

}

func (dm *dumbMovement) draw() error {
	return nil
}
