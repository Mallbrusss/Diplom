package membersheipdegree

import (
	// "fmt"

	ma "programm/Ma/MaFirst"
)

const (
	lowerBound  = 650
	centerPoint = 850
	upperBound  = 1000
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
	return mu
}

func MembershipDegreeLong() float64 {
	mu := calculateMembershipDegree(ma.LongMovingAverage())
	return mu
}

func Both() {
	MembershipDegreeShort()
	MembershipDegreeLong()
}
