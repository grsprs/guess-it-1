# Audit Guidelines

## Project Setup

### File Structure Verification

Ensure the following structure exists:

```
guess-it-1-final/
├── student/
│   ├── main.go       # Main program
│   └── go.mod        # Go module (go 1.21)
└── script.sh         # Execution script
```

### Script Verification

The `script.sh` must:
- Be executable
- Run from the tester's root directory
- Execute the student program correctly

Content:
```bash
#!/bin/sh
cd student
go run main.go
```

## Testing Instructions

### Prerequisites

1. Download the tester (guess-it-dockerized)
2. Place the `student/` folder in the tester's root directory
3. Ensure Go 1.21+ is installed

### Test Execution

The program will be tested against multiple AI opponents using Data 1, 2, and 3.

#### Required Tests

Test the student program 3 times on each dataset against:
- `big-range`
- `average`
- `median`

#### Bonus Test

Test the student program 3 times on each dataset against:
- `nic`

### Using the Standalone Tester

A standalone Go tester is provided for convenience:

```bash
# Test specific opponent and dataset
go run tester.go <opponent> <dataset>

# Examples
go run tester.go big-range 1
go run tester.go average 2
go run tester.go median 3
go run tester.go nic 1
```

The tester will:
1. Run 3 files from the specified dataset
2. Compare student vs opponent scores
3. Report wins/losses
4. Show pass/fail status (requires 2/3 wins)

## Evaluation Criteria

### Scoring System

For each prediction:
- If the actual value falls within the predicted range: `score += 1000 / (upper - lower)`
- Smaller ranges yield higher scores
- Incorrect predictions score 0

### Success Criteria

The student must win at least 2 out of 3 times for each dataset against each opponent.

**Required**:
- Beat `big-range` on Data 1, 2, 3 (2/3 wins each)
- Beat `average` on Data 1, 2, 3 (2/3 wins each)
- Beat `median` on Data 1, 2, 3 (2/3 wins each)

**Bonus**:
- Beat `nic` on Data 1, 2, 3 (2/3 wins each)

## Code Review

### Quality Checks

1. **Language**: Go (golang)
2. **Dependencies**: None (standard library only)
3. **Code Quality**:
   - Clean, readable structure
   - Proper comments
   - No hardcoded dataset-specific values
   - Generic algorithm

4. **Algorithm**:
   - Uses mathematical/statistical methods
   - Handles edge cases
   - Balances accuracy vs range size

### Expected Behavior

- Reads numbers from stdin (one per line)
- Outputs ranges in format: `lower upper`
- Handles invalid input gracefully
- Terminates cleanly on EOF

## Common Issues

### Script Execution

If the script fails:
- Check file permissions (`chmod +x script.sh`)
- Verify Go installation (`go version`)
- Ensure correct working directory

### Performance

The program should:
- Execute quickly (< 1 second per file)
- Use minimal memory
- Handle large datasets (12,000+ numbers)

## Verification Steps

1. ✅ Check file structure
2. ✅ Verify script.sh is executable
3. ✅ Test with Data 1 against big-range (3 times)
4. ✅ Test with Data 2 against big-range (3 times)
5. ✅ Test with Data 3 against big-range (3 times)
6. ✅ Repeat for average opponent
7. ✅ Repeat for median opponent
8. ✅ (Bonus) Repeat for nic opponent
9. ✅ Verify 2/3 wins for each dataset/opponent combination
10. ✅ Review code quality and comments

## Expected Results

The student program should:
- Achieve ~96% prediction accuracy
- Score ~70,000-76,000 points per file
- Win consistently against all opponents
- Demonstrate understanding of statistical methods

## Notes

- The program uses linear regression with a sliding window
- Confidence intervals balance accuracy and range size
- Edge cases are handled appropriately
- No external dependencies required
