package main

import (
	"fmt"
	"math"
	"strconv"
)

func stringToInt(str string) int {
	// convert to int
	int,_ := strconv.Atoi(str)
	return int
}

func stringToFloat(str string) float64 {
	// convert to float with 4 digits of precision
	float, _ := strconv.ParseFloat(str,64)
	return math.Round(float*10000) / 10000
}

func FloatToString(value float64) string {
	// convert to float with 2 digits of precision
	return strconv.FormatFloat(value, 'f', 2, 64)
}

func main4(){
	isFailed := false
	if stringToInt("10") != 10 {
		fmt.Println("Failed: stringToInt")
		isFailed = true
	}

	if stringToFloat("123.33333333333") != 123.3333 {
		fmt.Println("Failed: stringToFloat")
		isFailed = true
	}

	if FloatToString(1.0/3) != "0.33" {
		fmt.Println("Failed: FloatToString")
		isFailed = true
	}

	if !isFailed {
		fmt.Println("All tests passed")
	}
}