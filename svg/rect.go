package svg

import (
	"github.com/deadsy/sdfx/sdf"
	v2 "github.com/deadsy/sdfx/vec/v2"
)

type Rect struct {
	X      float64 `xml:"x,attr"`
	Y      float64 `xml:"y,attr"`
	Width  float64 `xml:"width,attr"`
	Height float64 `xml:"height,attr"`
}

func (r *Rect) Draw() (sdf.SDF2, error) {
	return sdf.Transform2D(
		sdf.Box2D(v2.Vec{X: r.Width, Y: r.Height}, 0),
		sdf.Translate2d(v2.Vec{X: r.X + r.Width/2, Y: r.Y + r.Height/2}),
	), nil
}
