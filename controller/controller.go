package controller

import (
	"github.com/deadsy/sdfx/sdf"
	v2 "github.com/deadsy/sdfx/vec/v2"
	v3 "github.com/deadsy/sdfx/vec/v3"
	"github.com/minodisk/crfrm/trrs"
)

var (
	jackBuffer    = 0.3
	jackSizeX     = 13.0 + jackBuffer
	jackSizeY     = 6.0 + jackBuffer
	wallThickness = 0.8
	// jackSizeZ = 6.0
)

func Create(
	shape sdf.SDF2,
	thickness, wallHeight, topZ float64,
) sdf.SDF3 {
	top := sdf.Transform3D(
		sdf.Extrude3D(
			shape,
			thickness,
		),
		sdf.Translate3d(v3.Vec{X: 0, Y: 0, Z: wallHeight - thickness/2}),
	)
	wall := sdf.Transform3D(
		sdf.Extrude3D(
			sdf.Difference2D(
				shape,
				sdf.Transform2D(
					shape,
					sdf.Translate2d(v2.Vec{X: -wallThickness, Y: -wallThickness}),
				),
			),
			wallHeight,
		),
		sdf.Translate3d(v3.Vec{X: 0, Y: 0, Z: wallHeight / 2}),
	)
	// space := sdf.Transform3D(
	// 	sdf.Transform3D(
	// 		sdf.ScaleExtrude3D(shape, thickness, v2.Vec{X: 0.9, Y: 0.9}),
	// 		sdf.MirrorXY(),
	// 	),
	// 	sdf.Translate3d(v3.Vec{X: 0, Y: 0, Z: topZ}),
	// )

	return sdf.Union3D(top, wall)
}

func CreateJack(
	thickness, wallThickness, wallHeight float64,
) (sdf.SDF3, error) {
	box, err := sdf.Box3D(v3.Vec{X: jackSizeX, Y: jackSizeY, Z: wallHeight}, 0)
	if err != nil {
		return nil, err
	}
	return sdf.Transform3D(
		box,
		sdf.Translate3d(v3.Vec{X: jackSizeX / 2, Y: trrs.Y(), Z: wallHeight/2 - thickness}),
	), nil
}
