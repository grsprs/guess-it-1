# Docker Testing Guide

## Prerequisites

**Docker Desktop must be running!**

1. Start Docker Desktop
2. Wait for it to fully start (icon in system tray)
3. Then run the commands below

## Quick Start

### Build and Run
```bash
docker compose up -d
```

### Test with Data
```bash
# Test with your data file
docker exec -i guess-it-student sh -c "cd student && go run main.go" < data.txt

# Test with sample data
echo -e "189\n113\n121\n114\n145" | docker exec -i guess-it-student sh -c "cd student && go run main.go"
```

### Run Tests
```bash
docker exec guess-it-student sh -c "cd student && go test -v"
```

### Stop
```bash
docker compose down
```

## Visual Testing (With Provided Tester)

For visual comparison with baseline algorithms:

1. Download guess-it-dockerized tester
2. Copy files:
   ```bash
   cp -r student/ guess-it-dockerized/
   cp script.sh guess-it-dockerized/
   ```
3. Start tester:
   ```bash
   cd guess-it-dockerized
   docker compose up
   ```
4. Open browser: http://localhost:3000/?guesser=big-range
5. Select dataset and compare

## Benefits

- ✅ Isolated environment
- ✅ No local Go installation needed
- ✅ Consistent testing across systems
- ✅ Visual comparison with baselines
- ✅ Easy cleanup

## Note

Docker is **optional**. The program works perfectly without it using:
```bash
cd student
go run main.go < data.txt
```
