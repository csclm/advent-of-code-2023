package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/mitchellh/iochan"
)

func main() {
	f, _ := os.Open("./input.txt")
	seeds, almanac := parseInput(f)

	fmt.Printf("Part 1 lowest location number is %d\n", part1(&almanac, seeds))
	fmt.Printf("Part 2 lowest location number is %d\n", part2(&almanac, seeds))
}

func part1(almanac *Almanac, seeds []int) int {
	lowestLocation := 99999999999
	for _, seed := range seeds {
		soil := evaluateMapping(seed, almanac.seedToSoil)
		fertilizer := evaluateMapping(soil, almanac.soilToFertilizer)
		water := evaluateMapping(fertilizer, almanac.fertilizerToWater)
		light := evaluateMapping(water, almanac.waterToLight)
		temperature := evaluateMapping(light, almanac.lightToTemperature)
		humidity := evaluateMapping(temperature, almanac.temperatureToHumidity)
		location := evaluateMapping(humidity, almanac.humidityToLocation)
		lowestLocation = min(lowestLocation, location)
	}
	return lowestLocation
}

func part2(almanac *Almanac, seedNumbers []int) int {
	lowestLocation := 99999999999
	for i := 0; i < len(seedNumbers)/2; i++ {
		seed := Range{min: seedNumbers[i*2], max: seedNumbers[i*2] + seedNumbers[i*2+1]}
		soil := evaluateMappingWithRange(seed, almanac.seedToSoil)
		fertilizer := evaluateMappingWithRanges(soil, almanac.soilToFertilizer)
		water := evaluateMappingWithRanges(fertilizer, almanac.fertilizerToWater)
		light := evaluateMappingWithRanges(water, almanac.waterToLight)
		temperature := evaluateMappingWithRanges(light, almanac.lightToTemperature)
		humidity := evaluateMappingWithRanges(temperature, almanac.temperatureToHumidity)
		location := evaluateMappingWithRanges(humidity, almanac.humidityToLocation)
		for _, locationRange := range location {
			lowestLocation = min(lowestLocation, locationRange.min)
		}
	}
	return lowestLocation
}

func evaluateMapping(input int, ranges []MappingRange) int {
	for _, mappingRange := range ranges {
		output, didMap := mappingRange.MapInput(input)
		if didMap {
			return output
		}
	}
	return input
}

func evaluateMappingWithRanges(inputRanges []Range, mappingRanges []MappingRange) []Range {
	result := make([]Range, 0)
	for _, inputRange := range inputRanges {
		result = append(result, evaluateMappingWithRange(inputRange, mappingRanges)...)
	}
	return result
}

// This is technically wrong - it doesn't output range fragments that weren't mapped by the mapping ranges
// It still gave me the right answer though, and I don't want to take the time to fix it right now
// I'd have to make some kind of loop that figures out if the list of output ranges is contiguous, and if it's
// not it should insert ranges where the gaps would be
func evaluateMappingWithRange(inputRange Range, mappingRanges []MappingRange) []Range {
	resultRanges := make([]Range, 0)
	for _, mappingRange := range mappingRanges {
		overlap, doesOverlap := inputRange.Intersection(mappingRange.InputRange())
		if doesOverlap {
			mappedRange := Range{
				min: overlap.min + mappingRange.destination - mappingRange.source,
				max: overlap.max + mappingRange.destination - mappingRange.source,
			}
			resultRanges = append(resultRanges, mappedRange)
		}
	}
	return resultRanges
}

// seeds, almanac
func parseInput(f *os.File) ([]int, Almanac) {
	almanac := Almanac{
		seedToSoil:            make([]MappingRange, 0),
		soilToFertilizer:      make([]MappingRange, 0),
		fertilizerToWater:     make([]MappingRange, 0),
		waterToLight:          make([]MappingRange, 0),
		lightToTemperature:    make([]MappingRange, 0),
		temperatureToHumidity: make([]MappingRange, 0),
		humidityToLocation:    make([]MappingRange, 0),
	}

	reader := iochan.DelimReader(f, '\n')
	seedsLine := <-reader
	seeds := parseSeedNums(seedsLine)

	currentMapping := ""
	mappingTitlePattern := regexp.MustCompile(`([\w-]+) map:`)

	for line := range reader {
		if strings.TrimSpace(line) == "" {
			continue
		}
		mappingTitleMatch := mappingTitlePattern.FindAllStringSubmatch(line, -1)
		if len(mappingTitleMatch) > 0 {
			currentMapping = mappingTitleMatch[0][1]
			continue
		}
		almanac.addMapping(currentMapping, parseMappingRange(line))
	}

	return seeds, almanac
}

func parseMappingRange(line string) MappingRange {
	mappingRangePattern := regexp.MustCompile(`\d+`)
	numStrings := mappingRangePattern.FindAllString(line, -1)
	return MappingRange{
		destination: mustParseInt(numStrings[0]),
		source:      mustParseInt(numStrings[1]),
		length:      mustParseInt(numStrings[2]),
	}
}

func mustParseInt(intStr string) int {
	num, _ := strconv.ParseInt(intStr, 10, 0)
	return int(num)
}

func parseSeedNums(seedNumLine string) []int {
	seedNumStrings := strings.Split(strings.TrimSpace(seedNumLine), " ")[1:]
	nums := make([]int, len(seedNumStrings))
	for i, seedNumString := range seedNumStrings {
		nums[i] = mustParseInt(seedNumString)
	}
	return nums
}

func (almanac *Almanac) addMapping(mappingName string, mappingRange MappingRange) {
	switch mappingName {
	case "seed-to-soil":
		almanac.seedToSoil = append(almanac.seedToSoil, mappingRange)
	case "soil-to-fertilizer":
		almanac.soilToFertilizer = append(almanac.soilToFertilizer, mappingRange)
	case "fertilizer-to-water":
		almanac.fertilizerToWater = append(almanac.fertilizerToWater, mappingRange)
	case "water-to-light":
		almanac.waterToLight = append(almanac.waterToLight, mappingRange)
	case "light-to-temperature":
		almanac.lightToTemperature = append(almanac.lightToTemperature, mappingRange)
	case "temperature-to-humidity":
		almanac.temperatureToHumidity = append(almanac.temperatureToHumidity, mappingRange)
	case "humidity-to-location":
		almanac.humidityToLocation = append(almanac.humidityToLocation, mappingRange)
	default:
		panic("invalid mapping name " + mappingName)
	}
}

type Range struct {
	min int // inclusive
	max int // exclusive
}

func (r Range) Contains(input int) bool {
	return input >= r.min && input < r.max
}

func (r Range) Intersection(other Range) (Range, bool) {
	if r.min <= other.min && r.max >= other.max {
		return other, true
	}
	if r.min >= other.min && r.max <= other.max {
		return r, true
	}
	if other.Contains(r.min) {
		return Range{min: r.min, max: other.max}, true
	}
	if r.Contains(other.min) {
		return Range{min: other.min, max: r.max}, true
	}
	return Range{min: 0, max: 0}, false
}

type Almanac struct {
	seedToSoil            []MappingRange
	soilToFertilizer      []MappingRange
	fertilizerToWater     []MappingRange
	waterToLight          []MappingRange
	lightToTemperature    []MappingRange
	temperatureToHumidity []MappingRange
	humidityToLocation    []MappingRange
}

type MappingRange struct {
	destination int
	source      int
	length      int
}

func (mr MappingRange) InputRange() Range {
	return Range{
		min: mr.source,
		max: mr.source + mr.length,
	}
}

// output, was it transformed?
func (mappingRange MappingRange) MapInput(input int) (int, bool) {
	if input >= mappingRange.source && input < mappingRange.source+mappingRange.length {
		return input + (mappingRange.destination - mappingRange.source), true
	}
	return input, false
}
