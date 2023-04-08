package rsi

import (
	//"fmt"
	"math"

	prs "programm/CsvPackges/Parse"
	err "programm/ErrorsCheck"
)

func rsi(prices []float64, period int) float64 {
	if len(prices) < period {
		err.CheckError(nil)
	}

	changes := make([]float64, len(prices)-1)
	for i := range changes {
		changes[i] = prices[i+1] - prices[i]
	}

	gains := make([]float64, len(changes))
	losses := make([]float64, len(changes))
	for i, change := range changes {
		if change >= 0 {
			gains[i] = change
		} else {
			losses[i] = math.Abs(change)
		}
	}

	avgGain := average(gains[:period])
	avgLoss := average(losses[:period])
	rs := avgGain / avgLoss
	var rsi float64
	if avgLoss == 0 {
		if avgGain == 0 {
			rsi = 50
		} else {
			rsi = 100
		}
	} else {
		rsi = 100 - (100 / (1 + rs))
	}
	// fmt.Println("rsi is: ", rsi)
	return rsi
}

func average(numbers []float64) float64 {
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func RsiMain() float64 {
	prices := prs.ParseCSVFirst()
	period := 10
	return rsi(prices, period)
}
