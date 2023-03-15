package main

import (
	"fmt"

	//che "programm/ErrorsCheck"
	//opf "programm/openCsvFile"
	//prs "programm/Parse"
	ma "programm/Ma/MaFirst"
	wrm "programm/Ma/WriteMa"
	gr "programm/Graphs"
)

func main() {
	movingAvg := ma.MovingAverage()
	movingAvgSec := ma.MovingAverageSecond()
	if err := wrm.WriteFirstMovingAvgToCsv(movingAvg); err != nil {
		fmt.Println("Failed to write moving average to CSV:", err)
	}
	if err := wrm.WriteFirstMovingAvgToCsv(movingAvgSec); err != nil {
		fmt.Println("Failed to write moving average to CSV:", err)
	}
	
	gr.Graphs()
}
