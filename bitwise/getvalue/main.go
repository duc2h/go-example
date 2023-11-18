package main

import (
	"fmt"
	"log"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common/math"
)

// how can we get value from a bit slot, then visualize it.
// This can be not a best way, i will improve it later.
// isPositive -> amount -> price: 1 -> 32 -> 160
// amount = 200
// price = 100000000 per amount.

var (
	amount_global = int64(600000)
	price_global  = int64(100000000)
)

func main() {
	fee0 := makeABitValuePositiveFalse()

	fmt.Println(fee0)

	a := new(big.Int).Rsh(fee0, 160)
	fmt.Println(a)

	c := BigToUint160(fee0)

	cc := new(big.Int).SetBytes(c[:])
	fmt.Println(cc)

}

func makeABitValuePositiveFalse() *big.Int {

	a := new(big.Int).Lsh(big.NewInt(1), 32)

	amountBig := big.NewInt(100000)
	al := new(big.Int).Or(a, amountBig)
	priceBig := math.MustParseBig256("0x4f82e73edb06d29ff62c91ec8f5ff06571bdeb29")
	fmt.Println("len 1: ", len(al.Bytes()))
	// add 160 bits slot for price. amountBig << 160
	amountBigLefted := new(big.Int).Lsh(al, 160)

	fmt.Println("len 2: ", len(amountBigLefted.Bytes()))
	// merge amountBigLefted and priceBig.
	fee := new(big.Int).Or(amountBigLefted, priceBig)

	return fee
}

func getValueFromBit(fee *big.Int) (int64, uint32, *big.Int) { // isPositive, amount, price
	// get price's value. uint160(fee)
	ui160Bytes := BigToUint160(fee)
	price := new(big.Int).SetBytes(ui160Bytes[:])

	// get amount's value. uint32(fee>>160)
	priceRighted := new(big.Int).Rsh(fee, 160)
	amount := getAmount(priceRighted)

	// get isPositive's value. fee>>192
	isPositive := new(big.Int).Rsh(fee, 192).Int64()

	return isPositive, amount, price
}

func getValueFromBit1(fee *big.Int) { // isPositive, amount, price
	// get price's value. uint160(fee)
	// ui160Bytes := BigToUint160(fee)
	// price := new(big.Int).SetBytes(ui160Bytes[:])

	// get amount's value. uint32(fee>>160)
	priceRighted := new(big.Int).Rsh(fee, 160)
	fmt.Println(priceRighted.String())

	return
}

func makeABitValuePositiveTrue() *big.Int {
	positiveBig := big.NewInt(1)
	amountBig := big.NewInt(amount_global)
	priceBig := big.NewInt(price_global)

	// add 32 bit slots. positiveBig << 32
	positiveLefted := new(big.Int).Lsh(positiveBig, 32) // => 33 bits

	// merge positiveLefted and amountBig.
	amountMerged := new(big.Int).Or(positiveLefted, amountBig) // 33 bits

	// add more 160 bit slots. amountMerged << 160
	amountLefted := new(big.Int).Lsh(amountMerged, 160) // 193 bits

	// merge amountLefted and priceBig.

	fee := new(big.Int).Or(amountLefted, priceBig) // 193 bits.

	return fee
}

func getAmount(amountBig *big.Int) uint32 {
	// convert big.int to uint64
	ui64, err := strconv.ParseInt(amountBig.String(), 10, 64) // there is can be 33 bits slot.
	if err != nil {
		log.Fatal(err)
	}

	// convert uint64 to uint32
	return uint32(ui64)
}

func BigToUint160(n *big.Int) uint160 {
	var result uint160
	bytes := n.Bytes()

	fmt.Println("len bytes: ", len(bytes))
	copy(result[:], bytes[max(0, len(bytes)-20):])
	return result
}

// https://docs.soliditylang.org/en/v0.8.20/types.html#address
type uint160 [20]byte

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func getValueFromBitPositive0(fee string) uint32 {

// }
