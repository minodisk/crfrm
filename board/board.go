package board

import (
	"github.com/deadsy/sdfx/sdf"
	v3 "github.com/deadsy/sdfx/vec/v3"
)

func Create(
	bottom sdf.SDF2,
	thickness, boardHeight float64,
) sdf.SDF3 {
	h := thickness + boardHeight
	return sdf.Transform3D(
		sdf.Extrude3D(bottom, h),
		sdf.Translate3d(v3.Vec{X: 0, Y: 0, Z: h / 2}),
	)
}
