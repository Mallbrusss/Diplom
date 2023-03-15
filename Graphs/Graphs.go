package graphs

import (
	"image/color"
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	prs "programm/CsvPackges/Parse"
)

func Graphs() {

	// Parse the values from the CSV file
	firstValues, secondValues, err := prs.ParseResult()
	if err != nil {
		log.Fatal(err)
	}

	// Create a new plot
	p := plot.New()

	// Set the title and labels for the plot
	p.Title.Text = "Graph of Trends"
	p.X.Label.Text = "days"
	p.Y.Label.Text = "coast"

	// Add the first values to the plot as a scatter plot
	firstData := make(plotter.XYs, len(firstValues))
	for i, y := range firstValues {
		firstData[i].X = float64(i)
		firstData[i].Y = y
	}

	firstScatter, err := plotter.NewLine(firstData)
	if err != nil {
		log.Fatal(err)
	}

	firstScatter.LineStyle.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255} // Red
	p.Add(firstScatter)

	// Save the plot to a PNG file
	// if err := p.Save(4*vg.Inch, 4*vg.Inch, "plot.png"); err != nil {
	// 	panic(err)
	// }

	// Add the second values to the plot as a line plot
	secondData := make(plotter.XYs, len(secondValues))
	for i, y := range secondValues {
		secondData[i].X = float64(i)
		secondData[i].Y = y
	}
	secondLine, err := plotter.NewLine(secondData)
	if err != nil {
		log.Fatal(err)
	}
	secondLine.LineStyle.Color = color.RGBA{R: 0, G: 0, B: 255, A: 255} // Blue
	p.Add(secondLine)

	// Save the plot to a PNG file

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "/home/mallbruss/university/programm/grapsResult/plot.png"); err != nil {
		panic(err)
	}

}
