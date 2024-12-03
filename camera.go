package main

type Camera struct {
	position Point
	focal    float64
	angle    float64
	scale    float64
}

func newCamera() Camera {
	return Camera{
		position: Point{0.0, 0.0, 0.0},
		focal:    0.5,
		angle:    0.0,
		scale:    100.0,
	}
}

func (c *Camera) move(x, y, z float64) {
	c.position[0] += x
	c.position[1] += y
	c.position[2] += z
}
