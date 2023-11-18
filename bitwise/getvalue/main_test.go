package main

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/stretchr/testify/assert"
)

func TestGetValueFromBitPositiveFalse(t *testing.T) {
	fee := makeABitValuePositiveFalse()

	// a := "146159165624476364475945418325124367680971254365544"
	// fee, _ := new(big.Int).SetString(a, 0)
	// fmt.Println(len(fee.Bytes()))

	isPositive, amount, price := getValueFromBit(fee)

	fmt.Println(amount)
	// fmt.Println(price.String())
	// getValueFromBit1(fee)
	b, _ := math.ParseBig256("0x0193a8a52d77e27bdd4f12e0cdd52d8ff1d97d68")

	c := hexutil.EncodeBig(price)
	fmt.Println(c)

	fmt.Println(b.String())

	assert.Equal(t, int64(0), isPositive)
	assert.Equal(t, uint32(amount_global), amount)
	assert.Equal(t, price_global, price.Int64())
}

func TestGetValueFromBitPositiveTrue(t *testing.T) {
	fee := makeABitValuePositiveTrue()
	fmt.Println(fee.String())

	isPositive, amount, price := getValueFromBit(fee)

	assert.Equal(t, int64(0), isPositive)
	assert.Equal(t, uint32(amount_global), amount)
	assert.Equal(t, price_global, price.Int64())
}
