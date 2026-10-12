# ESEP Grade Calculator

A Go library that averages assignment, exam, and essay scores, then applies
weights of 50%, 35%, and 15%. Letter grades are A >= 90, B >= 80, C >= 70,
D >= 60, and F < 60.

## Setup

Install Go 1.24.4 or newer and Git, then clone the repository:

```sh
git clone https://github.com/bryancarrillo1/esep-grade-calculator.git
cd esep-grade-calculator/grade-calculator
```

Run the following commands inside `grade-calculator`, where `go.mod` is located.
No third-party dependencies are needed.

## Usage

This is a library, not a standalone application. Within package `esepunittests`:

```go
calculator := NewGradeCalculator()
calculator.AddGrade("Homework", 100, Assignment)
calculator.AddGrade("Exam", 95, Exam)
calculator.AddGrade("Essay", 91, Essay)
letter := calculator.GetFinalGrade() // "A"
```

## Tests and coverage

```sh
go test -v ./...
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
```
