package top

import (
	"github.com/deadsy/sdfx/sdf"
	v3 "github.com/deadsy/sdfx/vec/v3"
)

func Create(
	shape, screw, key sdf.SDF2,
	thickness, z float64,
) sdf.SDF3 {
	return sdf.Transform3D(
		sdf.Extrude3D(
			sdf.Difference2D(
				shape,
				sdf.Union2D(key, screw),
			),
			thickness,
		),
		sdf.Translate3d(v3.Vec{X: 0, Y: 0, Z: z}),
	)
}
