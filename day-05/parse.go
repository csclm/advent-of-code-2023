package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/mitchellh/iochan"
)

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
