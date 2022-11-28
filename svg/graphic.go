package svg

import "github.com/deadsy/sdfx/sdf"

type Graphic struct {
	ID    string  `xml:"id,attr"`
	Paths []*Path `xml:"path"`
}

func (g *Graphic) Draw() (sdf.SDF2, error) {
	var polygons []sdf.SDF2
	for _, path := range g.Paths {
		b, err := path.ToBezier()
		if err != nil {
			return nil, err
		}
		p, err := b.Polygon()
		if err != nil {
			return nil, err
		}
		polygon, err := sdf.Polygon2D(p.Vertices())
		if err != nil {
			return nil, err
		}
		polygons = append(polygons, polygon)
	}
	return sdf.Union2D(polygons...), nil
}
