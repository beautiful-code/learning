package typeconversions

import "fmt"

func stringToInt(str string) int {
	// convert to int

	return 0
}

func stringToFloat(str string) float64 {
	// convert to float with 4 digits of precision

	return 0
}

func FloatToString(value float64) string {
	// convert to float with 2 digits of precision

	return "0"
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

	if FloatToString(1/3) != "0.33" {
		fmt.Println("Failed: stringToFloat")
		isFailed = true
	}

	if !isFailed {
		fmt.Println("All tests passed")
	}
}
