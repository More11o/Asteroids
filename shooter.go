package main

import (
	"math"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type shooter struct {
	parent     *entity
	cooldown   time.Duration
	lastShot   time.Time
	bulletPool [15]*entity
}

func createShooter(parent *entity, cooldown time.Duration) *shooter {
	s := shooter{}
	s.parent = parent
	s.cooldown = cooldown
	s.lastShot = time.Now()

	for i := 0; i < len(s.bulletPool); i++ {
		b := createBullet()
		s.bulletPool[i] = b
		entities = append(entities, b)
	}

	return &s
}

func (s *shooter) draw() error {
	return nil
}

func (s *shooter) update() error {
	for _, bullet := range s.bulletPool {
		bullet.update()
	}

	// TODO:
	//	This whole function needs tidying up. Do not like the use of magic numbers
	//	and math quirks (negative y in createDumbMovement).

	if rl.IsKeyDown(rl.KeySpace) {
		if time.Since(s.lastShot) > s.cooldown {
			b := s.getBullet()
			b.active = true

			b.rotation = s.parent.rotation

			// 24 is the ship size.
			b.position = rl.Vector2{
				X: s.parent.position.X + float32(math.Sin(float64(b.rotation*rl.Deg2rad))*24),
				Y: s.parent.position.Y - float32(math.Cos(float64(b.rotation*rl.Deg2rad))*24),
			}

			// playerSpeed is used here in place of a bullet speed.
			x := 1.5 * float32(math.Sin(float64(b.rotation*rl.Deg2rad))*playerSpeed)
			y := 1.5 * float32(math.Cos(float64(b.rotation*rl.Deg2rad))*playerSpeed)

			//dm := createDumbMovement(b, 0.5, 0, rl.Vector2{X: x, Y: -y})
			//b.addComponent(dm)

			ndm := b.getComponent(&dumbMovement{}).(*dumbMovement)
			ndm.direction = rl.Vector2{X: x, Y: -y}

			s.lastShot = time.Now()
		}

	}
	return nil
}

func createBullet() *entity {
	b := &entity{}
	b.active = false
	b.scale = 2.0

	sr := createRenderer(b, "sprites/shot.png")
	b.addComponent(sr)

	sk := createScreenKill(b)
	b.addComponent(sk)

	dm := createDumbMovement(b, 0.5, 0, rl.Vector2{})
	b.addComponent(dm)

	cc := createCollisionCircle(b, b.position, 5)
	b.addComponent(cc)

	return b
}

func (s *shooter) getBullet() *entity {
	for _, bullet := range s.bulletPool {
		if !bullet.active {
			return bullet
		}
	}
	return nil
}
