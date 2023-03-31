package maFirst

import (
	
	prs "programm/CsvPackges/Parse"
)

func ShortMovingAverage() float64 {
	// Extract the last ten closing prices from the slice
	firstTen := prs.ParseCSVFirst()[:10]

	for i := 0; i < 10; i++ {
		firstTen[i] = firstTen[len(firstTen)-i-1]
	}

	// Calculate the moving average of the last ten closing prices
	var sum float64
	for _, price := range firstTen {
		sum += price
	}
	movingAvg := sum / float64(len(firstTen))
	//fmt.Println("Moving average:", movingAvg)
	return movingAvg
}

func LongMovingAverage() float64 {
	// Extract the last ten closing prices from the slice
	firstTwenty := prs.ParseCSVSecond()[:20]

	for i := 0; i < 20; i++ {
		firstTwenty[i] = firstTwenty[len(firstTwenty)-i-1]
	}

	// Calculate the moving average of the last ten closing prices
	var sum float64
	for _, price := range firstTwenty {
		sum += price
	}
	movingAvg := sum / float64(len(firstTwenty))
	// fmt.Println("Moving average Second:", movingAvg)
	return movingAvg
}

