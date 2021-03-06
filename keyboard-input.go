package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type keyboardInput struct {
	speed  float32
	parent *entity
}

func createKeyboardInput(parent *entity, speed float32) *keyboardInput {
	k := keyboardInput{}

	k.parent = parent
	k.speed = speed

	return &k
}

// TODO: Change movement to be relative to rotation
func (k *keyboardInput) update() error {
	if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		k.parent.position.X -= k.speed * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
		k.parent.position.X += k.speed * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		k.parent.position.Y -= k.speed * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
		k.parent.position.Y += k.speed * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeyQ) {
		k.parent.rotation -= k.speed * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeyE) {
		k.parent.rotation += k.speed * rl.GetFrameTime()
	}

	return nil
}

func (k *keyboardInput) draw() error {
	return nil
}
