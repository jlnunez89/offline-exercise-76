package calculator

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCalculateInterest(t *testing.T) {
	t.Run("Negative balance should return 0", func(t *testing.T) {
		// arrange
		negativeBalance := -1.0
		expectedInterest := 0.0

		// act
		interest := CalculateInterest(negativeBalance)

		// assert
		assert.Equal(t, expectedInterest, interest)
	})

	t.Run("Zero balance should return 0", func(t *testing.T) {
		// arrange
		zeroBalance := 0.0
		expectedInterest := 0.0

		// act
		interest := CalculateInterest(zeroBalance)

		// assert
		assert.Equal(t, expectedInterest, interest)
	})

	t.Run("Test balances between 0 and 1000", func(t *testing.T) {
		// arrange
		epsilon := 0.01 // since one cent is normally the smallest unit of currency
		nonZeroBalance := 0.0 + epsilon
		lessThan1000Balance := 1000.0 - epsilon
		expectedInterestPercentage := DefaultInterest

		for balance := nonZeroBalance; balance < lessThan1000Balance; balance += epsilon {
			// act
			calculatedInterest := CalculateInterest(balance)
			expectedInterest := math.RoundToEven(balance*expectedInterestPercentage*100.0) / 100.0 // round to nearest even number because finance

			// assert
			require.Equal(t, expectedInterest, calculatedInterest)
		}
	})

	t.Run("Test balances between 1000 and 5000", func(t *testing.T) {
		// arrange
		epsilon := 0.01 // since one cent is normally the smallest unit of currency
		lowerLimitBalance := 1000.0
		upperLimitBalance := 5000.0 - epsilon
		expectedInterestPercentage := 0.015

		for balance := lowerLimitBalance; balance < upperLimitBalance; balance += epsilon {
			// act
			calculatedInterest := CalculateInterest(balance)
			expectedInterest := math.RoundToEven(balance*expectedInterestPercentage*100.0) / 100.0 // round to nearest even number because finance

			// assert
			require.Equal(t, expectedInterest, calculatedInterest, "Expected calculated interest for %v to be %v, but got %v", balance, expectedInterest, calculatedInterest)
		}
	})

	t.Run("Test balances between 5000 and 10000", func(t *testing.T) {
		// arrange
		epsilon := 0.01 // since one cent is normally the smallest unit of currency
		lowerLimitBalance := 5000.0
		upperLimitBalance := 10000.0 - epsilon
		expectedInterestPercentage := 0.02

		for balance := lowerLimitBalance; balance < upperLimitBalance; balance += epsilon {
			// act
			calculatedInterest := CalculateInterest(balance)
			expectedInterest := math.RoundToEven(balance*expectedInterestPercentage*100.0) / 100.0 // round to nearest even number because finance

			// assert
			require.Equal(t, expectedInterest, calculatedInterest, "Expected calculated interest for %v to be %v, but got %v", balance, expectedInterest, calculatedInterest)
		}
	})

	t.Run("Test balances between 10000 and 50000", func(t *testing.T) {
		// arrange
		epsilon := 0.01 // since one cent is normally the smallest unit of currency
		lowerLimitBalance := 10000.0
		upperLimitBalance := 50000.0 - epsilon
		expectedInterestPercentage := 0.025

		for balance := lowerLimitBalance; balance < upperLimitBalance; balance += epsilon {
			// act
			calculatedInterest := CalculateInterest(balance)
			expectedInterest := math.RoundToEven(balance*expectedInterestPercentage*100.0) / 100.0 // round to nearest even number because finance

			// assert
			require.Equal(t, expectedInterest, calculatedInterest, "Expected calculated interest for %v to be %v, but got %v", balance, expectedInterest, calculatedInterest)
		}
	})

	t.Run("Test balances above 50000", func(t *testing.T) {
		// arrange
		epsilon := 0.01 // since one cent is normally the smallest unit of currency
		lowerLimitBalance := 50000.0
		upperLimitBalance := 100000.0 - epsilon
		expectedInterestPercentage := 0.03

		for balance := lowerLimitBalance; balance < upperLimitBalance; balance += epsilon {
			// act
			calculatedInterest := CalculateInterest(balance)
			expectedInterest := math.RoundToEven(balance*expectedInterestPercentage*100.0) / 100.0 // round to nearest even number because finance

			// assert
			require.Equal(t, expectedInterest, calculatedInterest, "Expected calculated interest for %v to be %v, but got %v", balance, expectedInterest, calculatedInterest)
		}

		// sporadically test a few values above 100,000:
		discreetBalances := []float64{150000.0, 300000.0, 500000.0, 1000000.0, 2000000.0, 10000000.00}
		for _, balance := range discreetBalances {
			// act
			calculatedInterest := CalculateInterest(balance)
			expectedInterest := math.RoundToEven(balance*expectedInterestPercentage*100.0) / 100.0 // round to nearest even number because finance

			// assert
			require.Equal(t, expectedInterest, calculatedInterest, "Expected calculated interest for %v to be %v, but got %v", balance, expectedInterest, calculatedInterest)
		}
	})
}
