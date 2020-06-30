package plotter

import (
	"fmt"
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// LineOpts defines options for creating line plot
type LineOpts struct {
	XLabel    string
	YLabel    string
	Data      plotter.XYs
	ImageName string
	Text      string
	ImgWidth  float64
	ImgHeight float64
}

// Validate provides validation of the LineOpts
func (l LineOpts) Validate() error {
	if l.Text == "" {
		return fmt.Errorf("text is not defined")
	}
	if l.ImgWidth == 0 {
		return fmt.Errorf("ingWidth is not defined")
	}
	if l.ImgHeight == 0 {
		return fmt.Errorf("imgHeight is not defined")
	}
	return nil
}

// Plotter defines struct for project
type Plotter struct {
	plot *plot.Plot
}

// New provides initialization of the plotter
func New() (*Plotter, error) {
	p, err := plot.New()
	if err != nil {
		return nil, err
	}
	return &Plotter{
		plot: p,
	}, nil
}

// Line provides creating and saving of the line plot
func (p *Plotter) Line(opts LineOpts) error {
	if err := opts.Validate(); err != nil {
		return err
	}
	p.plot.Title.Text = opts.Text
	p.plot.X.Label.Text = opts.XLabel
	p.plot.Y.Label.Text = opts.YLabel
	p.plot.Add(plotter.NewGrid())
	rssLine, err := plotter.NewLine(opts.Data)
	if err != nil {
		return fmt.Errorf("unable to create new line: %v", err)
	}
	rssLine.LineStyle.Width = vg.Points(1)
	rssLine.LineStyle.Color = color.RGBA{R: 100, G: 100, B: 0, A: 255}
	p.plot.Add(rssLine)
	if err := p.plot.Save(vg.Length(opts.ImgWidth), vg.Length(opts.ImgHeight), opts.ImageName); err != nil {
		return fmt.Errorf("unable to save plot: %s %v", opts.ImageName, err)
	}
	return nil
}
