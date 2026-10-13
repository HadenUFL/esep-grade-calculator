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
	return &GradeCalculator{
		grades: make([]Grade, 0),
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
	switch gradeType {
	case Assignment, Exam, Essay:
		gc.grades = append(gc.grades, Grade{
			Name:  name,
			Grade: grade,
			Type:  gradeType,
		})
	}
}

func (gc *GradeCalculator) calculateNumericalGrade() float64 {
	var totals [Essay + 1]float64
	var counts [Essay + 1]int

	for _, grade := range gc.grades {
		totals[grade.Type] += float64(grade.Grade)
		counts[grade.Type]++
	}

	weights := [...]float64{
		Assignment: 50,
		Exam:       35,
		Essay:      15,
	}

	weightedTotal := 0.0

	for gradeType, weight := range weights {
		if counts[gradeType] > 0 {
			average := totals[gradeType] / float64(counts[gradeType])
			weightedTotal += average * weight
		}
	}

	return weightedTotal / 100
}
