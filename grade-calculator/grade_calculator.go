package esepunittests

type GradeCalculator struct {
	grades []Grade
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{grades: make([]Grade, 0)}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	}

	return "F"
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	switch gradeType {
	case Assignment, Exam, Essay:
		gc.grades = append(gc.grades, Grade{Name: name, Grade: grade, Type: gradeType})
	}
}

func (gc *GradeCalculator) calculateNumericalGrade() float64 {
	assignmentAverage := computeAverage(gc.grades, Assignment)
	examAverage := computeAverage(gc.grades, Exam)
	essayAverage := computeAverage(gc.grades, Essay)
	return assignmentAverage*.5 + examAverage*.35 + essayAverage*.15
}

func computeAverage(grades []Grade, kind GradeType) float64 {
	sum, count := 0, 0
	for _, grade := range grades {
		if grade.Type == kind {
			sum += grade.Grade
			count++
		}
	}
	// An empty category contributes zero; its weight is not redistributed.
	if count == 0 {
		return 0
	}
	return float64(sum) / float64(count)
}
