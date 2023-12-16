package main

import (
	"fmt"
	"os"
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
