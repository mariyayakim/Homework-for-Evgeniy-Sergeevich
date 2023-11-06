package main

import (
	"fmt"
	"math"
)

func main() {
	example()
}

func example() {
	a := 7.2
	b := 1.3
	xn := 1.56
	xk := 4.71
	deltaX := 0.63
	x1 := 2.4
	x2 := 2.8
	x3 := 3.9
	x4 := 4.7
	x5 := 3.16
	x := xn
	for x <= xk {
		result := (a + (b * x))
		result = math.Pow(result/(math.Log(x)/3.0), 1.0/5.0)
		fmt.Printf("x=%.2f, y=%.2f\n", x, result)
		x += deltaX
	}
	values := []float64{x1, x2, x3, x4, x5}

	for _, x := range values {
		result := (a + (b * x))
		result = math.Pow(result/(math.Log(x)/3.0), 1.0/5.0)
		formattedResult := fmt.Sprintf("%.3f", result)
		fmt.Println("При x=", x, " y=", formattedResult)
	}

}
