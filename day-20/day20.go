package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/mitchellh/iochan"
)

func main() {
	f, _ := os.Open("./input.txt")
	mms := parseInput(f)
	moduleMachinePart1 := mms.CreateModuleMachine()

	totalLows := 0
	totalHighs := 0
	for i := 0; i < 1000; i++ {
		sr := moduleMachinePart1.simulate(Pulse{level: false, source: "button"})
		totalLows += sr.totalLowPulsesSent
		totalHighs += sr.totalHighPulsesSent
	}

	fmt.Printf("Simulate result after 1000 presses - lows: %d, highs: %d\n", totalLows, totalHighs)
	fmt.Printf("Multiplied %d\n", totalLows*totalHighs)

	// TODO this is too slow. Looks like we need to make some kind of optimizer.
	// moduleMachinePart2 := mms.CreateModuleMachine()
	// presses := 0
	// for {
	// 	presses++
	// 	sr := moduleMachinePart2.simulate(Pulse{level: false, source: "button"})
	// 	if sr.sentLowPulseToRx {
	// 		break
	// 	}
	// }

	// fmt.Printf("Presses before low pulse sent to rx %d\n", presses)
}

func parseInput(f *os.File) ModuleMachineSchema {
	connections := make(map[string]([]string))
	moduleTypes := make(map[string]string)
	pattern := regexp.MustCompile(`([&%]?)(\w+) -> (.*)$`)
	for line := range iochan.DelimReader(f, '\n') {
		matches := pattern.FindStringSubmatch(strings.TrimSpace(line))
		moduleType := matches[1]
		moduleName := matches[2]
		moduleConnections := strings.Split(matches[3], ", ")
		connections[moduleName] = moduleConnections
		if len(moduleType) == 0 { // Empty type indicates broadcaster
			moduleTypes[moduleName] = "broadcaster"
		} else {
			moduleTypes[moduleName] = moduleType
		}
	}
	moduleTypes["button"] = "broadcaster"
	connections["button"] = []string{"broadcaster"}
	return ModuleMachineSchema{
		modules:     moduleTypes,
		connections: connections,
	}
}
