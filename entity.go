package main

import (
	"fmt"
	"log"
	"reflect"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type entity struct {
	// position should be center of object.
	position rl.Vector2
	rotation float32
	// Scale default should be 1.0f
	scale      float32
	active     bool
	components []component
}

type component interface {
	update() error
	draw() error
}

func (e *entity) addComponent(new component) error {
	for _, existing := range e.components {
		if reflect.TypeOf(new) == reflect.TypeOf(existing) {
			return fmt.Errorf("already have component of type: %v", reflect.TypeOf(new))
		}
	}
	e.components = append(e.components, new)
	return nil
}

// entity.getComponent(componentName{}).(*componentName)
func (e *entity) getComponent(c component) component {
	for _, component := range e.components {
		if reflect.TypeOf(component) == reflect.TypeOf(c) {
			return component
		}
	}
	return nil
}

// Itterate though attached componets and call thier draw function.
func (e *entity) draw() error {
	for _, component := range e.components {
		err := component.draw()
		if err != nil {
			log.Fatalf("failed to draw: %s \n", err)
		}
	}
	return nil
}

// Itterate though attached componets and call thier update function.
func (e *entity) update() error {
	for _, component := range e.components {
		err := component.update()
		if err != nil {
			log.Fatalf("failed to update entity: %s \n", err)
		}
	}
	return nil
}

func (e *entity) checkCollision() error {

	if e.getComponent(&collisonCircle{}) == nil {
		return nil
	}

	cc := e.getComponent(&collisonCircle{}).(*collisonCircle)

	for _, otherEntity := range entities {
		if e == otherEntity {
			return nil
		}
		if otherEntity.getComponent(&collisonCircle{}) == nil {
			return nil
		}
		occ := otherEntity.getComponent(&collisonCircle{}).(*collisonCircle)
		if rl.CheckCollisionCircles(cc.position, cc.radius, occ.position, occ.radius) {
			e.active = false
			occ.parent.active = false
		}
	}
	return nil
}
