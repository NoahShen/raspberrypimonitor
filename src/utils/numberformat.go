package utils

import (
	"strconv"
)

func FormatFloatToPercent(f float64) string {
	formated := strconv.FormatFloat(f*100, 'f', 2, 64)
	return formated
}
