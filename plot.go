package main

func plot(point Point, camera Camera) Pixel {
	var plotedX, plotedY int

	dx := point[0] - camera.position[0]
	dy := point[1] - camera.position[1]
	dz := point[2] - camera.position[2]

	plotedX = halfWidth + round(camera.scale*camera.focal*dx/dz)
	plotedY = halfHeight - round(camera.scale*camera.focal*dy/dz)

	return Pixel{plotedX, plotedY}
}
