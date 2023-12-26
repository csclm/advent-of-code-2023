package day8

// truncates a string such that it has no repetitions, e.g. ABCABC -> ABC
func truncateSymmetricalString(directions string) string {
	symmetry := findRotationalSymmetry(directions)
	return directions[:symmetry]
}

func findRotationalSymmetry(s string) int {
	for i := 1; i < len(s)/2; i++ {
		rotated := s[i:] + s[:i]
		if rotated == s {
			return i
		}
	}
	return len(s)
}
