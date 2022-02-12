package main

import rl "github.com/gen2brain/raylib-go/raylib"

type renderer struct {
	parent  *entity
	texture rl.Texture2D
}

func createRenderer(parent *entity, filename string) *renderer {
	r := renderer{}

	r.texture = rl.LoadTexture(filename)
	if (r.texture == rl.Texture2D{}) {
		return nil
	}

	r.parent = parent

	return &r
}

func (r *renderer) draw() error {
	position := r.parent.position

	position.X -= float32(r.texture.Width)
	position.Y -= float32(r.texture.Height)

	// FIXME: Should rotate around center point
	rl.DrawTextureEx(r.texture, position, r.parent.rotation, r.parent.scale, rl.White)

	return nil
}

func (r *renderer) update() error {
	return nil
}
