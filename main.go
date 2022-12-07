package main

import (
	"fmt"
	"log"

	"github.com/minodisk/crfrm/oled"
	"github.com/minodisk/crfrm/svg"
	"github.com/pkg/errors"

	"github.com/deadsy/sdfx/render"
	"github.com/deadsy/sdfx/sdf"
	v3 "github.com/deadsy/sdfx/vec/v3"
)

var (
	meshCells     = 300
	buffer        = 0.2
	thickness     = 1.6
	wallThickness = 1.6
	boardHeight   = 4.0
	wallHeight    = 13.8
)

func main() {
	svg, err := svg.NewSVG("./assets.svg")
	if err != nil {
		log.Fatal(errors.Wrap(err, "fail to parse SVG"))
	}

	f, err := top(svg)
	if err != nil {
		log.Fatal(errors.Wrap(err, "fail to create top"))
	}
	render.ToSTL(f, "crfrm.stl", render.NewMarchingCubesUniform(meshCells))
}

func top(svg *svg.SVG) (sdf.SDF3, error) {
	top, err := svg.Find("top").Draw()
	if err != nil {
		return nil, err
	}
	bottom, err := svg.Find("bottom").Draw()
	if err != nil {
		return nil, err
	}
	oled2D, err := svg.Find("oled").Draw()
	if err != nil {
		return nil, err
	}
	screw, err := svg.Find("screw").Draw()
	if err != nil {
		return nil, err
	}
	key, err := svg.Find("key").Draw()
	if err != nil {
		return nil, err
	}

	wall := createWall(top, bottom, oled2D)
	keyPlate := createKeyPlate(top, screw, key)
	boardSpace := createBoardSpace(bottom)
	o, err := oled.CreateOLED(oled2D, wallHeight)
	if err != nil {
		return nil, err
	}
	usb := oled.CreateUSBHole()
	trrs, err := oled.CreateTRRSHole()
	if err != nil {
		return nil, err
	}

	return sdf.Difference3D(
		sdf.Difference3D(
			sdf.Union3D(wall, keyPlate, o, trrs),
			boardSpace,
		),
		sdf.Union3D(usb, trrs),
	), nil
}

func createWall(top, bottom, oled sdf.SDF2) sdf.SDF3 {
	return sdf.Transform3D(
		sdf.Extrude3D(sdf.Difference2D(sdf.Difference2D(sdf.Offset2D(bottom, buffer+wallThickness), sdf.Offset2D(top, buffer)), sdf.Offset2D(oled, buffer)), wallHeight),
		sdf.Translate3d(v3.Vec{X: 0, Y: 0, Z: wallHeight / 2}),
	)
}

func createKeyPlate(top, screw, key sdf.SDF2) sdf.SDF3 {
	fmt.Println(screw)
	return sdf.Transform3D(
		sdf.Extrude3D(
			sdf.Difference2D(
				sdf.Difference2D(
					sdf.Offset2D(top, buffer),
					key,
				),
				screw,
			),
			thickness,
		),
		sdf.Translate3d(
			v3.Vec{X: 0, Y: 0, Z: boardHeight + thickness*1.5},
		),
	)
}

func createBoardSpace(bottom sdf.SDF2) sdf.SDF3 {
	h := thickness + boardHeight
	return sdf.Transform3D(
		sdf.Extrude3D(sdf.Offset2D(bottom, buffer), h),
		sdf.Translate3d(v3.Vec{X: 0, Y: 0, Z: h / 2}),
	)
}
