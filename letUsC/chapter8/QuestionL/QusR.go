package main

import "fmt"

/*
*This function is used to find the value of a and b by given formulae
 */
func main() {
	var x [10]float64
	var y [10]float64
	sumx := 0.0
	sumy := 0.0
	sumxy := 0.0
	sumx2 := 0.0
	fmt.Println("Enter the value of x and y")
	for i := 0; i < 10; i++ {
		fmt.Scanln(&x[i], &y[i])
		sumx = sumx + x[i]
		sumy = sumy + y[i]
		sumxy = sumxy + (x[i] * y[i])
		sumx2 = sumx2 + (x[i] * x[i])
	}
	meanX := sumx / 10
	meanY := sumy / 10
	b := ((10*sumxy)-(sumx*sumy))/10*sumx2 - (sumx * sumx)
	a := meanY - b*meanX
	fmt.Println("value of a is:", a, "and The value of b is", b)
	return
}
