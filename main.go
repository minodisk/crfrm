package main

import (
	"log"

	"github.com/minodisk/crfrm/svg"

	"github.com/deadsy/sdfx/render"
	"github.com/deadsy/sdfx/sdf"
	v3 "github.com/deadsy/sdfx/vec/v3"
)

var (
	buffer         = 0.2
	thickness      = 2.0
	wallThickness  = 2.4
	boardHeight    = 4.0
	wallHeight     = 13.8
	oledWallHeight = 13.8
	oledHeight     = 2.0
)

func main() {
	svg, err := svg.NewSVG("./assets.svg")
	if err != nil {
		log.Fatal(err)
	}

	f, err := top(svg)
	if err != nil {
		log.Fatal(err)
	}
	render.ToSTL(f, "crfrm.stl", render.NewMarchingCubesUniform(1000))
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
	oled, err := svg.Find("oled").Draw()
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

	wall := createWall(top, bottom, oled)
	keyPlate := createKeyPlate(top, screw, key)
	boardSpace := createBoardSpace(bottom)
	o := createOLED(oled)
	return sdf.Difference3D(sdf.Union3D(wall, keyPlate, o), boardSpace), nil
}

func createWall(top, bottom, oled sdf.SDF2) sdf.SDF3 {
	return sdf.Transform3D(
		sdf.Extrude3D(sdf.Difference2D(sdf.Difference2D(sdf.Offset2D(bottom, buffer+wallThickness), sdf.Offset2D(top, buffer)), sdf.Offset2D(oled, buffer)), wallHeight),
		sdf.Translate3d(v3.Vec{X: 0, Y: 0, Z: wallHeight / 2}),
	)
}

func createKeyPlate(top, screw, key sdf.SDF2) sdf.SDF3 {
	return sdf.Transform3D(
		sdf.Extrude3D(
			sdf.Difference2D(sdf.Difference2D(sdf.Offset2D(top, buffer), key), screw),
			thickness,
		),
		sdf.Translate3d(
			v3.Vec{X: 0, Y: 0, Z: boardHeight + thickness*1.5},
		),
	)
}

func createOLED(oled sdf.SDF2) sdf.SDF3 {
	return sdf.Transform3D(
		sdf.Extrude3D(sdf.Difference2D(oled, sdf.Offset2D(oled, -1.0)), oledHeight),
		sdf.Translate3d(v3.Vec{X: 0, Y: 0, Z: wallHeight - oledHeight/2}),
	)
}

func createBoardSpace(bottom sdf.SDF2) sdf.SDF3 {
	h := thickness + boardHeight
	return sdf.Transform3D(
		sdf.Extrude3D(sdf.Offset2D(bottom, buffer), h),
		sdf.Translate3d(v3.Vec{X: 0, Y: 0, Z: h / 2}),
	)
}
