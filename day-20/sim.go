package day20

type SimTime struct {
	buttonPress int
	ticks       int // 1 "tick" is the time it takes for a gate to evaluate
}

type PulseAtTime struct {
	time  SimTime
	level bool
}

type GateBehavior struct {
	cycleLength       int
	cycleStart        int
	pulsesBeforeCycle []PulseAtTime
	pulsesInCycle     []PulseAtTime
}

// type OptimizerGate interface {
// }

// func (mm ModuleMachine) simulate(initialPulse Pulse) SimulateResult {
// 	behaviors := make(map[string]GateBehavior)

// 	behaviors[initialPulse.source] = GateBehavior{
// 		cycleLength:       1,
// 		cycleStart:        0,
// 		pulsesBeforeCycle: []PulseAtTime{},
// 		pulsesInCycle: []PulseAtTime{
// 			PulseAtTime{time: SimTime{0, 0}, level: false},
// 		},
// 	}

// 	pulses := queue.New()
// 	pulses.Enqueue(initialPulse)
// 	for pulses.Len() != 0 {
// 		pulse := pulses.Dequeue().(Pulse)
// 		desintationModuleNames := mm.connections[pulse.source]
// 		if pulse.level {
// 			highPulseCount += len(desintationModuleNames)
// 		} else {
// 			lowPulseCount += len(desintationModuleNames)
// 		}
// 		for _, destinationName := range desintationModuleNames {
// 			if destinationName == "rx" && !pulse.level {
// 				sentLowPulseToRx = true
// 			}
// 			module, exists := mm.modules[destinationName]
// 			if !exists {
// 				continue
// 			}
// 			sent, level := module.receivePulse(pulse)
// 			if sent {
// 				pulses.Enqueue(Pulse{source: destinationName, level: level})
// 			}
// 		}
// 	}
// 	return SimulateResult{
// 		totalHighPulsesSent: highPulseCount,
// 		totalLowPulsesSent:  lowPulseCount,
// 		sentLowPulseToRx:    sentLowPulseToRx,
// 	}
// }
