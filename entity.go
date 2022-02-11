package main

import (
	"fmt"
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

/* func (e *entity) getComponent(c component) component {
	for _, component := range e.components {
		if reflect.TypeOf(component) == reflect.TypeOf(c) {
			return component
		}
	}
	return nil
} */

// Itterate though attached componets and call thier draw function.
func (e *entity) draw() error {
	for _, component := range e.components {
		err := component.draw()
		if err != nil {
			return err
		}
	}
	return nil
}

// Itterate though attached componets and call thier update function.
func (e *entity) update() error {
	for _, component := range e.components {
		err := component.update()
		if err != nil {
			return err
		}
	}
	return nil
}
