package main

import (
	"math"
	"testing"
)

func TestWindowMean(t *testing.T) {
	w := NewWindow(5)
	w.Add(10)
	w.Add(20)
	w.Add(30)
	
	mean := w.Mean()
	expected := 20.0
	if math.Abs(mean-expected) > 0.01 {
		t.Errorf("Mean = %.2f, want %.2f", mean, expected)
	}
}

func TestWindowStdDev(t *testing.T) {
	w := NewWindow(5)
	w.Add(10)
	w.Add(20)
	w.Add(30)
	
	stdDev := w.StdDev()
	if stdDev <= 0 {
		t.Errorf("StdDev should be positive, got %.2f", stdDev)
	}
}

func TestLinearRegression(t *testing.T) {
	w := NewWindow(5)
	w.Add(1)
	w.Add(2)
	w.Add(3)
	w.Add(4)
	
	slope, _ := w.LinearRegression()
	if math.Abs(slope-1.0) > 0.01 {
		t.Errorf("Slope = %.2f, want 1.0", slope)
	}
}

func TestPredictEdgeCases(t *testing.T) {
	w := NewWindow(5)
	lower, upper := w.Predict()
	
	if lower >= upper {
		t.Errorf("Lower bound should be less than upper bound")
	}
}

func TestPredictWithData(t *testing.T) {
	w := NewWindow(5)
	w.Add(100)
	w.Add(110)
	w.Add(120)
	
	lower, upper := w.Predict()
	
	if lower >= upper {
		t.Errorf("Lower bound should be less than upper bound")
	}
}
