package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64 = 2000.0
	var years float64
	expectedReturnRate := 5.5

	outPutText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	outPutText("expectedReturnRate: ")
	fmt.Scan(&expectedReturnRate)

	outPutText("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	//fmt.Println("future value", futureValue)
	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for Inflation): %.1f\n", futureRealValue)
	// formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	// formattedRFV := fmt.Sprintf("Future Value (adjusted for Inflation): %.1f\n", futureRealValue)

	// fmt.Printf(`Future Value: %.1f\n
	// Future Value (adjusted for Inflation): %.1f`, futureValue, futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}

func outPutText(text string) {
	fmt.Print(text)
}
