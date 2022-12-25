package trrs

import (
	"github.com/deadsy/sdfx/sdf"
	v3 "github.com/deadsy/sdfx/vec/v3"
)

var (
	y            = 55.8
	z            = 7.6
	holeDiameter = 7.0 // real: 6.0
)

func CreateHole() (sdf.SDF3, error) {
	circle, err := sdf.Circle2D(holeDiameter / 2)
	if err != nil {
		return nil, err
	}
	return sdf.Transform3D(
		sdf.Transform3D(
			sdf.Extrude3D(
				circle,
				10,
			),
			sdf.RotateY(sdf.Pi/2),
		),
		sdf.Translate3d(v3.Vec{X: 0, Y: y, Z: z}),
	), nil
}

func Y() float64 {
	return y
}
