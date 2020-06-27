package plotter

import (
	"fmt"
	"image/color"

	"gonum.org/v1/plot"
)

type LineOpts struct {
	XLabel string
	YLabel string
}

// Plotter defines struct for project
type Plotter struct {
	plot *plot.Plot
}

// New provides initialization of the plotter
func New() (*Plotter, error) {
	p, err := plot.New()
	if err != nil {
		return err
	}
	return &Plotter{
		plot: p,
	}
}

// Line provides creating and saving of the line plot
func (p *Plotter) Line(opts LineOpts) error {
	p.Title.Text = fmt.Sprintf("Memory Plot of PID %d", 10)
	p.X.Label.Text = opts.XLabel
	p.Y.Label.Text = opts.YLabel
	p.Add(plotter.NewGrid())
	rssLine, err := plotter.NewLine(pts)
	if err != nil {
		panic(err)
	}
	rssLine.LineStyle.Width = vg.Points(1)
	rssLine.LineStyle.Color = color.RGBA{R: 100, G: 100, B: 0, A: 255}
	p.Add(rssLine)
	p.Save(500, 500, "test.png")
}
