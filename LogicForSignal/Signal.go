package logicforsignal

import (
	"fmt"
	mu "programm/CsvPackges/MemberShipDegree"
	ma "programm/Ma/MaFirst"
	rsi "programm/RsiModel"
)


func BetterSignal() string {
	rsiValue := rsi.RsiMain()          // assume this function returns the current RSI value
	shortMA := ma.ShortMovingAverage() // assume this function returns the current short-term moving average
	longMA := ma.LongMovingAverage()   // assume this function returns the current long-term moving average

	rsiSignal := ""
	if rsiValue >= 70 {
		rsiSignal = "strong sell"
	} else if rsiValue >= 55 && rsiValue < 70 {
		rsiSignal = "moderate sell"
	} else if rsiValue >= 45 && rsiValue < 55 {
		rsiSignal = "neutral"
	} else if rsiValue >= 30 && rsiValue < 45 {
		rsiSignal = "moderate buy"
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
	// for strong sell

	case rsiSignal == "strong sell" && maSignal == "strong sell":

		return "sell"

	case rsiSignal == "strong sell" && maSignal == "moderate sell":

		return "moderate sell"

	case rsiSignal == "strong sell" && maSignal == "moderate buy":

		return "moderate sell"

		// for moderate sell
	case rsiSignal == "moderate sell" && (maSignal == "strong buy" || maSignal == "moderated buy"):

		return "moderate buy"

	case rsiSignal == "moderate sell" && (maSignal == "moderate sell" || maSignal == "strong sell"):

		return "moderate sell"

		// for strong buy

	case rsiSignal == "strong buy" && maSignal == "strong buy":

		return "buy"

	case rsiSignal == "strong buy" && maSignal == "moderate sell":

		return "moderate buy"

	case rsiSignal == "strong buy" && maSignal == "moderate buy":

		return "moderate buy"

		// for moderate buy

	case rsiSignal == "moderate buy" && maSignal == "moderate buy":

		return "moderate buy"
	case rsiSignal == "moderate buy" && maSignal == "strong sell":

		return "moderate sell"
	case rsiSignal == "moderate buy" && maSignal == "strong buy":

		return "moderate buy"

	default:
		fmt.Println("wait")
		return "wait"
	}

}

func MainSignal() float64 {

	if BetterSignal() == "sell" {
		return float64(-1)
	} else if BetterSignal() == "buy" {
		return float64(1)
	} else if BetterSignal() == "wait" {
		return float64(0)
	} else if BetterSignal() == "moderate sell" {
		return float64(-0.5)
	} else if BetterSignal() == "moderate buy" {
		return float64(0.5)
	} else {
		return 0
	}
}

func SignalWithMemberShipDegree() {

	sigFirst := mu.MembershipDegreeShort() * MainSignal()
	sigSecond := mu.MembershipDegreeLong() * MainSignal()
	fmt.Println("Signal for short: ", sigFirst, " Signal for long: ", sigSecond)
}
