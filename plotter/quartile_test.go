// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"math/rand"
	"testing"

	"github.com/gonum/plot"
	"github.com/gonum/plot/vg"
)

func ExampleQuartPlot() {
	// Create the example data.
	rand.Seed(int64(0))
	n := 100
	uniform := make(Values, n)
	normal := make(Values, n)
	expon := make(Values, n)
	for i := 0; i < n; i++ {
		uniform[i] = rand.Float64()
		normal[i] = rand.NormFloat64()
		expon[i] = rand.ExpFloat64()
	}

	// Create the QuartPlots
	qp1, err := NewQuartPlot(0, uniform)
	handleEx(err)
	qp2, err := NewQuartPlot(1, normal)
	handleEx(err)
	qp3, err := NewQuartPlot(2, expon)
	handleEx(err)

	// Create a vertical plot
	p1, err := plot.New()
	handleEx(err)
	p1.Title.Text = "Quartile Plot"
	p1.Y.Label.Text = "plotter.Values"
	p1.Add(qp1, qp2, qp3)

	// Set the X axis of the plot to nominal with
	// the given names for x=0, x=1 and x=2.
	p1.NominalX("Uniform\nDistribution", "Normal\nDistribution",
		"Exponential\nDistribution")

	checkPlot("examplePlots", "verticalQuartPlot", "png", p1, 200, 200,
		handleEx, exampleLog)

	// Create a horizontal plot
	qp1.Horizontal = true
	qp2.Horizontal = true
	qp3.Horizontal = true

	p2, err := plot.New()
	handleEx(err)
	p2.Title.Text = "Quartile Plot"
	p2.X.Label.Text = "plotter.Values"
	p2.Add(qp1, qp2, qp3)

	// Set the Y axis of the plot to nominal with
	// the given names for y=0, y=1 and y=2.
	p2.NominalY("Uniform\nDistribution", "Normal\nDistribution",
		"Exponential\nDistribution")

	checkPlot("examplePlots", "horizontalQuartPlot", "png", p2, 200, 200,
		handleEx, exampleLog)

	// Output:
	// Image can be seen at https://github.com/gonum/plot/plotter/examplePlots/verticalQuartPlot.png.
	// Normally, you would use plot.Save().
	// Image can be seen at https://github.com/gonum/plot/plotter/examplePlots/horizontalQuartPlot.png.
	// Normally, you would use plot.Save().
}

func TestQuartPlot(t *testing.T) {
	// Create the example data.
	rand.Seed(int64(0))
	n := 100
	uniform := make(Values, n)
	normal := make(Values, n)
	expon := make(Values, n)
	for i := 0; i < n; i++ {
		uniform[i] = rand.Float64()
		normal[i] = rand.NormFloat64()
		expon[i] = rand.ExpFloat64()
	}

	// Create the QuartPlots
	qp1, err := NewQuartPlot(0, uniform)
	handleTest(t)(err)
	qp2, err := NewQuartPlot(1, normal)
	handleTest(t)(err)
	qp3, err := NewQuartPlot(2, expon)
	handleTest(t)(err)

	// Create a vertical plot
	p1, err := plot.New()
	handleTest(t)(err)
	p1.Title.Text = "Quartile Plot"
	p1.Y.Label.Text = "plotter.Values"
	p1.Add(qp1, qp2, qp3)

	// Add a GlyphBox plotter for debugging.
	p1.Add(NewGlyphBoxes())

	// Set the X axis of the plot to nominal with
	// the given names for x=0, x=1 and x=2.
	p1.NominalX("Uniform\nDistribution", "Normal\nDistribution",
		"Exponential\nDistribution")

	checkPlot("examplePlots", "verticalQuartPlotTest", "png", p1, 200, 200,
		handleTest(t), testLog(t))

	// Create a horizontal plot
	qp1.Horizontal = true
	qp2.Horizontal = true
	qp3.Horizontal = true

	p2, err := plot.New()
	handleTest(t)(err)
	p2.Title.Text = "Quartile Plot"
	p2.X.Label.Text = "plotter.Values"
	p2.Add(qp1, qp2, qp3)

	// Add a GlyphBox plotter for debugging.
	p2.Add(NewGlyphBoxes())

	// Set the Y axis of the plot to nominal with
	// the given names for y=0, y=1 and y=2.
	p2.NominalY("Uniform\nDistribution", "Normal\nDistribution",
		"Exponential\nDistribution")

	checkPlot("examplePlots", "horizontalQuartPlotTest", "png", p2, 200, 200,
		handleTest(t), testLog(t))
}

func TestGroupedQuartPlots(t *testing.T) {
	rand.Seed(int64(0))
	n := 100
	uniform := make(Values, n)
	normal := make(Values, n)
	expon := make(Values, n)
	for i := 0; i < n; i++ {
		uniform[i] = rand.Float64()
		normal[i] = rand.NormFloat64()
		expon[i] = rand.ExpFloat64()
	}

	p, err := plot.New()
	handleTest(t)(err)
	p.Title.Text = "Box Plot"
	p.Y.Label.Text = "plotter.Values"

	w := vg.Points(10)
	for x := 0.0; x < 3.0; x++ {
		b0, err := NewQuartPlot(x, uniform)
		handleTest(t)(err)
		b0.Offset = -w
		b1, err := NewQuartPlot(x, normal)
		handleTest(t)(err)
		b2, err := NewQuartPlot(x, expon)
		handleTest(t)(err)
		b2.Offset = w
		p.Add(b0, b1, b2)
	}
	p.Add(NewGlyphBoxes())

	p.NominalX("Group 0", "Group 1", "Group 2")

	checkPlot("examplePlots", "groupedQuartPlotTest", "png", p, 200, 200,
		handleTest(t), testLog(t))
}
