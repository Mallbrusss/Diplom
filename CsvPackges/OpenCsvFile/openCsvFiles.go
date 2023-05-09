package openCsv

import (
	"os"
	che "programm/ErrorsCheck"
)

func OpenCsvFile() *os.File {
	// Open the CSV file
	file, err := os.Open("/home/mallbruss/university/programm/data/testData/Severstal.csv")
	che.CheckError(err)
	return file
}
