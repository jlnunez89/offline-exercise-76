package calculator

import "math"

const DefaultInterest float64 = 0.01

func CalculateInterest(balance float64) float64 {
	if balance <= 0 {
		return 0
	}

	balanceInCents := balance * 100.0
	interestRate := DefaultInterest

	if balance >= 50000.00 {
		interestRate = 0.03
	} else if balance >= 10000.00 {
		interestRate = 0.025
	} else if balance >= 5000.00 {
		interestRate = 0.02
	} else if balance >= 1000.00 {
		interestRate = 0.015
	}

	// round to nearest even number because financial calculations are like that.
	roundedInterest := math.RoundToEven(balanceInCents*interestRate) / 100.0

	return roundedInterest
}
