package esepunittests

type GradeCalculator struct {
	assignments []Grade
	exams       []Grade
	essays      []Grade
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
	return &GradeCalculator{
		assignments: make([]Grade, 0),
		exams:       make([]Grade, 0),
		essays:      make([]Grade, 0),
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	switch {
	case numericalGrade >= 90:
		return "A"
	case numericalGrade >= 80:
		return "B"
	case numericalGrade >= 70:
		return "C"
	case numericalGrade >= 60:
		return "D"
	default:
		return "F"
	}
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	entry := Grade{
		Name:  name,
		Grade: grade,
		Type:  gradeType,
	}

	switch gradeType {
	case Assignment:
		gc.assignments = append(gc.assignments, entry)
	case Exam:
		gc.exams = append(gc.exams, entry)
	case Essay:
		gc.essays = append(gc.essays, entry)
	}
}

func (gc *GradeCalculator) calculateNumericalGrade() float64 {
	assignmentAverage := computeAverage(gc.assignments)
	examAverage := computeAverage(gc.exams)
	essayAverage := computeAverage(gc.essays)

	return (assignmentAverage*50 + examAverage*35 + essayAverage*15) / 100
}

func computeAverage(grades []Grade) float64 {
	if len(grades) == 0 {
		return 0
	}

	total := 0.0

	for _, grade := range grades {
		total += float64(grade.Grade)
	}

	return total / float64(len(grades))
}
