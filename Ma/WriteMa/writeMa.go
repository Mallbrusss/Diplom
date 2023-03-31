package writema

import (
	"encoding/csv"
	"fmt"
	"os"

	ma "programm/Ma/MaFirst"
)

func WriteMovingAvgToCsv(float64) error {

	file, err := os.OpenFile("/home/mallbruss/university/programm/data/result/result1.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	// Write the moving averages to the CSV file
	writer := csv.NewWriter(file)
	err = writer.Write([]string{fmt.Sprintf("%.2f", ma.ShortMovingAverage()), fmt.Sprintf("%.2f", ma.LongMovingAverage())})
	if err != nil {
		return err
	}
	writer.Flush()

	fmt.Println("Moving averages written to result.csv")

	return nil

}
