package main

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	playerSpeed        = 80
	playerShotCooldown = time.Millisecond * 250
)

func createPlayer() *entity {
	player := &entity{}

	player.position = rl.Vector2{X: screenWidth / 2.0, Y: screenHeight / 2.0}
	player.scale = 1.0
	player.active = true

	r := createRenderer(player, "sprites/Spaceship#02.png")
	player.addComponent(r)

	ki := createKeyboardInput(player, playerSpeed)
	player.addComponent(ki)

	sw := createScreenWrap(player)
	player.addComponent(sw)

	return player
}
