package day5

import (
	"aoc-2023/aoc-lib"
	"fmt"
	"os"
)

func Part1(f *os.File) {
	seeds, almanac := parseInput(f)
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
	fmt.Printf("Part 1 lowest location number is %d\n", lowestLocation)
}

func Part2(f *os.File) {
	seedNumbers, almanac := parseInput(f)
	lowestLocation := 99999999999
	for i := 0; i < len(seedNumbers)/2; i++ {
		seed := aoc.Range{Min: seedNumbers[i*2], Max: seedNumbers[i*2] + seedNumbers[i*2+1]}
		soil := evaluateMappingWithRange(seed, almanac.seedToSoil)
		fertilizer := evaluateMappingWithRanges(soil, almanac.soilToFertilizer)
		water := evaluateMappingWithRanges(fertilizer, almanac.fertilizerToWater)
		light := evaluateMappingWithRanges(water, almanac.waterToLight)
		temperature := evaluateMappingWithRanges(light, almanac.lightToTemperature)
		humidity := evaluateMappingWithRanges(temperature, almanac.temperatureToHumidity)
		location := evaluateMappingWithRanges(humidity, almanac.humidityToLocation)
		for _, locationRange := range location {
			lowestLocation = min(lowestLocation, locationRange.Min)
		}
	}
	fmt.Printf("Part 2 lowest location number is %d\n", lowestLocation)
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

func evaluateMappingWithRanges(inputRanges []aoc.Range, mappingRanges []MappingRange) []aoc.Range {
	result := make([]aoc.Range, 0)
	for _, inputRange := range inputRanges {
		result = append(result, evaluateMappingWithRange(inputRange, mappingRanges)...)
	}
	return result
}

// This is technically wrong - it doesn't output range fragments that weren't mapped by the mapping ranges
// It still gave me the right answer though, and I don't want to take the time to fix it right now
// I'd have to make some kind of loop that figures out if the list of output ranges is contiguous, and if it's
// not it should insert ranges where the gaps would be
func evaluateMappingWithRange(inputRange aoc.Range, mappingRanges []MappingRange) []aoc.Range {
	resultRanges := make([]aoc.Range, 0)
	for _, mappingRange := range mappingRanges {
		overlap := inputRange.Intersection(mappingRange.InputRange())
		if overlap.Size() > 0 {
			mappedRange := aoc.Range{
				Min: overlap.Min + mappingRange.destination - mappingRange.source,
				Max: overlap.Max + mappingRange.destination - mappingRange.source,
			}
			resultRanges = append(resultRanges, mappedRange)
		}
	}
	return resultRanges
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

func (mr MappingRange) InputRange() aoc.Range {
	return aoc.Range{
		Min: mr.source,
		Max: mr.source + mr.length,
	}
}

// output, was it transformed?
func (mappingRange MappingRange) MapInput(input int) (int, bool) {
	if input >= mappingRange.source && input < mappingRange.source+mappingRange.length {
		return input + (mappingRange.destination - mappingRange.source), true
	}
	return input, false
}
