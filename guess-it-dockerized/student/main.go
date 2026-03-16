package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var data []float64
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		num, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			continue
		}

		data = append(data, num)

		if len(data) < 2 {
			continue
		}

		// Window of last 4 values
		start := len(data) - 4
		if start < 0 {
			start = 0
		}
		window := data[start:]

		// Mean
		sum := 0.0
		for _, v := range window {
			sum += v
		}
		mean := sum / float64(len(window))

		// Standard deviation
		varSum := 0.0
		for _, v := range window {
			d := v - mean
			varSum += d * d
		}
		stdDev := math.Sqrt(varSum / float64(len(window)))

		lower := math.Round(mean - 2*stdDev)
		upper := math.Round(mean + 2*stdDev)

		fmt.Printf("%.0f %.0f\n", lower, upper)
	}
}
