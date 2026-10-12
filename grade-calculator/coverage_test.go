package esepunittests

import "testing"

// These tests use behavior rather than GradeCalculator's private storage.
func TestLetterGradeBoundaries(t *testing.T) {
	cases := []struct {
		score int
		want  string
	}{
		{0, "F"}, {59, "F"}, {60, "D"}, {69, "D"},
		{70, "C"}, {79, "C"}, {80, "B"}, {89, "B"},
		{90, "A"}, {100, "A"},
	}
	for _, tc := range cases {
		gc := NewGradeCalculator()
		gc.AddGrade("assignment", tc.score, Assignment)
		gc.AddGrade("exam", tc.score, Exam)
		gc.AddGrade("essay", tc.score, Essay)
		if got := gc.GetFinalGrade(); got != tc.want {
			t.Errorf("score %d: got %s, want %s", tc.score, got, tc.want)
		}
	}
}

func TestCategoryAveragesAndWeights(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("assignment 1", 100, Assignment)
	gc.AddGrade("assignment 2", 80, Assignment)
	gc.AddGrade("exam 1", 80, Exam)
	gc.AddGrade("exam 2", 60, Exam)
	gc.AddGrade("essay 1", 100, Essay)
	// Category averages: 90, 70, 100. Weighted total: 84.5 (B).
	if got := gc.GetFinalGrade(); got != "B" {
		t.Errorf("unequal category sizes: got %s, want B", got)
	}
}

func TestEssayUsesItsOwnScores(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("assignment", 80, Assignment)
	gc.AddGrade("exam", 60, Exam)
	gc.AddGrade("essay", 100, Essay)
	// Correct: 76 (C). Reusing exam scores for essays gives 70 (C),
	// so use a second essay to make the error cross a grade boundary.
	gc.AddGrade("essay 2", 0, Essay)
	// Correct: 68.5 (D); the buggy exam substitution gives 70 (C).
	if got := gc.GetFinalGrade(); got != "D" {
		t.Errorf("essay average: got %s, want D", got)
	}
}

func TestFractionalCategoryAverage(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("assignment 1", 89, Assignment)
	gc.AddGrade("assignment 2", 90, Assignment)
	gc.AddGrade("exam", 91, Exam)
	gc.AddGrade("essay", 90, Essay)
	// 89.5*.5 + 91*.35 + 90*.15 = 90.1; premature truncation yields 89.85.
	if got := gc.GetFinalGrade(); got != "A" {
		t.Errorf("fractional average: got %s, want A", got)
	}
}

func TestEmptyCategories(t *testing.T) {
	gc := NewGradeCalculator()
	if got := gc.GetFinalGrade(); got != "F" {
		t.Errorf("empty calculator: got %s, want F", got)
	}
	gc.AddGrade("assignment only", 100, Assignment)
	if got := gc.GetFinalGrade(); got != "F" {
		t.Errorf("missing exam and essay contribute zero: got %s, want F", got)
	}
}

func TestGradeTypeNames(t *testing.T) {
	for _, tc := range []struct {
		kind GradeType
		want string
	}{{Assignment, "assignment"}, {Exam, "exam"}, {Essay, "essay"}, {GradeType(99), ""}} {
		if got := tc.kind.String(); got != tc.want {
			t.Errorf("type %d: got %q, want %q", int(tc.kind), got, tc.want)
		}
	}
}

func TestUnknownGradeTypeIsIgnored(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("unknown", 100, GradeType(99))
	if got := gc.GetFinalGrade(); got != "F" {
		t.Errorf("unknown category: got %s, want F", got)
	}
}
