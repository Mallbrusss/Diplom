package membersheipdegree

import (
	"fmt"
	ma "programm/Ma/MaFirst"
)

func Membersheipdegree() float64 {
	var fastMA float64 = ma.MovingAverage() // Fast Moving Average value from trading data
	var a float64 = 20                      // Lower bound for "low" membership function
	var b float64 = 30                      // Center point for "medium" membership function
	var c float64 = 40                      // Upper bound for "high" membership function

	var mu float64 = 0.0
	if fastMA <= a || fastMA >= c {
		mu = 0.0
	} else if fastMA >= b {
		mu = (c - fastMA) / (c - b)
	} else {
		mu = (fastMA - a) / (b - a)
	}

	fmt.Println("Membership degree:", mu)
	return mu
}
	