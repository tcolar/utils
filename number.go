// History: May 02 14 tcolar Creation

package utils

import (
	"math"
	"strconv"
)

// StringToFloat64 transforms a string to a float
// As a special case an empty string returns 0.0
func StrToFloat64(str string) (val float64, err error) {
	if str == "" {
		str = "0"
	}
	val, err = strconv.ParseFloat(str, 64)
	return val, err
}

// RoundFloat rounds a float into another float with the given number of decimals(precision).
func RoundFloat(val float64, decimals int) (rounded float64) {
	var rounder float64
	pow := math.Pow(10, float64(decimals))
	intermed := val * pow
	_, frac := math.Modf(intermed)
	if math.Abs(frac) >= 0.5 {
		if val > 0 {
			rounder = math.Ceil(intermed)
		} else {
			rounder = math.Floor(intermed)
		}
	} else {
		if val > 0 {
			rounder = math.Floor(intermed)
		} else {
			rounder = math.Ceil(intermed)
		}
	}
	return rounder / pow
}
