package main

import "fmt"

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))

}
func main() {
	a := []float64{1, 2, 3, 4, 5}
	fmt.Println(average(a))
}
