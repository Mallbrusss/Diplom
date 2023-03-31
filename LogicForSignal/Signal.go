package logicforsignal

import (
	"fmt"
	// mu "programm/CsvPackges/MemberShipDegree"
	ma "programm/Ma/MaFirst"
	rsi "programm/RsiModel"
)

// const (
// 	oversoldThreshold   = 30
// 	overboughtThreshold = 70
// )

// func determineMovingAverageSignal() string {
// 	shortMA := ma.ShortMovingAverage()
// 	longMA := ma.LongMovingAverage()
// 	if shortMA > longMA {
// 		return "buy"
// 	} else if shortMA < longMA {
// 		return "sell"
// 	} else {
// 		return "wait"
// 	}
// }

// func rsiSignal() string {

// 	currentRsi := rsi.RsiMain()

// 	if currentRsi <= oversoldThreshold {
// 		return "buy"
// 	} else if currentRsi >= overboughtThreshold {
// 		return "sell"
// 	} else {
// 		return "wait"
// 	}

// }

func BetterSignal() string {
	rsiValue := rsi.RsiMain()  // assume this function returns the current RSI value
	shortMA := ma.ShortMovingAverage() // assume this function returns the current short-term moving average
	longMA := ma.LongMovingAverage()   // assume this function returns the current long-term moving average

	rsiSignal := ""
	if rsiValue >= 70 {
		rsiSignal = "strong sell"
	} else if rsiValue >= 30 && rsiValue < 70 {
		rsiSignal = "neutral"
	} else {
		rsiSignal = "strong buy"
	}

	maSignal := ""
	diff := shortMA - longMA
	if diff >= 3 {
		maSignal = "strong buy"
	} else if diff > 0 && diff < 3 {
		maSignal = "moderate buy"
	} else if diff > -3 && diff < 0 {
		maSignal = "moderate sell"
	} else {
		maSignal = "strong sell"
	}

	switch {
	case rsiSignal == "strong sell" && maSignal == "strong sell":
		fmt.Println("sell")
		return "sell"
	case rsiSignal == "strong buy" && maSignal == "strong buy":
		fmt.Println("buy")
		return "buy"
	case rsiSignal == "strong sell" && maSignal == "moderate buy":
		fmt.Println("moderate sell")
		return "moderate sell"
	case rsiSignal == "moderate sell" && (maSignal == "strong buy" || maSignal == "moderated buy"):
		fmt.Println("moderate sell")
		return "moderate sell"
	case rsiSignal == "strong buy" && maSignal == "moderate sell":
		fmt.Println("moderate buy")
		return "moderate buy"
	case rsiSignal == "moderate buy" && (maSignal == "strong sell" || maSignal == "moderated  buy"):
		fmt.Println("moderate buy")
		return "moderate buy"
	case (rsiSignal == "neutral" && maSignal == "strong buy") || (rsiSignal == "strong buy" && maSignal == "neutral"):
		fmt.Println("moderate buy")
		return "moderate buy"
	case (rsiSignal == "neutral" && maSignal == "strong sell") || (rsiSignal == "strong sell" && maSignal == "neutral"):
		fmt.Println("moderate sell")
		return "moderate sell"
	default:
		fmt.Println("wait")
		return "wait"
	}

}

// func MainSignal() float64 {

// 	if betterSignal() == "sell" {
// 		fmt.Println("sell")
// 		return float64(-1)
// 	} else if betterSignal() == "buy" {
// 		fmt.Println("buy")
// 		return float64(1)
// 	} else if betterSignal() == "wait" {
// 		fmt.Println("wait")
// 		return float64(0)
// 	} else if betterSignal() == "moderate sell" {
// 		fmt.Println("moderate sell")
// 		return float64(-0.5)
// 	} else if betterSignal() == "moderate buy" {
// 		fmt.Println("moderate buy")
// 		return float64(0.5)
// 	} else {
// 		return float64(0)
// 	}
// }

// func SignalWithMemberShipDegree() {

// 	sigFirst := mu.MembershipDegreeShort() * MainSignal()
// 	sigSecond := mu.MembershipDegreeLong() * MainSignal()
// 	fmt.Println("Signal for short: ", sigFirst, " Signal for long: ", sigSecond)
// }
