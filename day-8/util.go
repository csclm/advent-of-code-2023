package main

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

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}
