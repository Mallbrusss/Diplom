package main

import (
	"fmt"

	ma "programm/Ma/MaFirst"
	wrm "programm/Ma/WriteMa"
	gr "programm/Graphs"
	mu "programm/CsvPackges/MemberSheipDegree"
	rsi "programm/RsiModel"
	sig "programm/LogicForSignal"
)

func main() {
	movingAvg := ma.MovingAverage()
	if err := wrm.WriteMovingAvgToCsv(movingAvg); err != nil {
		fmt.Println("Failed to write moving average to CSV:", err)
	}	
	gr.Graphs()
	mu.Membersheipdegree()
	rsi.RsiMain()
	sig.DetermineMovingAverageSignal()
}
