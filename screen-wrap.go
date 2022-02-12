package main

type screenWrap struct {
	parent *entity
}

type screenKill struct {
	parent *entity
}

func createScreenWrap(parent *entity) *screenWrap {
	sw := screenWrap{}

	sw.parent = parent

	return &sw
}

func createScreenKill(parent *entity) *screenKill {
	sk := screenKill{}

	sk.parent = parent

	return &sk
}

func (sw *screenWrap) update() error {
	if sw.parent.position.X > screenWidth {
		sw.parent.position.X -= screenWidth
	}

	if sw.parent.position.X < 0 {
		sw.parent.position.X += screenWidth
	}

	if sw.parent.position.Y > screenHeight {
		sw.parent.position.Y -= screenHeight
	}

	if sw.parent.position.Y < 0 {
		sw.parent.position.Y += screenHeight
	}
	return nil
}

func (sw *screenWrap) draw() error {
	return nil
}

func (sw *screenKill) update() error {
	if sw.parent.position.X > screenWidth {
		sw.parent.active = false
	}

	if sw.parent.position.X < 0 {
		sw.parent.active = false
	}

	if sw.parent.position.Y > screenHeight {
		sw.parent.active = false
	}

	if sw.parent.position.Y < 0 {
		sw.parent.active = false
	}
	return nil
}

func (sw *screenKill) draw() error {
	return nil
}
