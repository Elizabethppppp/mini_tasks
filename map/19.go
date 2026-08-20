package _map

func SetGrade(m map[string]map[string]int, student string, subject string, grade int) {
	if _, ok := m[student]; !ok {
		m[student] = make(map[string]int)
	}
	m[student][subject] = grade
}

func GetGrade(m map[string]map[string]int, student string, subject string) (int, bool) {
	studentGrades, ok := m[student]
	if !ok {
		return 0, false
	}

	grade, ok := studentGrades[subject]
	if !ok {
		return 0, false
	}
	return grade, true
}
