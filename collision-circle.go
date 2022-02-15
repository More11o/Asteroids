package main

import rl "github.com/gen2brain/raylib-go/raylib"

type collisonCircle struct {
	parent   *entity
	position rl.Vector2
	radius   float32
}

func createCollisionCircle(parent *entity, position rl.Vector2, radius float32) *collisonCircle {
	cc := collisonCircle{}
	cc.parent = parent
	cc.position = position
	cc.radius = radius

	return &cc

}

func (cc *collisonCircle) update() error {
	cc.position = cc.parent.position
	return nil
}

func (cc *collisonCircle) draw() error {
	return nil
}
