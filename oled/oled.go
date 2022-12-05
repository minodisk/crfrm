package oled

import (
	"github.com/deadsy/sdfx/sdf"
	v2 "github.com/deadsy/sdfx/vec/v2"
	v3 "github.com/deadsy/sdfx/vec/v3"
)

var (
	oledWidth    = 20.0
	oledHeight   = 2.0
	usbPadding   = 2.0
	usbZ         = 7.5
	usbWidth     = 7.6
	usbHeight    = 2.4
	usbHoleScale = 1.2
)

func CreateOLED(oled sdf.SDF2, wallHeight float64) (sdf.SDF3, error) {
	return sdf.Transform3D(
		sdf.Extrude3D(sdf.Difference2D(oled, sdf.Offset2D(oled, -1.0)), oledHeight),
		sdf.Translate3d(v3.Vec{X: 0, Y: 0, Z: wallHeight - oledHeight/2}),
	), nil
}

func CreateUSBHole() sdf.SDF3 {
	return sdf.Transform3D(
		sdf.Transform3D(
			sdf.Extrude3D(
				sdf.Box2D(v2.Vec{X: usbWidth * usbHoleScale, Y: usbHeight * usbHoleScale}, 0),
				100,
			),
			sdf.Rotate3d(v3.Vec{X: 1, Y: 0, Z: 0}, sdf.Pi/2),
		),
		sdf.Translate3d(v3.Vec{X: oledWidth / 2, Y: 0, Z: usbZ}),
	)
}

var (
	trrsY         = 52.0
	trrsZ         = 6.8
	trrsRadius    = 2.5
	trrsHoleScale = 1.2
)

func CreateTRRSHole() (sdf.SDF3, error) {
	circle, err := sdf.Circle2D(trrsRadius * trrsHoleScale)
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
		sdf.Translate3d(v3.Vec{X: 0, Y: trrsY, Z: trrsZ}),
	), nil
}
