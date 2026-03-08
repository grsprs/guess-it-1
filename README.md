# guess-it-1 - Statistical Prediction System

## Repository Links
- **Zone01**: https://platform.zone01.gr/git/snikolou/guess-it-1.git
- **GitHub**: https://github.com/grsprs/guess-it-1.git

A Go program that predicts numerical ranges using linear regression with a sliding window approach.

## Overview

This program reads numbers from standard input and predicts a range `[lower, upper]` for each subsequent number. It uses statistical methods to balance prediction accuracy with range size optimization.

## Algorithm

### Core Components

1. **Sliding Window (7 points)**
   - Maintains the last 7 data points
   - Adapts to recent trends
   - O(1) space complexity

2. **Linear Regression**
   - Detects trends using least squares method
   - Calculates slope and intercept
   - Predicts next value based on trend

3. **Confidence Intervals**
   - 95% confidence level (±1.96σ)
   - Uses standard deviation for margin calculation
   - Adjusts for high variance (>100)

4. **Edge Case Handling**
   - 0 points: Default range [0, 100]
   - 1 point: value ± 50
   - 2 points: mean ± 2×difference
   - 3+ points: Full statistical prediction

### Statistical Formulas

**Mean**: `μ = Σx / n`

**Variance**: `σ² = Σ(x - μ)² / (n - 1)`

**Standard Deviation**: `σ = √σ²`

**Linear Regression**:
- Slope: `m = (n·Σxy - Σx·Σy) / (n·Σx² - (Σx)²)`
- Intercept: `b = (Σy - m·Σx) / n`
- Prediction: `y = m·x + b`

**Confidence Interval**:
- Margin: `1.96 × σ`
- If variance > 100: `margin × 1.5`
- Range: `[predicted - margin, predicted + margin]`

## Project Structure

```
guess-it-1-final/
├── student/
│   ├── main.go       # Main program with prediction logic
│   └── go.mod        # Go module definition
├── script.sh         # Execution script for tester
├── tester.go         # Standalone test program
├── README.md         # This file
└── AUDIT.md          # Audit guidelines
```

## Requirements

- Go 1.21 or higher
- No external dependencies (pure Go standard library)

## Usage

### Standard Execution

```bash
# From project root
./script.sh < data.txt

# Or directly
cd student
go run main.go < ../data.txt
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
0 100
120 200
160 230
110 140
100 200
```

## Testing

### With Provided Tester

1. Place the `student/` folder in the tester's root directory
2. Ensure `script.sh` is executable
3. Run the tester according to its instructions

### Manual Testing

```bash
# Test against specific opponent and dataset
go run tester.go <opponent> <dataset>

# Examples
go run tester.go big-range 1
go run tester.go average 2
go run tester.go median 3
go run tester.go nic 1
```

## Performance

- **Time Complexity**: O(n) per prediction where n = window size (7)
- **Space Complexity**: O(1) - constant space (7 floats)
- **Accuracy**: ~96% correct predictions
- **Average Score**: ~72,000 points per file

## Implementation Details

### Why This Works

1. **Linear Regression** captures trends (increasing/decreasing patterns)
2. **Sliding Window** adapts to recent changes, ignoring old data
3. **Confidence Intervals** balance accuracy vs range size
4. **Variance Adjustment** handles volatile data appropriately
5. **Edge Case Fallbacks** ensure robustness with limited data

### Design Decisions

- **Window Size (7)**: Balances responsiveness vs stability
- **95% Confidence**: Standard statistical confidence level
- **Variance Threshold (100)**: Empirically determined for data characteristics
- **Minimum StdDev (1)**: Prevents division by zero and overly narrow ranges

## Code Quality

- Clean, readable code structure
- Comprehensive comments explaining logic
- No hardcoded values for specific datasets
- Generic algorithm works on any numerical sequence
- Proper error handling for invalid input

## Author

Zone01 Module #585 - Statistical Prediction Exercise
