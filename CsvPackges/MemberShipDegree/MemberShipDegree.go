package membersheipdegree

import (
	// "fmt"

	ma "programm/Ma/MaFirst"
)

const (
	lowerBound  = 15.0
	centerPoint = 30.0
	upperBound  = 45.0
)

func calculateMembershipDegree(maValue float64) float64 {
	if maValue < lowerBound || maValue > upperBound {
		return 0.0
	}

	if maValue == centerPoint {
		return 1.0
	}

	if maValue > centerPoint {
		return (upperBound - maValue) / (upperBound - centerPoint)
	}

	return (maValue - lowerBound) / (centerPoint - lowerBound)
}

func MembershipDegreeShort() float64 {
	mu := calculateMembershipDegree(ma.ShortMovingAverage())
	// fmt.Println("Membership degree short: ", mu)
	return mu
}

func MembershipDegreeLong() float64 {
	mu := calculateMembershipDegree(ma.LongMovingAverage())
	// fmt.Println("Membership degree long: ", mu)
	return mu
}

func Both() {
	MembershipDegreeShort()
	MembershipDegreeLong()
}
