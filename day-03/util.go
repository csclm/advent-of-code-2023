package main

func isSchematicSymbol(char rune) bool {
	if char == '.' {
		return false
	}
	if char >= 33 && char <= 47 {
		return true
	}
	if char >= 58 && char <= 64 {
		return true
	}
	if char >= 133 && char <= 140 {
		return true
	}
	if char == 126 {
		return true
	}
	return false
}
