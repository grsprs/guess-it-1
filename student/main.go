package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// Window maintains a sliding window of recent data points
type Window struct {
	data []float64 // Historical values
	size int       // Maximum window size
}

// NewWindow creates a new sliding window with specified size
func NewWindow(size int) *Window {
	return &Window{data: make([]float64, 0, size), size: size}
}

// Add inserts a new value and maintains window size
func (w *Window) Add(val float64) {
	w.data = append(w.data, val)
	if len(w.data) > w.size {
		w.data = w.data[1:] // Remove oldest value
	}
}

// Mean calculates the average of window values
func (w *Window) Mean() float64 {
	if len(w.data) == 0 {
		return 0
	}
	sum := 0.0
	for _, v := range w.data {
		sum += v
	}
	return sum / float64(len(w.data))
}

// Variance measures data spread using sample variance
func (w *Window) Variance() float64 {
	if len(w.data) < 2 {
		return 0
	}
	mean := w.Mean()
	sum := 0.0
	for _, v := range w.data {
		diff := v - mean
		sum += diff * diff
	}
	return sum / float64(len(w.data)-1) // Sample variance
}

// StdDev returns the standard deviation
func (w *Window) StdDev() float64 {
	return math.Sqrt(w.Variance())
}

// LinearRegression calculates trend using least squares method
func (w *Window) LinearRegression() (slope, intercept float64) {
	n := len(w.data)
	if n < 2 {
		return 0, w.Mean()
	}

	var sumX, sumY, sumXY, sumX2 float64
	for i, y := range w.data {
		x := float64(i)
		sumX += x
		sumY += y
		sumXY += x * y
		sumX2 += x * x
	}

	// Least squares formulas
	slope = (float64(n)*sumXY - sumX*sumY) / (float64(n)*sumX2 - sumX*sumX)
	intercept = (sumY - slope*sumX) / float64(n)
	return slope, intercept
}

// Predict returns the range [lower, upper] for the next value
func (w *Window) Predict() (lower, upper float64) {
	n := len(w.data)

	// Edge case: no data
	if n == 0 {
		return 0, 100
	}

	// Edge case: single point
	if n == 1 {
		val := w.data[0]
		return val - 50, val + 50
	}

	// Edge case: two points
	if n == 2 {
		mean := w.Mean()
		diff := math.Abs(w.data[1] - w.data[0])
		margin := math.Max(diff*2, 10)
		return mean - margin, mean + margin
	}

	// Full prediction: linear regression + confidence interval
	slope, intercept := w.LinearRegression()
	predicted := slope*float64(n) + intercept

	stdDev := w.StdDev()
	if stdDev < 1 {
		stdDev = 1 // Minimum standard deviation
	}

	// 95% confidence interval (±1.96σ)
	margin := 1.96 * stdDev

	// Adjust for high variance data
	variance := w.Variance()
	if variance > 100 {
		margin *= 1.5
	}

	return predicted - margin, predicted + margin
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	window := NewWindow(7) // Sliding window of 7 data points

	for scanner.Scan() {
		line := scanner.Text()
		val, err := strconv.ParseFloat(line, 64)
		if err != nil {
			continue // Skip invalid input
		}

		// Predict range before adding current value
		lower, upper := window.Predict()
		fmt.Printf("%.0f %.0f\n", lower, upper)

		// Add current value to window for next prediction
		window.Add(val)
	}
}
