package wall

import (
	"github.com/deadsy/sdfx/sdf"
	v3 "github.com/deadsy/sdfx/vec/v3"
)

func Create(
	bottom sdf.SDF2,
	wallThickness, wallHeight, skirtHeight float64,
) sdf.SDF3 {
	return sdf.Transform3D(
		sdf.Extrude3D(
			sdf.Difference2D(
				sdf.Offset2D(bottom, wallThickness),
				bottom,
			),
			wallHeight+skirtHeight,
		),
		sdf.Translate3d(v3.Vec{X: 0, Y: 0, Z: (wallHeight+skirtHeight)/2 - skirtHeight}),
	)
}
