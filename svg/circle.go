package svg

import (
	"github.com/deadsy/sdfx/sdf"
	v2 "github.com/deadsy/sdfx/vec/v2"
)

type Circle struct {
	CenterX float64 `xml:"cx,attr"`
	CenterY float64 `xml:"cy,attr"`
	Radius  float64 `xml:"r,attr"`
}

func (c *Circle) Draw() (sdf.SDF2, error) {
	c2, err := sdf.Circle2D(c.Radius)
	if err != nil {
		return nil, err
	}
	return sdf.Transform2D(
		c2,
		sdf.Translate2d(v2.Vec{X: c.CenterX, Y: c.CenterY}),
	), nil
}
