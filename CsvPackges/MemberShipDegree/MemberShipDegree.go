package membersheipdegree

import (
	"fmt"
	"math"
	prs "programm/CsvPackges/Parse"
	ma "programm/Ma/MaFirst"
	"sort"
)

var (
	lowerBound  = CalculateSupport()
	centerPoint = CalculateMedianPrice()
	upperBound  = CalculateResistance()
)

// const (
// 	lowerBound = 750
// 	centerPoint =800
// 	upperBound = 950

// )
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
	fmt.Println("short mbsd", mu)
	return math.Abs(mu)
}

func MembershipDegreeLong() float64 {
	mu := calculateMembershipDegree(ma.LongMovingAverage())
	fmt.Println("long mbsd", mu)
	return math.Abs(mu)
}

func Both() {
	MembershipDegreeShort()
	MembershipDegreeLong()
}

func getPrices() []float64 {
	prices := prs.ParseCSVFirst()
	filteredPrices := make([]float64, len(prices))
	copy(filteredPrices, prices)
	sort.Float64s(filteredPrices)
	return filteredPrices
}

func CalculateSupport() float64 {
	filteredPrices := getPrices()
	minPrice := filteredPrices[0]
	maxPrice := filteredPrices[len(filteredPrices)-1]
	support := minPrice + 0.3*(maxPrice-minPrice)
	fmt.Println("support: ", support)
	return support
}
func CalculateResistance() float64 {
	filteredPrices := getPrices()
	minPrice := filteredPrices[0]
	maxPrice := filteredPrices[len(filteredPrices)-1]
	resistance := maxPrice - 0.3*(maxPrice-minPrice)
	fmt.Println("resistance: ", resistance)
	return resistance
}
func CalculateMedianPrice() float64 {
	filteredPrices := getPrices()
	var medianPrice float64
	if len(filteredPrices)%2 == 0 {
		medianPrice = (filteredPrices[len(filteredPrices)/2-1] + filteredPrices[len(filteredPrices)/2]) / 2
	} else {
		medianPrice = filteredPrices[len(filteredPrices)/2]
	}
	fmt.Println("Median: ", medianPrice)
	return medianPrice
}
