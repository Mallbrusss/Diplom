package main

import (
	"fmt"

	ma "programm/Ma/MaFirst"
	wrm "programm/Ma/WriteMa"
	gr "programm/Graphs"
	sig "programm/LogicForSignal"

)

func main() {
	movingAvg := ma.ShortMovingAverage()
	if err := wrm.WriteMovingAvgToCsv(movingAvg); err != nil {
		fmt.Println("Failed to write moving average to CSV:", err)
	}	
	 gr.Graphs()
	sig.BetterSignal()
}
