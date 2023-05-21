package openCsv

import (
	"encoding/csv"
	"strconv"

	opf "programm/CsvPackges/OpenCsvFile"
)

func ParseCSVFirst() []float64 {
	file := opf.OpenCsvFile()
	defer file.Close()
	// Parse the CSV file into a slice of closing prices
	reader := csv.NewReader(file)
	var closingPrices []float64
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		// Extract the closing price from the "Close" column
		closePrice, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			continue
		}
		closingPrices = append(closingPrices, closePrice)
	}
	return closingPrices
}

func ParseCSVEndClose() float64 {
	file := opf.OpenCsvFile()
	defer file.Close()

	reader := csv.NewReader(file)

	// Skip first line (header)
	if _, err := reader.Read(); err != nil {
		return 0
	}

	record, err := reader.Read()
	if err != nil {
		return 0.0
	}

	closePrice, err := strconv.ParseFloat(record[1], 64)
	if err != nil {
		return 0.0
	}

	return closePrice
}
