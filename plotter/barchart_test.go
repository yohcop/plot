// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"image/color"
	"testing"

	"github.com/gonum/plot"
	"github.com/gonum/plot/vg"
)

func ExampleBarChart() {
	// Create the plot values and labels.
	values := Values{0.5, 10, 20, 30}
	verticalLabels := []string{"A", "B", "C", "D"}
	horizontalLabels := []string{"Label A", "Label B", "Label C", "Label D"}

	// Create a vertical BarChart
	p1, err := plot.New()
	handleEx(err)
	verticalBarChart, err := NewBarChart(values, 0.5*vg.Centimeter)
	handleEx(err)
	p1.Add(verticalBarChart)
	p1.NominalX(verticalLabels...)
	checkPlot("examplePlots", "verticalBarChart", "png", p1, 100, 100,
		handleEx, exampleLog)

	// Create a horizontal BarChart
	p2, err := plot.New()
	handleEx(err)
	horizontalBarChart, err := NewBarChart(values, 0.5*vg.Centimeter)
	horizontalBarChart.Horizontal = true // Specify a horizontal BarChart.
	handleEx(err)
	p2.Add(horizontalBarChart)
	p2.NominalY(horizontalLabels...)
	checkPlot("examplePlots", "horizontalBarChart", "png", p2, 100, 100,
		handleEx, exampleLog)

	// Output:
	// Image can be seen at https://github.com/gonum/plot/tree/master/plotter/examplePlots/verticalBarChart.png.
	// Normally, you would use plot.Save().
	// Image can be seen at https://github.com/gonum/plot/tree/master/plotter/examplePlots/horizontalBarChart.png.
	// Normally, you would use plot.Save().

}

func TestBarChart(t *testing.T) {
	// Create the plot values and labels.
	values := Values{0.5, 10, 20, 30}
	verticalLabels := []string{"A", "B", "C", "D"}
	horizontalLabels := []string{"Label A", "Label B", "Label C", "Label D"}

	// Create a vertical BarChart
	p1, err := plot.New()
	handleTest(t)(err)
	verticalBarChart, err := NewBarChart(values, 0.5*vg.Centimeter)
	handleTest(t)(err)
	p1.Add(verticalBarChart)
	p1.NominalX(verticalLabels...)
	// Add a GlyphBox plotter for debugging.
	p1.Add(NewGlyphBoxes())
	checkPlot("examplePlots", "verticalBarChartTest", "png", p1, 100, 100,
		handleTest(t), testLog(t))

	// Create a horizontal BarChart
	p2, err := plot.New()
	handleTest(t)(err)
	horizontalBarChart, err := NewBarChart(values, 0.5*vg.Centimeter)
	horizontalBarChart.Horizontal = true // Specify a horizontal BarChart.
	handleTest(t)(err)
	p2.Add(horizontalBarChart)
	p2.NominalY(horizontalLabels...)
	// Add a GlyphBox plotter for debugging.
	p2.Add(NewGlyphBoxes())
	checkPlot("examplePlots", "horizontalBarChartTest", "png", p2, 100, 100,
		handleTest(t), testLog(t))
}

// An example of making a bar chart.
func TestBarChart2(t *testing.T) {
	groupA := Values{20, 35, 30, 35, 27}
	groupB := Values{25, 32, 34, 20, 25}
	groupC := Values{12, 28, 15, 21, 8}
	groupD := Values{30, 42, 6, 9, 12}

	p, err := plot.New()
	handleTest(t)(err)
	p.Title.Text = "Bar chart"
	p.Y.Label.Text = "Heights"

	w := vg.Points(8)

	barsA, err := NewBarChart(groupA, w)
	handleTest(t)(err)
	barsA.Color = color.RGBA{R: 255, A: 255}
	barsA.Offset = -w / 2

	barsB, err := NewBarChart(groupB, w)
	handleTest(t)(err)
	barsB.Color = color.RGBA{R: 196, G: 196, A: 255}
	barsB.Offset = w / 2

	barsC, err := NewBarChart(groupC, w)
	handleTest(t)(err)
	barsC.XMin = 6
	barsC.Color = color.RGBA{B: 255, A: 255}
	barsC.Offset = -w / 2

	barsD, err := NewBarChart(groupD, w)
	handleTest(t)(err)
	barsD.Color = color.RGBA{B: 255, R: 255, A: 255}
	barsD.XMin = 6
	barsD.Offset = w / 2

	p.Add(barsA, barsB, barsC, barsD)
	p.Legend.Add("A", barsA)
	p.Legend.Add("B", barsB)
	p.Legend.Add("C", barsC)
	p.Legend.Add("D", barsD)
	p.Legend.Top = true
	p.NominalX("Zero", "One", "Two", "Three", "Four", "",
		"Six", "Seven", "Eight", "Nine", "Ten")

	p.Add(NewGlyphBoxes())
	checkPlot("examplePlots", "barChartTest2", "png", p, 300, 250,
		handleTest(t), testLog(t))
}

// An example of making a stacked bar chart.
func TestStackedBarChart(t *testing.T) {
	groupA := Values{20, 35, 30, 35, 27}
	groupB := Values{25, 32, 34, 20, 25}
	groupC := Values{12, 28, 15, 21, 8}
	groupD := Values{30, 42, 6, 9, 12}

	p, err := plot.New()
	handleTest(t)(err)
	p.Title.Text = "Bar chart"
	p.Y.Label.Text = "Heights"

	w := vg.Points(15)

	barsA, err := NewBarChart(groupA, w)
	handleTest(t)(err)
	barsA.Color = color.RGBA{R: 255, A: 255}
	barsA.Offset = -w / 2

	barsB, err := NewBarChart(groupB, w)
	handleTest(t)(err)
	barsB.Color = color.RGBA{R: 196, G: 196, A: 255}
	barsB.StackOn(barsA)

	barsC, err := NewBarChart(groupC, w)
	handleTest(t)(err)
	barsC.Offset = w / 2
	barsC.Color = color.RGBA{B: 255, A: 255}

	barsD, err := NewBarChart(groupD, w)
	handleTest(t)(err)
	barsD.StackOn(barsC)
	barsD.Color = color.RGBA{B: 255, R: 255, A: 255}

	p.Add(barsA, barsB, barsC, barsD)
	p.Legend.Add("A", barsA)
	p.Legend.Add("B", barsB)
	p.Legend.Add("C", barsC)
	p.Legend.Add("D", barsD)
	p.Legend.Top = true
	p.NominalX("Zero", "One", "Two", "Three", "Four", "",
		"Six", "Seven", "Eight", "Nine", "Ten")

	p.Add(NewGlyphBoxes())
	checkPlot("examplePlots", "stackedBarChartTest", "png", p, 250, 250,
		handleTest(t), testLog(t))
}
