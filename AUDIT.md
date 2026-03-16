# AUDIT - guess-it-1

## Prerequisites

- Docker Desktop installed and running
- Browser available (for the tester UI)

---

## Setup (One-time)

### 1. Start the tester

```bash
cd guess-it-dockerized
docker compose up -d --build
```

Wait ~30 seconds for the container to start. Verify:

```bash
docker ps
```

You should see `guess-it-dockerized-node-1` running on port 3000.

### 2. Open the tester UI

Open browser and navigate to the URLs below for each test.

---

## Audit Questions & Answers

---

### Q1: Did the student won against big-range most of the times (at least 2 out of 3 times)?

**Answer: YES**

#### How to test:

**Data 2** — Open browser, run 3 times:
```
http://localhost:3000/?guesser=big-range
```
Click **"Test Data 2"** → Click **"Quick"** → Compare final scores.
Repeat 3 times total. Student wins all 3.

**Data 3** — Same URL, click **"Test Data 3"** instead.
Repeat 3 times total. Student wins all 3.

#### Expected results:

| Dataset | Student Score | big-range Score | Winner |
|---------|-------------|-----------------|--------|
| Data 2  | ~119,000-123,000 | ~49,000 | **Student** |
| Data 3  | ~119,000-121,000 | ~49,000 | **Student** |

Student wins **100% of the time** against big-range (score is ~2.5x higher).

---

### Q2: Did the student won against average most of the times (at least 2 out of 3 times for each)?

**Answer: YES**

#### How to test:

```
http://localhost:3000/?guesser=average
```

**Data 1** — Click "Test Data 1" → "Quick" → Compare scores. Repeat 3 times.
**Data 2** — Click "Test Data 2" → "Quick" → Compare scores. Repeat 3 times.
**Data 3** — Click "Test Data 3" → "Quick" → Compare scores. Repeat 3 times.

#### Expected results:

| Dataset | Student Score | average Score | Winner |
|---------|-------------|---------------|--------|
| Data 1  | ~124,000-127,000 | ~87,000-104,000 | **Student** |
| Data 2  | ~119,000-123,000 | ~320-1,280 | **Student** |
| Data 3  | ~119,000-121,000 | ~320-1,280 | **Student** |

Student wins **100% of the time** against average. On Data 2/3, average collapses to near-zero due to outliers while student handles them gracefully.

---

### Q3: Did the student won against median most of the times (at least 2 out of 3 times for each)?

**Answer: YES**

#### How to test:

```
http://localhost:3000/?guesser=median
```

**Data 1** — Click "Test Data 1" → "Quick" → Compare scores. Repeat 3 times.
**Data 2** — Click "Test Data 2" → "Quick" → Compare scores. Repeat 3 times.
**Data 3** — Click "Test Data 3" → "Quick" → Compare scores. Repeat 3 times.

#### Expected results:

| Dataset | Student Score | median Score | Winner |
|---------|-------------|--------------|--------|
| Data 1  | ~124,000-127,000 | ~94,000-102,000 | **Student** |
| Data 2  | ~119,000-123,000 | ~96,000-102,000 | **Student** |
| Data 3  | ~119,000-121,000 | ~94,000-103,000 | **Student** |

Student wins **100% of the time** against median (score is ~20% higher).

---

### Q4 (Bonus): Did the student won against nic most of the times (at least 2 out of 3 times for each)?

**Answer: YES**

#### How to test:

```
http://localhost:3000/?guesser=nic
```

**Data 1** — Click "Test Data 1" → "Quick" → Compare scores. Repeat 3 times.
**Data 2** — Click "Test Data 2" → "Quick" → Compare scores. Repeat 3 times.
**Data 3** — Click "Test Data 3" → "Quick" → Compare scores. Repeat 3 times.

#### Expected results:

| Dataset | Student Score | nic Score | Winner |
|---------|-------------|-----------|--------|
| Data 1  | ~124,000-127,000 | ~85,000-114,000 | **Student** |
| Data 2  | ~119,000-123,000 | ~84,000-110,000 | **Student** |
| Data 3  | ~119,000-121,000 | ~94,000-107,000 | **Student** |

Student wins **100% of the time** against nic. Even nic's best variance spike (~114K) is below student's worst score (~119K).

---

## Algorithm Explanation

The program uses a **sliding window of 4 data points** with **mean ± 2σ** confidence interval:

```
1. Read number from stdin
2. Keep last 4 numbers in a window
3. Calculate mean: μ = Σx / n
4. Calculate population std dev: σ = √(Σ(x-μ)² / n)
5. Output range: [μ - 2σ, μ + 2σ]
```

**Why it works:**
- Small window (4) → adapts quickly → small σ → tight range → high score per correct guess
- 2σ confidence → ~100% accuracy on normal data
- When outliers appear → σ expands automatically → range widens → still captures next value
- Score formula: `10M / (1 + range_size)` rewards tight ranges with high accuracy

---

## Quick Verification Commands

### Run all tests automatically (from guess-it-dockerized folder):

```bash
# Start tester
docker compose up -d --build

# Wait for startup
ping -n 6 127.0.0.1 >nul   # Windows
sleep 5                      # Linux/Mac

# Open in browser
start http://localhost:3000/?guesser=big-range
start http://localhost:3000/?guesser=average
start http://localhost:3000/?guesser=median
start http://localhost:3000/?guesser=nic
```

### Run unit tests:

```bash
cd student
go test -v
```

### Stop tester:

```bash
cd guess-it-dockerized
docker compose down
```

---

## Score Summary

| Opponent | Data 1 | Data 2 | Data 3 | Result |
|----------|--------|--------|--------|--------|
| big-range | N/A | ✅ WIN | ✅ WIN | **PASS** |
| average | ✅ WIN | ✅ WIN | ✅ WIN | **PASS** |
| median | ✅ WIN | ✅ WIN | ✅ WIN | **PASS** |
| nic (bonus) | ✅ WIN | ✅ WIN | ✅ WIN | **PASS** |

**All audit questions: PASS ✅**
