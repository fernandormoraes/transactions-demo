package helpers

import "math"

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

const (
	LayoutDate       string = "2006-01-02"
	SixMonthsToHours int32  = 4380
)
