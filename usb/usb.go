package usb

import (
	"github.com/deadsy/sdfx/sdf"
	v2 "github.com/deadsy/sdfx/vec/v2"
	v3 "github.com/deadsy/sdfx/vec/v3"
)

var (
	usbX      = 9.2
	usbZ      = 7.2
	usbWidth  = 8.4 // real: 7.4
	usbHeight = 3.4 // real: 2.4
)

func CreateHole() sdf.SDF3 {
	return sdf.Transform3D(
		sdf.Transform3D(
			sdf.Extrude3D(
				sdf.Box2D(v2.Vec{X: usbWidth, Y: usbHeight}, 0),
				100,
			),
			sdf.Rotate3d(v3.Vec{X: 1, Y: 0, Z: 0}, sdf.Pi/2),
		),
		sdf.Translate3d(v3.Vec{X: usbX, Y: 0, Z: usbZ}),
	)
}
