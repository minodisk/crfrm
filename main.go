package main

import (
	"log"

	"github.com/minodisk/crfrm/board"
	"github.com/minodisk/crfrm/controller"
	"github.com/minodisk/crfrm/svg"
	"github.com/minodisk/crfrm/top"
	"github.com/minodisk/crfrm/trrs"
	"github.com/minodisk/crfrm/usb"
	"github.com/minodisk/crfrm/wall"
	"github.com/pkg/errors"

	"github.com/deadsy/sdfx/render"
	"github.com/deadsy/sdfx/sdf"
)

var (
	// meshCells = 400
	meshCells     = 1000
	buffer        = 0.5
	pad           = 0.6
	thickness     = 1.6
	wallThickness = 1.6
	boardHeight   = 4.0
	wallHeight    = 13.0
	skirtHeight   = 1.0
	topZ          = boardHeight + thickness*1.5
)

func main() {
	svg, err := svg.NewSVG("./assets.svg")
	if err != nil {
		log.Fatal(errors.Wrap(err, "fail to parse SVG"))
	}

	left, err := create(svg, true)
	if err != nil {
		log.Fatal(errors.Wrap(err, "fail to create left"))
	}

	r, err := create(svg, false)
	if err != nil {
		log.Fatal(errors.Wrap(err, "fail to create right"))
	}
	right := sdf.Transform3D(
		r,
		sdf.MirrorYZ(),
	)

	render.ToSTL(left, "left.stl", render.NewMarchingCubesUniform(meshCells))
	render.ToSTL(right, "right.stl", render.NewMarchingCubesUniform(meshCells))
}

func create(svg *svg.SVG, withUSB bool) (sdf.SDF3, error) {
	topPath, err := svg.Find("top").Draw()
	if err != nil {
		return nil, err
	}
	bottomPath, err := svg.Find("bottom").Draw()
	if err != nil {
		return nil, err
	}
	screwShape, err := svg.Find("screw").Draw()
	if err != nil {
		return nil, err
	}
	keyShape, err := svg.Find("key").Draw()
	if err != nil {
		return nil, err
	}

	topShape := sdf.Offset2D(topPath, buffer)
	bottomShape := sdf.Offset2D(bottomPath, buffer)
	controllerShape := sdf.Difference2D(bottomShape, topShape)

	wa := wall.Create(
		bottomShape,
		wallThickness, wallHeight, skirtHeight,
	)
	to := top.Create(
		topShape, screwShape, keyShape,
		thickness, topZ,
	)
	ctrl := controller.Create(
		controllerShape,
		thickness, wallHeight, topZ,
	)
	// jack, err := controller.CreateJack(
	// 	thickness, wallThickness, wallHeight,
	// )
	// if err != nil {
	// 	return nil, err
	// }
	boardSpace := board.Create(
		bottomShape,
		thickness, boardHeight,
	)
	u := usb.CreateHole()
	t, err := trrs.CreateHole()
	if err != nil {
		return nil, err
	}

	var hole sdf.SDF3
	if withUSB {
		hole = sdf.Union3D(u, t)
	} else {
		hole = t
	}

	return sdf.Difference3D(
		sdf.Union3D(
			sdf.Difference3D(
				wa,
				hole,
			),
			to,
			ctrl,
		),
		// sdf.Union3D(
		boardSpace,
		// 	jack,
		// ),
	), nil
}
