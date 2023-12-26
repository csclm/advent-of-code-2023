package day15

import "slices"

func HolidayHash(s string) int {
	curr := 0
	for _, char := range s {
		curr += int(char)
		curr *= 17
		curr %= 256
	}
	return curr
}

type HolidayHashmap struct {
	contents [256][]Lens
}

type Lens struct {
	label       string
	focalLength int
}

func (hm *HolidayHashmap) Insert(label string, focalLength int) {
	boxNum := HolidayHash(label)
	box := hm.contents[boxNum]
	existing := slices.IndexFunc(box, func(entry Lens) bool { return entry.label == label })
	if existing == -1 {
		hm.contents[boxNum] = append(box, Lens{label, focalLength})
	} else {
		hm.contents[boxNum][existing].focalLength = focalLength
	}
}

func (hm *HolidayHashmap) Remove(label string) {
	boxNum := HolidayHash(label)
	box := hm.contents[boxNum]
	existing := slices.IndexFunc(box, func(entry Lens) bool { return entry.label == label })
	if existing == -1 {
		return
	}
	hm.contents[boxNum] = slices.Delete(box, existing, existing+1)
}

func (hm *HolidayHashmap) TotalFocusingPower() int {
	totalPower := 0
	for boxNum, box := range hm.contents {
		for lensNum, lens := range box {
			totalPower += (boxNum + 1) * (lensNum + 1) * lens.focalLength
		}
	}
	return totalPower
}
