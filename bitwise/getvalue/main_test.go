package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetValueFromBitPositiveFalse(t *testing.T) {
	fee := makeABitValuePositiveFalse()

	isPositive, amount, price := getValueFromBit(fee)

	assert.Equal(t, int64(0), isPositive)
	assert.Equal(t, uint32(amount_global), amount)
	assert.Equal(t, price_global, price.Int64())
}

func TestGetValueFromBitPositiveTrue(t *testing.T) {
	fee := makeABitValuePositiveTrue()

	isPositive, amount, price := getValueFromBit(fee)

	assert.Equal(t, int64(1), isPositive)
	assert.Equal(t, uint32(amount_global), amount)
	assert.Equal(t, price_global, price.Int64())
}
