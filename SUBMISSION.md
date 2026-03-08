# Submission Instructions

## What to Submit

Submit the entire `guess-it-1-final/` folder containing:

### Required Files
```
guess-it-1-final/
├── student/
│   ├── main.go       # Main program (commented)
│   └── go.mod        # Go module
└── script.sh         # Execution script
```

### Documentation Files
```
├── README.md         # Complete documentation
├── AUDIT.md          # Audit guidelines
└── SUMMARY.md        # Project summary
```

### Optional Files
```
├── tester.go         # Standalone test program
└── TEST_RESULTS.md   # Test results reference
```

## Pre-Submission Checklist

### Code Quality
- ✅ All functions have comments
- ✅ Code is clean and readable
- ✅ No hardcoded dataset-specific values
- ✅ Proper error handling

### Functionality
- ✅ Program reads from stdin
- ✅ Outputs format: `lower upper`
- ✅ Handles edge cases (0, 1, 2 points)
- ✅ Uses statistical methods

### Files
- ✅ `script.sh` is executable
- ✅ `go.mod` specifies Go 1.21
- ✅ No external dependencies
- ✅ Documentation is complete

### Testing
- ✅ Program compiles without errors
- ✅ Runs successfully with sample data
- ✅ Produces valid output format

## Verification Commands

```bash
# Check Go version
go version

# Test compilation
cd student
go build main.go

# Test execution
echo -e "189\n113\n121" | go run main.go

# Expected output (example):
# 0 100
# 120 200
# 160 230
```

## Submission Format

### Option 1: Archive
```bash
# Create zip archive
zip -r guess-it-1-final.zip guess-it-1-final/

# Or tar archive
tar -czf guess-it-1-final.tar.gz guess-it-1-final/
```

### Option 2: Git Repository
```bash
# Initialize git (if not already)
cd guess-it-1-final
git init
git add student/ script.sh README.md AUDIT.md SUMMARY.md
git commit -m "Initial submission"
```

## Important Notes

### Do NOT Include
- ❌ `guess-it-dockerized/` folder (tester)
- ❌ Test data files
- ❌ Compiled binaries
- ❌ IDE configuration files
- ❌ Personal notes or drafts

### Do Include
- ✅ Source code (`student/main.go`)
- ✅ Module file (`student/go.mod`)
- ✅ Execution script (`script.sh`)
- ✅ Documentation (README, AUDIT, SUMMARY)

## Final Check

Before submission, verify:

1. **Structure**: All required files in correct locations
2. **Compilation**: `go build` succeeds without errors
3. **Execution**: `./script.sh` works correctly
4. **Documentation**: README explains algorithm clearly
5. **Comments**: Code has comprehensive comments
6. **Clean**: No unnecessary files included

## Support

If you encounter issues:
1. Check Go version (must be 1.21+)
2. Verify file permissions on `script.sh`
3. Test with sample data first
4. Review AUDIT.md for requirements

## Submission Deadline

Follow your institution's submission guidelines and deadlines.

---

**Ready to submit**: Once all checklist items are verified ✅
