package calculator

import "math"

const DefaultInterest float64 = 0.01

var lowerInterestLimits map[int]float64 = map[int]float64{
	1000:  0.015,
	5000:  0.02,
	10000: 0.025,
	50000: 0.03,
}

func CalculateInterest(balance float64) float64 {
	if balance <= 0 {
		return 0
	}

	// round to nearest even number because financial calculations are like that.
	interest := math.RoundToEven(balance*DefaultInterest*100.0) / 100.0

	return interest
}
