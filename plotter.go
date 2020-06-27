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
	XLabel   string
	YLabel   string
	Data     plotter.XYs
	Filename string
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
	p.plot.Title.Text = fmt.Sprintf("Memory Plot of PID %d", 10)
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
	if err := p.plot.Save(500, 500, "test.png"); err != nil {
		return fmt.Errorf("unable to save plot: %s %v", opts.Filename, err)
	}
	return nil
}
