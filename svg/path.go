package svg

import (
	"fmt"
	"strconv"

	"github.com/deadsy/sdfx/sdf"
	"github.com/pkg/errors"
)

type Path struct {
	Draw           string `xml:"d,attr"`
	Stroke         string `xml:"stroke,attr"`
	StrokeLinecap  string `xml:"stroke-linecap,attr"`
	StrokeLineJoin string `xml:"stroke-linejoin,attr"`
}

func (p *Path) ToBezier() (*sdf.Bezier, error) {
	b := sdf.NewBezier()
	d := p.Draw

	l := len(d)
	for i := 0; i < l; i++ {
		switch d[i] {
		case 'M':
			p, j, err := readPoint2D(d, i+1, l)
			if err != nil {
				return nil, err
			}
			i = j
			NewMoveTo(p).DrawTo(b)
		case 'L':
			p, j, err := readPoint2D(d, i+1, l)
			if err != nil {
				return nil, err
			}
			i = j
			NewLineTo(p).DrawTo(b)
		case 'C':
			start, j, err := readPoint2D(d, i+1, l)
			if err != nil {
				return nil, err
			}
			mid, j, err := readPoint2D(d, j+2, l)
			if err != nil {
				return nil, err
			}
			end, j, err := readPoint2D(d, j+2, l)
			if err != nil {
				return nil, err
			}
			i = j
			NewCubicBezier(start, mid, end).DrawTo(b)
		case 'Z':
			NewClosePath().DrawTo(b)
		}
	}

	return b, nil
}

func readPoint2D(d string, start int, last int) (*Point2D, int, error) {
	x, j, err := readFloat64(d, start, last)
	if err != nil {
		return nil, j, err
	}
	y, j, err := readFloat64(d, j+2, last)
	if err != nil {
		return nil, j, err
	}
	return NewPoint2D(x, y), j, err
}

func readFloat64(d string, start int, last int) (float64, int, error) {
	n := ""
	i := start
	for ; i < last; i++ {
		c := d[i]
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.':
			n = n + string(c)
		default:
			f, err := strconv.ParseFloat(n, 64)
			return f, i - 1, errors.Wrap(err, fmt.Sprintf("fail to parse float %s", n))
		}
	}

	f, err := strconv.ParseFloat(n, 64)
	return f, i, err
}

type Point2D struct {
	X float64
	Y float64
}

func NewPoint2D(x float64, y float64) *Point2D {
	return &Point2D{x, y}
}

type Drawable interface {
	DrawTo(b *sdf.Bezier)
}

type MoveTo struct {
	Point *Point2D
}

func NewMoveTo(p *Point2D) *MoveTo {
	return &MoveTo{p}
}

func (m *MoveTo) DrawTo(b *sdf.Bezier) {
	b.Add(m.Point.X, m.Point.Y)
}

type LineTo struct {
	Point *Point2D
}

func NewLineTo(p *Point2D) *LineTo {
	return &LineTo{p}
}

func (l *LineTo) DrawTo(b *sdf.Bezier) {
	b.Add(l.Point.X, l.Point.Y)
}

type CubicBezier struct {
	Start *Point2D
	Mid   *Point2D
	End   *Point2D
}

func NewCubicBezier(start *Point2D, mid *Point2D, end *Point2D) *CubicBezier {
	return &CubicBezier{start, mid, end}
}

func (m *CubicBezier) DrawTo(b *sdf.Bezier) {
	b.Add(m.Start.X, m.Start.Y).Mid()
	b.Add(m.Mid.X, m.Mid.Y).Mid()
	b.Add(m.End.X, m.End.Y)
}

type ClosePath struct{}

func NewClosePath() *ClosePath {
	return &ClosePath{}
}

func (m *ClosePath) DrawTo(b *sdf.Bezier) {
	b.Close()
}
