// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/gonum/plot"
	"github.com/gonum/plot/vg"
)

func ExampleBoxPlot() {
	// Create the sample data.
	rand.Seed(0) // The default random seed is 1.
	n := 100
	uniform := make(ValueLabels, n)
	normal := make(ValueLabels, n)
	expon := make(ValueLabels, n)
	for i := 0; i < n; i++ {
		uniform[i].Value = rand.Float64()
		uniform[i].Label = fmt.Sprintf("%4.4f", uniform[i].Value)
		normal[i].Value = rand.NormFloat64()
		normal[i].Label = fmt.Sprintf("%4.4f", normal[i].Value)
		expon[i].Value = rand.ExpFloat64()
		expon[i].Label = fmt.Sprintf("%4.4f", expon[i].Value)
	}

	// Make boxes for our data and add them to the plot.
	uniBox, err := NewBoxPlot(vg.Points(20), 0, uniform)
	handleEx(err)
	normBox, err := NewBoxPlot(vg.Points(20), 1, normal)
	handleEx(err)
	expBox, err := NewBoxPlot(vg.Points(20), 2, expon)
	handleEx(err)

	// Make a vertical box plot.
	uniLabels, err := uniBox.OutsideLabels(uniform)
	handleEx(err)
	normLabels, err := normBox.OutsideLabels(normal)
	handleEx(err)
	expLabels, err := expBox.OutsideLabels(expon)
	handleEx(err)

	p1, err := plot.New()
	handleEx(err)
	p1.Title.Text = "Vertical Box Plot"
	p1.Y.Label.Text = "plotter.Values"
	p1.Add(uniBox, uniLabels, normBox, normLabels, expBox, expLabels)

	// Set the X axis of the plot to nominal with
	// the given names for x=0, x=1 and x=2.
	p1.NominalX("Uniform\nDistribution", "Normal\nDistribution",
		"Exponential\nDistribution")

	// Normally, you would use plot.Save() instead of checkPlot().
	checkPlot("examplePlots", "verticalBoxPlot", "png", p1, 200, 200,
		handleEx, exampleLog)

	// Now, make the same plot but horizontal.
	normBox.Horizontal = true
	expBox.Horizontal = true
	uniBox.Horizontal = true
	// We can use the same plotters but the labels need to be recreated.
	uniLabels, err = uniBox.OutsideLabels(uniform)
	handleEx(err)
	normLabels, err = normBox.OutsideLabels(normal)
	handleEx(err)
	expLabels, err = expBox.OutsideLabels(expon)
	handleEx(err)

	p2, err := plot.New()
	handleEx(err)
	p2.Title.Text = "Horizontal Box Plot"
	p2.X.Label.Text = "plotter.Values"

	p2.Add(uniBox, uniLabels, normBox, normLabels, expBox, expLabels)

	// Set the Y axis of the plot to nominal with
	// the given names for y=0, y=1 and y=2.
	p2.NominalY("Uniform\nDistribution", "Normal\nDistribution",
		"Exponential\nDistribution")

	// Normally, you would use plot.Save() instead of checkPlot().
	checkPlot("examplePlots", "horizontalBoxPlot", "png", p2, 200, 200,
		handleEx, exampleLog)

	// Output:
	// Image can be seen at https://github.com/gonum/plot/tree/master/plotter/examplePlots/verticalBoxPlot.png.
	// Image can be seen at https://github.com/gonum/plot/tree/master/plotter/examplePlots/horizontalBoxPlot.png.
}

func TestBoxPlot(t *testing.T) {
	// Create sample data.
	rand.Seed(0) // The default random seed is 1.
	n := 100
	uniform := make(ValueLabels, n)
	normal := make(ValueLabels, n)
	expon := make(ValueLabels, n)
	for i := 0; i < n; i++ {
		uniform[i].Value = rand.Float64()
		uniform[i].Label = fmt.Sprintf("%4.4f", uniform[i].Value)
		normal[i].Value = rand.NormFloat64()
		normal[i].Label = fmt.Sprintf("%4.4f", normal[i].Value)
		expon[i].Value = rand.ExpFloat64()
		expon[i].Label = fmt.Sprintf("%4.4f", expon[i].Value)
	}

	// Make boxes for our data and add them to the plot.
	uniBox, err := NewBoxPlot(vg.Points(20), 0, uniform)
	handleTest(t)(err)
	normBox, err := NewBoxPlot(vg.Points(20), 1, normal)
	handleTest(t)(err)
	expBox, err := NewBoxPlot(vg.Points(20), 2, expon)
	handleTest(t)(err)

	// Make a vertical box plot.
	uniLabels, err := uniBox.OutsideLabels(uniform)
	handleTest(t)(err)
	normLabels, err := normBox.OutsideLabels(normal)
	handleTest(t)(err)
	expLabels, err := expBox.OutsideLabels(expon)
	handleTest(t)(err)

	p1, err := plot.New()
	handleTest(t)(err)
	p1.Title.Text = "Vertical Box Plot"
	p1.Y.Label.Text = "plotter.Values"
	p1.Add(uniBox, uniLabels, normBox, normLabels, expBox, expLabels)

	// Add a GlyphBox plotter for debugging.
	p1.Add(NewGlyphBoxes())

	// Set the X axis of the plot to nominal with
	// the given names for x=0, x=1 and x=2.
	p1.NominalX("Uniform\nDistribution", "Normal\nDistribution",
		"Exponential\nDistribution")

	checkPlot("examplePlots", "verticalBoxPlotTest", "png", p1, 200, 200,
		handleTest(t), testLog(t))

	// Now, make the same plot but horizontal.
	normBox.Horizontal = true
	expBox.Horizontal = true
	uniBox.Horizontal = true
	// We can use the same plotters but the labels need to be recreated.
	uniLabels, err = uniBox.OutsideLabels(uniform)
	handleTest(t)(err)
	normLabels, err = normBox.OutsideLabels(normal)
	handleTest(t)(err)
	expLabels, err = expBox.OutsideLabels(expon)
	handleTest(t)(err)

	p2, err := plot.New()
	handleTest(t)(err)
	p2.Title.Text = "Horizontal Box Plot"
	p2.X.Label.Text = "plotter.Values"

	p2.Add(uniBox, uniLabels, normBox, normLabels, expBox, expLabels)

	// Add a GlyphBox plotter for debugging.
	p2.Add(NewGlyphBoxes())

	// Set the Y axis of the plot to nominal with
	// the given names for y=0, y=1 and y=2.
	p2.NominalY("Uniform\nDistribution", "Normal\nDistribution",
		"Exponential\nDistribution")

	checkPlot("examplePlots", "horizontalBoxPlotTest", "png", p2, 200, 200,
		handleTest(t), testLog(t))
}

func TestGroupedBoxPlots(t *testing.T) {
	rand.Seed(0) // The default random seed is 1.
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

	w := vg.Points(20)
	for x := 0.0; x < 3.0; x++ {
		b0, err := NewBoxPlot(w, x, uniform)
		handleTest(t)(err)
		b0.Offset = -w - vg.Points(3)
		b1, err := NewBoxPlot(w, x, normal)
		handleTest(t)(err)
		b2, err := NewBoxPlot(w, x, expon)
		handleTest(t)(err)
		b2.Offset = w + vg.Points(3)
		p.Add(b0, b1, b2)
	}
	// Add a GlyphBox plotter for debugging.
	p.Add(NewGlyphBoxes())

	// Set the X axis of the plot to nominal with
	// the given names for x=0, x=1 and x=2.
	p.NominalX("Group 0", "Group 1", "Group 2")
	checkPlot("examplePlots", "groupedBoxPlotTest", "png", p, 300, 300,
		handleTest(t), testLog(t))
}
