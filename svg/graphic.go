package svg

import "github.com/deadsy/sdfx/sdf"

type Graphic struct {
	ID      string    `xml:"id,attr"`
	Paths   []*Path   `xml:"path"`
	Rects   []*Rect   `xml:"rect"`
	Circles []*Circle `xml:"circle"`
}

func (g *Graphic) Draw() (sdf.SDF2, error) {
	var polygons []sdf.SDF2
	for _, path := range g.Paths {
		bs, err := path.ToBeziers()
		if err != nil {
			return nil, err
		}
		for _, b := range bs {
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
	}
	for _, r := range g.Rects {
		p, err := r.Draw()
		if err != nil {
			return nil, err
		}
		polygons = append(polygons, p)
	}
	for _, c := range g.Circles {
		p, err := c.Draw()
		if err != nil {
			return nil, err
		}
		polygons = append(polygons, p)
	}
	return sdf.Union2D(polygons...), nil
}
