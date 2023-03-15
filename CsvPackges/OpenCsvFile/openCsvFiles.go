package openCsv

import (
	"os"
	che "programm/ErrorsCheck"
)

func OpenCsvFile() *os.File {
	// Open the CSV file
	file, err := os.Open("/home/mallbruss/university/programm/data/testData/Download Data - STOCK_US_XNYS_CSV.csv")
	che.CheckError(err)
	return file
}

func OpenCsvFileResults() *os.File {
	// Open the CSV file
	file, err := os.Open("/home/mallbruss/university/programm/data/result/result1.csv")
	che.CheckError(err)
	return file
}
