# Project Summary

## Deliverables

### Core Files (Required)
- `student/main.go` - Main prediction program
- `student/go.mod` - Go module definition
- `script.sh` - Execution script

### Documentation
- `README.md` - Complete usage guide
- `AUDIT.md` - Audit checklist
- `SUBMISSION.md` - Submission instructions

## Algorithm

Linear regression with sliding window (7 points)
- 95% confidence intervals
- Variance-based adjustment
- Edge case handling

## Requirements

- Go 1.21+
- No external dependencies
- No Docker required
- No npm/node required

## Testing

Use the standalone tester:
```bash
go run tester.go <opponent> <dataset>
```

Ready for auditor testing with Data 1, 2, 3.

## Compliance

✅ All audit requirements met
✅ Autonomous project
✅ No external tools needed
✅ Complete documentation
