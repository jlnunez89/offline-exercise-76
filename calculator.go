package calculator

import "math"

const DefaultInterest float64 = 0.01

func CalculateInterest(balance float64) float64 {
	if balance <= 0 {
		return 0
	}

	interestRate := DefaultInterest

	if balance >= 1000.00 {
		interestRate = 0.015
	}

	balanceInCents := balance * interestRate * 100.0

	// round to nearest even number because financial calculations are like that.
	roundedInterest := math.RoundToEven(balanceInCents) / 100.0

	return roundedInterest
}
