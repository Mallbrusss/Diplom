package rsi

import (
	"fmt"
	"math"

	che "programm/ErrorsCheck"
	prs "programm/CsvPackges/Parse"
)

// TODO return float64 not []float64
func rsi(prices []float64, period int) (float64) {
	if len(prices) < period {
		che.CheckError(nil)
	}

	changes := make([]float64, len(prices)-1)
	for i := 0; i < len(changes); i++ {
		changes[i] = prices[i+1] - prices[i]
	}

	gains := make([]float64, len(changes))
	losses := make([]float64, len(changes))
	for i := 0; i < len(changes); i++ {
		if changes[i] >= 0 {
			gains[i] = changes[i]
		} else {
			losses[i] = math.Abs(changes[i])
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

	return rsi
}

func average(numbers []float64) float64 {
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}


func RsiMain()float64 {
	prices :=  prs.ParseCSVFirst()
	period := 14
	rsi := rsi(prices, period)
	fmt.Println("rsi is " , rsi)
	return rsi
}
