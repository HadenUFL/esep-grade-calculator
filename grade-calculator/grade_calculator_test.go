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

func TestGetFinalGradeBoundaries(t *testing.T) {
	tests := []struct {
		name  string
		score int
		want  string
	}{
		{"Zero", 0, "F"},
		{"BelowD", 59, "F"},
		{"AtD", 60, "D"},
		{"BelowC", 69, "D"},
		{"AtC", 70, "C"},
		{"BelowB", 79, "C"},
		{"AtB", 80, "B"},
		{"BelowA", 89, "B"},
		{"AtA", 90, "A"},
		{"Maximum", 100, "A"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			calculator := NewGradeCalculator()

			calculator.AddGrade("assignment", test.score, Assignment)
			calculator.AddGrade("exam", test.score, Exam)
			calculator.AddGrade("essay", test.score, Essay)

			got := calculator.GetFinalGrade()

			if got != test.want {
				t.Errorf("GetFinalGrade() = %q, want %q", got, test.want)
			}
		})
	}
}

func TestGetFinalGradeCategoryAverages(t *testing.T) {
	tests := []struct {
		name        string
		assignments []int
		exams       []int
		essays      []int
		want        string
	}{
		{
			name:        "WeightedAverages",
			assignments: []int{100, 80},
			exams:       []int{100, 80},
			essays:      []int{0},
			want:        "C",
		},
		{
			name:        "FractionalAverage",
			assignments: []int{69, 70},
			exams:       []int{70},
			essays:      []int{72},
			want:        "C",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			calculator := NewGradeCalculator()

			groups := []struct {
				gradeType GradeType
				scores    []int
			}{
				{Assignment, test.assignments},
				{Exam, test.exams},
				{Essay, test.essays},
			}

			for _, group := range groups {
				for _, score := range group.scores {
					calculator.AddGrade("score", score, group.gradeType)
				}
			}

			got := calculator.GetFinalGrade()

			if got != test.want {
				t.Errorf("GetFinalGrade() = %q, want %q", got, test.want)
			}
		})
	}
}

func TestGetFinalGradeWithoutGrades(t *testing.T) {
	calculator := NewGradeCalculator()

	got := calculator.GetFinalGrade()

	if got != "F" {
		t.Errorf("GetFinalGrade() = %q, want %q", got, "F")
	}
}

func TestGradeTypeString(t *testing.T) {
	tests := []struct {
		name      string
		gradeType GradeType
		want      string
	}{
		{"Assignment", Assignment, "assignment"},
		{"Exam", Exam, "exam"},
		{"Essay", Essay, "essay"},
		{"Unknown", GradeType(-1), ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.gradeType.String()

			if got != test.want {
				t.Errorf("GradeType(%d).String() = %q, want %q",
					int(test.gradeType), got, test.want)
			}
		})
	}
}
