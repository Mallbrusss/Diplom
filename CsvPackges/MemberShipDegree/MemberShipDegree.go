package membersheipdegree

import (
	// "fmt"

	ma "programm/Ma/MaFirst"
)

const (
	lowerBound  = 15
	centerPoint = 30
	upperBound  = 45
)

func calculateMembershipDegree(maValue float64) float64 {
	if maValue <= lowerBound || maValue >= upperBound {
		return 0.0
	} else if maValue >= centerPoint && maValue < upperBound {
		return (upperBound - maValue) / (upperBound - centerPoint)
	} else {
		return (maValue - lowerBound) / (centerPoint - lowerBound)
	}
}

func MembershipDegreeShort() float64 {
	fastMA := ma.ShortMovingAverage()
	mu := calculateMembershipDegree(fastMA)
	// fmt.Println("Membership degree short: ", mu)
	return mu
}

func MembershipDegreeLong() float64 {
	longMA := ma.LongMovingAverage()
	mu := calculateMembershipDegree(longMA)
	// fmt.Println("Membership degree long: ", mu)
	return mu
}

func Both() {
	MembershipDegreeShort()
	MembershipDegreeLong()
}
