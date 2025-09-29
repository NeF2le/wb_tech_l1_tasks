package main

import (
	"fmt"
	"math/big"
)

type BigCalc struct{}

func (c *BigCalc) Sum(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func (c *BigCalc) Sub(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func (c *BigCalc) Mul(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func (c *BigCalc) Div(a, b *big.Int) *big.Int {
	return new(big.Int).Div(a, b)
}

func main() {
	bc := BigCalc{}

	a, _ := new(big.Int).SetString("123456789012345678901234567890", 10)
	b, _ := new(big.Int).SetString("123456789123456789001234506789", 10)

	fmt.Println(bc.Sum(a, b))
	fmt.Println(bc.Sub(a, b))
	fmt.Println(bc.Div(a, b))
	fmt.Println(bc.Mul(a, b))
}
