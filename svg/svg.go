package svg

import (
	"encoding/xml"
	"os"
)

type SVG struct {
	Width    string     `xml:"width,attr"`
	Height   string     `xml:"height,attr"`
	ViewBox  string     `xml:"viewBox,attr"`
	Fill     string     `xml:"fill,attr"`
	XMLNS    string     `xml:"xmlns,attr"`
	Graphics []*Graphic `xml:"g"`
}

func NewSVG(name string) (*SVG, error) {
	r, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	var svg SVG
	if err := xml.NewDecoder(r).Decode(&svg); err != nil {
		return nil, err
	}
	return &svg, nil
}

func (svg *SVG) Find(id string) *Graphic {
	for _, g := range svg.Graphics {
		if g.ID == id {
			return g
		}
	}
	return nil
}
