package logicforsignal

import (
	"fmt"
	mu "programm/CsvPackges/MemberShipDegree"
	ma "programm/Ma/MaFirst"
	rsi "programm/RsiModel"
)

func Signal() string {
	rsiValue := rsi.RsiMain()          
	shortMA := ma.ShortMovingAverage() 
	longMA := ma.LongMovingAverage()   

	trendDirection := ""
	if rsiValue >= 70 {
		trendDirection = "strong down"
	} else if rsiValue >= 55 && rsiValue < 70 {
		trendDirection = "moderate down"
	} else if rsiValue >= 45 && rsiValue < 55 {
		trendDirection = "neutral"
	} else if rsiValue >= 30 && rsiValue < 45 {
		trendDirection = "moderate up"
	} else {
		trendDirection = "strong up"
	}
	
	maSignal := ""
	diff := shortMA - longMA
	if diff >= 3 {
		maSignal = "strong up"
	} else if diff > 0 && diff < 3 {
		maSignal = "moderate up"
	} else if diff > -3 && diff < 0 {
		maSignal = "moderate down"
	} else {
		maSignal = "strong down"
	}

	switch {
	case trendDirection == "strong down":
		switch maSignal {
		case "strong down":
			return "strong down"
		case "moderate down", "moderate up":
			return "moderate down"
		default:
			return "neutral"
		}

	case trendDirection == "moderate down":
		switch maSignal {
		case "strong up":
			return "moderate up"
		case "moderate down", "strong down":
			return "moderate down"
		default:
			return "neutral"
		}

	case trendDirection == "strong up":
		switch maSignal {
		case "strong up":
			return "strong up"
		case "moderate down", "moderate up":
			return "moderate up"
		default:
			return "neutral"
		}

	case trendDirection == "moderate up":
		switch maSignal {
		case "strong up", "moderate up":
			return "moderate up"
		case "strong down":
			return "moderate down"
		default:
			return "neutral"
		}

	default:
		return "neutral"
	}
}

func MainSignal() float64 {
	switch Signal() {
	case "strong down":
		return -1
	case "moderate down":
		return -0.5
	case "moderate up":
		return 0.5
	case "strong up":
		return 1
	default:
		return 0
	}
}

func SignalWithMemberShipDegree() {

	sigFirst := mu.MembershipDegreeShort() * MainSignal()
	sigSecond := mu.MembershipDegreeLong() * MainSignal()
	fmt.Println("trend line: ", MainSignal(), "\n Signal for short: ", sigFirst, "\n Signal for long: ", sigSecond)
}

func SignalWithMemberShipDegreeShort()float64{
	sig:=mu.MembershipDegreeShort()*MainSignal()
	return sig
}
func SignalWithMemberShipDegreeLong()float64{
	sig:=mu.MembershipDegreeLong()*MainSignal()
	return sig
}

func SignalForBuyOrSail() string{
	switch Signal() {
	case "strong down":
		return "buy" 
	case "moderate down":
		return "moderate buy" 
	case "moderate up":
		return "moderate sell" 
	case "strong up":
		return "sell" 
	default:
		return "wait" 
	}
}