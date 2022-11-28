package main

import (
	"log"

	"github.com/minodisk/crfrm/svg"

	"github.com/deadsy/sdfx/render"
	"github.com/deadsy/sdfx/sdf"
	v3 "github.com/deadsy/sdfx/vec/v3"
)

var (
	thickness     = 2.0
	wallThickness = 2.4
	boardHeight   = 4.0
	wallHeight    = 16.8
	oledHeight    = 2.0
)

func top(svg *svg.SVG) (sdf.SDF3, error) {
	k, err := keyPlate(svg)
	if err != nil {
		return nil, err
	}
	// b, err := bottom(svg)
	// if err != nil {
	// 	return nil, err
	// }
	w, err := wall(svg)
	if err != nil {
		return nil, err
	}
	bo, err := board(svg)
	if err != nil {
		return nil, err
	}
	o, err := oled(svg)
	if err != nil {
		return nil, err
	}
	return sdf.Difference3D(sdf.Union3D(k /*b,*/, w, o), bo), nil
}

func keyPlate(svg *svg.SVG) (sdf.SDF3, error) {
	top, err := svg.Find("top").Draw()
	if err != nil {
		return nil, err
	}
	screw, err := svg.Find("screw").Draw()
	if err != nil {
		return nil, err
	}
	base := sdf.Offset2D(sdf.Difference2D(top, screw), thickness)

	keyHole, err := svg.Find("key").Draw()
	if err != nil {
		return nil, err
	}

	return sdf.Transform3D(
		sdf.Extrude3D(
			sdf.Difference2D(base, keyHole),
			thickness,
		),
		sdf.Translate3d(
			v3.Vec{X: 0, Y: 0, Z: boardHeight + thickness*1.5},
		),
	), nil
}

func bottom(svg *svg.SVG) (sdf.SDF3, error) {
	inner, err := svg.Find("bottom").Draw()
	if err != nil {
		return nil, err
	}
	screw, err := svg.Find("screw").Draw()
	if err != nil {
		return nil, err
	}
	return sdf.Transform3D(sdf.Extrude3D(sdf.Difference2D(inner, screw), thickness), sdf.Translate3d(v3.Vec{X: 0, Y: 0, Z: thickness / 2})), nil
}

func wall(svg *svg.SVG) (sdf.SDF3, error) {
	bottom, err := svg.Find("bottom").Draw()
	if err != nil {
		return nil, err
	}
	inner := sdf.Extrude3D(bottom, wallHeight)
	outer := sdf.Extrude3D(sdf.Offset2D(bottom, wallThickness), wallHeight)
	wall := sdf.Transform3D(sdf.Difference3D(outer, inner), sdf.Translate3d(v3.Vec{X: 0, Y: 0, Z: wallHeight / 2}))
	return wall, nil
}

func oled(svg *svg.SVG) (sdf.SDF3, error) {
	bottom, err := svg.Find("bottom").Draw()
	if err != nil {
		return nil, err
	}
	top, err := svg.Find("top").Draw()
	if err != nil {
		return nil, err
	}
	outer := sdf.Difference2D(bottom, top)
	oled, err := svg.Find("oled").Draw()
	if err != nil {
		return nil, err
	}

	wallZ := thickness + boardHeight
	wallH := wallHeight - wallZ
	wall := sdf.Transform3D(
		sdf.Extrude3D(
			sdf.Difference2D(outer, oled),
			wallH,
		),
		sdf.Translate3d(v3.Vec{X: 0, Y: 0, Z: wallZ + wallH/2}),
	)

	frame := sdf.Transform3D(
		sdf.Extrude3D(sdf.Difference2D(oled, sdf.Offset2D(oled, -2.0)), oledHeight),
		sdf.Translate3d(v3.Vec{X: 0, Y: 0, Z: wallHeight - thickness - thickness/2}),
	)
	return sdf.Union3D(wall, frame), nil
}

func board(svg *svg.SVG) (sdf.SDF3, error) {
	bottom, err := svg.Find("bottom").Draw()
	if err != nil {
		return nil, err
	}
	return sdf.Transform3D(
		sdf.Extrude3D(bottom, boardHeight),
		sdf.Translate3d(v3.Vec{X: 0, Y: 0, Z: thickness + boardHeight/2}),
	), nil
}

func main() {
	svg, err := svg.NewSVG("./assets.svg")
	if err != nil {
		log.Fatal(err)
	}

	f, err := top(svg)
	if err != nil {
		log.Fatal(err)
	}
	render.ToSTL(f, "crfrm.stl", render.NewMarchingCubesUniform(300))
}
