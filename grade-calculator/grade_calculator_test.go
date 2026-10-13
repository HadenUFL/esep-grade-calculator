package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	want := "A"
	calculator := NewGradeCalculator()

	calculator.AddGrade("open source assignment", 100, Assignment)
	calculator.AddGrade("exam 1", 100, Exam)
	calculator.AddGrade("essay on ai ethics", 100, Essay)

	got := calculator.GetFinalGrade()

	if got != want {
		t.Errorf("GetFinalGrade() = %q, want %q", got, want)
	}
}

func TestGetGradeB(t *testing.T) {
	want := "B"
	calculator := NewGradeCalculator()

	calculator.AddGrade("open source assignment", 80, Assignment)
	calculator.AddGrade("exam 1", 81, Exam)
	calculator.AddGrade("essay on ai ethics", 85, Essay)

	got := calculator.GetFinalGrade()

	if got != want {
		t.Errorf("GetFinalGrade() = %q, want %q", got, want)
	}
}

func TestGetGradeF(t *testing.T) {
	want := "F"
	calculator := NewGradeCalculator()

	calculator.AddGrade("open source assignment", 50, Assignment)
	calculator.AddGrade("exam 1", 45, Exam)
	calculator.AddGrade("essay on ai ethics", 40, Essay)

	got := calculator.GetFinalGrade()

	if got != want {
		t.Errorf("GetFinalGrade() = %q, want %q", got, want)
	}
}
