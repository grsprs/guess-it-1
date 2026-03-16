package main

import (
	"math"
	"testing"
)

func meanStdDev(vals []float64) (float64, float64) {
	sum := 0.0
	for _, v := range vals {
		sum += v
	}
	mean := sum / float64(len(vals))

	varSum := 0.0
	for _, v := range vals {
		d := v - mean
		varSum += d * d
	}
	return mean, math.Sqrt(varSum / float64(len(vals)))
}

func TestMeanStdDev(t *testing.T) {
	vals := []float64{150, 160, 140, 155}
	mean, sd := meanStdDev(vals)

	if math.Abs(mean-151.25) > 0.01 {
		t.Errorf("Mean = %f, want 151.25", mean)
	}
	if sd < 7 || sd > 8 {
		t.Errorf("StdDev = %f, want ~7.4", sd)
	}
}

func TestWindowSize(t *testing.T) {
	data := []float64{100, 120, 140, 160, 180, 200}
	start := len(data) - 4
	if start < 0 {
		start = 0
	}
	window := data[start:]

	if len(window) != 4 {
		t.Errorf("Window size = %d, want 4", len(window))
	}
	if window[0] != 140 {
		t.Errorf("Window[0] = %f, want 140", window[0])
	}
}

func TestRangeCalculation(t *testing.T) {
	vals := []float64{150, 150, 150, 150}
	mean, sd := meanStdDev(vals)

	lower := math.Round(mean - 2*sd)
	upper := math.Round(mean + 2*sd)

	if lower != 150 || upper != 150 {
		t.Errorf("Identical values: range [%f, %f], want [150, 150]", lower, upper)
	}
}

func TestRangeWithVariance(t *testing.T) {
	vals := []float64{100, 200, 100, 200}
	mean, sd := meanStdDev(vals)

	lower := math.Round(mean - 2*sd)
	upper := math.Round(mean + 2*sd)

	if lower > 100 {
		t.Errorf("Lower = %f, should be <= 100", lower)
	}
	if upper < 200 {
		t.Errorf("Upper = %f, should be >= 200", upper)
	}
}

func TestSmallWindow(t *testing.T) {
	data := []float64{150, 160}
	start := len(data) - 4
	if start < 0 {
		start = 0
	}
	window := data[start:]

	if len(window) != 2 {
		t.Errorf("Window size = %d, want 2", len(window))
	}
}
