package main

import (
	"fmt"
	"math"
)

/*
*This function is used to compute the correlation coefficient r
 */
func main() {
	var x [11]float64
	var y [11]float64
	sumx := 0.0
	sumy := 0.0
	sumxy := 0.0
	sumx2 := 0.0
	sumy2 := 0.0
	var r float64

	fmt.Println("Enter the value of x and y")
	for i := 0; i < 11; i++ {
		fmt.Scanln(&x[i], &y[i])
		sumx = sumx + x[i]
		sumy = sumy + y[i]
		sumxy = sumxy + (x[i] * y[i])
		sumx2 = sumx2 + (x[i] * x[i])
		sumy2 = sumy2 + (y[i] * y[i])
	}
	r = (sumxy - (sumx * sumy)) / math.Sqrt((11*sumx2-(sumx*sumx))*(11*sumy2-(sumy*sumy)))
	fmt.Println("Correlation coefficient of set is : ", r)
	return
}
