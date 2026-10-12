# ESEP Grade Calculator

A Go library that calculates a course grade from assignments, exams, and essays.
It is a library package, so use the test runner or import it from another Go program;
there is no command-line `main` application.

## Requirements and layout

Install Go 1.24.4 or newer and Git. No third-party Go dependencies are required.

- `README.MD`: project usage and development instructions.
- `grade-calculator/go.mod`: module `esep/grade-calculator` and minimum Go version.
- `grade-calculator/grade_calculator.go`: grade types, storage, weighted averages,
  and letter-grade evaluation; package name `esepunittests`.
- `grade-calculator/grade_calculator_test.go`: tests using Go's `testing` package.

## Setup and tests

Fork the original repository on GitHub, clone your fork, and enter its module:

```sh
git clone https://github.com/YOUR-USERNAME/esep-grade-calculator.git
cd esep-grade-calculator/grade-calculator
go version
go test -v -count=1 ./...
```

The directory is `grade-calculator` (hyphen), not `grade_calculator`.
Run Go commands inside that directory because the repository root has no go.mod.
`-v` lists tests; `-count=1` runs them without reusing cached results.

## Coverage and checks

```sh
go test -count=1 -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
go tool cover -html=coverage.out -o coverage.html
go vet ./...
```

The text report gives statement coverage by function and overall. The HTML report
shows executed and unexecuted source statements. Coverage does not establish that
assertions are meaningful or all input combinations were checked.

## Public API and intended calculation

```go
calculator := NewGradeCalculator()
calculator.AddGrade("Homework 1", 100, Assignment)
calculator.AddGrade("Exam 1", 95, Exam)
calculator.AddGrade("Essay 1", 91, Essay)
letter := calculator.GetFinalGrade()
```

This snippet runs within package `esepunittests`. From a separate Go package,
import `esep/grade-calculator` with an alias and qualify the exported names.
Average scores within each category first, then apply assignment/exam/essay
weights of 50%/35%/15%. The example's intended total is 96.9.
The starter implementation uses A >= 90, B >= 80, C >= 70, D >= 60, and F below 60.
These cutoffs should be checked against the assignment's complete requirements.

## Starter status

At the original snapshot, TestGetGradeA and TestGetGradeB fail, while
TestGetGradeF passes incorrectly. The averaging loop sums indices, the essay
average reads exams, and the F test's high scores should yield A. Empty categories
also cause integer division by zero. Later assignment commits repair these issues
and expand the tests; inspect the commit history to reproduce each stage.

## Resources

- https://go.dev/doc/install
- https://go.dev/doc/tutorial/add-a-test
- https://go.dev/blog/cover
- https://pkg.go.dev/testing
