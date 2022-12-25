package bottom

import (
	"github.com/deadsy/sdfx/sdf"
	v3 "github.com/deadsy/sdfx/vec/v3"
)

func Create(
	bottomShape sdf.SDF2,
	screwShape sdf.SDF2,
	thickness float64,
	skirtHeight float64,
) sdf.SDF3 {
	return sdf.Transform3D(
		sdf.Extrude3D(
			sdf.Difference2D(
				bottomShape,
				screwShape,
			),
			thickness,
		),
		sdf.Translate3d(v3.Vec{X: 0, Y: 0, Z: thickness/2 - skirtHeight}),
	)
}
