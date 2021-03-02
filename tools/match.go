package tools

import (
	"math/big"
	"strconv"
)

func Float64Add(x float64, y float64, more ...float64) float64 {
	floatX := new(big.Float).SetFloat64(x)
	floatY := new(big.Float).SetFloat64(y)
	result := new(big.Float).Add(floatX, floatY)
	if len(more) > 0 {
		for _, m := range more {
			floatM := new(big.Float).SetFloat64(m)
			result = new(big.Float).Add(result, floatM)
		}
	}
	f, _ := strconv.ParseFloat(result.String(), 64) // 忽略 error
	return f
}

