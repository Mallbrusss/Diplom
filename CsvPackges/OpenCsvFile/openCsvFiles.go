package openCsv

import (
	"os"
	che "programm/ErrorsCheck"
)

func OpenCsvFile() *os.File {
	// Open the CSV file
	file, err := os.Open("/home/mallbruss/university/programm/data/testData/Brent.csv")
	che.CheckError(err)
	return file
}

func OpenCsvFileResults() *os.File {
	// Open the CSV file
	file, err := os.Open("/home/mallbruss/university/programm/data/result/result1.csv")
	che.CheckError(err)
	return file
}
