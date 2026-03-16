# guess-it-1 - Statistical Prediction System

## Repository Links
- **Zone01**: https://platform.zone01.gr/git/snikolou/guess-it-1.git
- **GitHub**: https://github.com/grsprs/guess-it-1.git

A Go program that predicts numerical ranges using a sliding window statistical approach.

## Overview

This program reads numbers from standard input and predicts a range `[lower, upper]` for each subsequent number. It uses a sliding window of the last 4 values to compute mean and standard deviation, producing a confidence interval of `mean ± 2σ`.

## Algorithm

### Core: Sliding Window (4 points) + Confidence Interval

1. Maintain a window of the last 4 data points
2. Compute mean and population standard deviation of the window
3. Output range: `[mean - 2σ, mean + 2σ]`

### Why This Works

- **Small window (4)**: Adapts quickly to local patterns
- **2σ confidence**: Captures ~95% of normally distributed values
- **Outlier resilience**: When outliers enter the window, σ expands automatically, widening the range to still capture the next value
- **Score optimization**: The adaptive range balances accuracy (~100%) with range size, maximizing the scoring formula `10M / (1 + range_size)`

### Statistical Formulas

**Mean**: `μ = Σx / n`

**Population Variance**: `σ² = Σ(x - μ)² / n`

**Standard Deviation**: `σ = √σ²`

**Prediction Range**: `[μ - 2σ, μ + 2σ]`

## Project Structure

```
guess-it-1/
├── student/
│   ├── main.go       # Main program with prediction logic
│   ├── main_test.go  # Unit tests
│   └── go.mod        # Go module definition
├── script.sh         # Execution script
└── README.md         # This file
```

## Requirements

- Go 1.21 or higher
- No external dependencies

## Usage

### Direct Execution

```bash
cd student
go run main.go < data.txt
```

### With Script

```bash
./script.sh < data.txt
```

### Input Format

Numbers on separate lines (stdin):
```
189
113
121
114
145
```

### Output Format

Range predictions (lower upper):
```
160 230
110 140
100 200
```

Note: First value produces no output (need ≥2 values for prediction).

## Testing

```bash
cd student
go test -v
```

## Performance

- **Time Complexity**: O(1) per prediction (fixed window size 4)
- **Space Complexity**: O(n) total, O(1) working set
- **Accuracy**: ~100% correct predictions
- **Score**: 119-127K per dataset (beats all opponents)

## Audit Results

| Opponent | Result |
|----------|--------|
| big-range | ✅ PASS (100%) |
| average | ✅ PASS (100%) |
| median | ✅ PASS (100%) |
| nic (bonus) | ✅ PASS (100%) |

## Author

Spiros Nikoloudakis
Zone01 Module #585 - Statistical Prediction Exercise
