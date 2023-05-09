package maFirst

import (
	prs "programm/CsvPackges/Parse"
)

func ShortMovingAverage() float64 {
	firstTwenty := prs.ParseCSVFirst()[:20]

	for i := 0; i < 20; i++ {
		firstTwenty[i] = firstTwenty[len(firstTwenty)-i-1]
	}
	var sum float64

	for _, price := range firstTwenty {
		sum += price
	}
	movingAvg := sum / float64(len(firstTwenty))
	return movingAvg
}

func LongMovingAverage() float64 {
	firstFifty := prs.ParseCSVFirst()[:50]

	for i := 0; i < 50; i++ {
		firstFifty[i] = firstFifty[len(firstFifty)-i-1]
	}

	var sum float64
	for _, price := range firstFifty {
		sum += price
	}
	movingAvg := sum / float64(len(firstFifty))
	return movingAvg
}
