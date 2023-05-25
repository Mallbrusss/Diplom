package openCsv

import (
	"os"
	che "programm/ErrorsCheck"
)


const ( 
	yand = "/home/mallbruss/university/programm/data/testData/Yandex.csv"
	acron = "/home/mallbruss/university/programm/data/testData/Acron.csv"
)


func OpenCsvFile() *os.File {
	// Open the CSV file
	file, err := os.Open(acron)
	che.CheckError(err)
	return file
}
