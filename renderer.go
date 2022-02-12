package main

import (
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type renderer struct {
	parent  *entity
	texture rl.Texture2D
}

func createRenderer(parent *entity, filename string) *renderer {
	r := renderer{}

	r.texture = rl.LoadTexture(filename)
	if (r.texture == rl.Texture2D{}) {
		log.Fatalf("failed to create texture: %s \n", filename)
	}

	r.parent = parent

	return &r
}

func (r *renderer) draw() error {
	position := r.parent.position

	position.X -= float32(r.texture.Width)
	position.Y -= float32(r.texture.Height)

	rl.DrawTexturePro(
		r.texture,
		rl.Rectangle{X: 0, Y: 0, Width: float32(r.texture.Width), Height: float32(r.texture.Height)},
		rl.Rectangle{X: position.X, Y: position.Y, Width: float32(r.texture.Width), Height: float32(r.texture.Height)},
		rl.Vector2{X: float32(r.texture.Width) / 2.0, Y: float32(r.texture.Height) / 2.0},
		r.parent.rotation,
		rl.White)

	return nil
}

func (r *renderer) update() error {
	return nil
}
