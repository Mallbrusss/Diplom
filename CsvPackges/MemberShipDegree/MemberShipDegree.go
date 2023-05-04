package membersheipdegree

import (
	ma "programm/Ma/MaFirst"
)

const (
	lowerBound  = 60.0
	centerPoint = 75.0    
	upperBound  = 90.0
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
