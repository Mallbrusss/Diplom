package openCsv

import (
	"encoding/csv"
	"io"
	"strconv"

	opf "programm/CsvPackges/OpenCsvFile"
)

func ParseCSVFirst() []float64 {
	file := opf.OpenCsvFile()
	// Parse the CSV file into a slice of closing prices
	reader := csv.NewReader(file)
	var closingPrices []float64
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		// Extract the closing price from the "Close" column
		closePrice, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			continue
		}
		closingPrices = append(closingPrices, closePrice)
	}
	return closingPrices
}

func ParseCSVSecond() []float64 {
	file := opf.OpenCsvFile()
	// Parse the CSV file into a slice of closing prices
	reader := csv.NewReader(file)
	var closingPrices []float64
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		// Extract the closing price from the "Close" column
		closePrice, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			continue
		}
		closingPrices = append(closingPrices, closePrice)
	}
	return closingPrices
}

func ParseResult() ([]float64, []float64, error) {
	file := opf.OpenCsvFileResults()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	var firstValues []float64
	var secondValues []float64

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, nil, err
		}

		firstValue, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			continue
		}
		secondValue, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			continue
		}

		firstValues = append(firstValues, firstValue)
		secondValues = append(secondValues, secondValue)
	}

	return firstValues, secondValues, nil
}

// func ParseDate() ([]float64, error) {
// 	file := opf.OpenCsvFile()
// 	// Parse the CSV file into a slice of closing prices
// 	reader := csv.NewReader(file)
// 	reader.FieldsPerRecord = -1

// 	var dates []float64
// 	for {
// 		record, err := reader.Read()
// 		if err == io.EOF {
// 			break
// 		} else if err != nil {
// 			fmt.Println("Error reading record:", err)
// 			continue
// 		}

// 		dateStr := record[0]
// 		date, err := time.Parse("01/02/2006", dateStr)
// 		if err != nil {
// 			return nil, fmt.Errorf("error parsing date: %w", err)
// 		}
// 		date = append(dates, float64(date.Unix()))
// 		//fmt.Println("Date:", date)
// 	}
// 	return dates, nil
// }

// func ParseIndex()[]float64{
// 	file := opf.OpenCsvFileResults()

// 	reader := csv.NewReader(file)

// 	reader.FieldsPerRecord = -1

// }
