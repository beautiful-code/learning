package typeconversions

import (
	"fmt"
	"math"
	"strconv"
)

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func stringToInt(str string) int {
	// convert to int
	res, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return res

}

func stringToFloat(str string) float64 {
	// convert to float with 4 digits of precision
	s, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic(err)
	}
	s = roundFloat(s, 4)
	return s
}

func FloatToString(value float64) string {
	// convert to float with 2 digits of precision
	value = roundFloat(value, 2)
	s := fmt.Sprintf("%v", value)
	fmt.Println( value)
	return s
}

func TestStringConverstions() {
	isFailed := false
	if stringToInt("10") != 10 {
		fmt.Println("Failed: stringToInt")
		isFailed = true
	}

	if stringToFloat("123.33333333333") != 123.3333 {
		fmt.Println("Failed: stringToFloat")
		isFailed = true
	}

	if FloatToString(0.3333) != "0.33" {
		fmt.Println("Failed: FloatToString")
		isFailed = true
	}

	if !isFailed {
		fmt.Println("All tests passed")
	}
}
