package main

import "github.com/golang-collections/collections/queue"

type ModuleMachineSchema struct {
	modules     map[string]string // name -> module type ('&', '%', or 'broadcaster')
	connections map[string]([]string)
}

func (mms ModuleMachineSchema) CreateModuleMachine() ModuleMachine {
	modules := make(map[string]Module)
	for name, moduleType := range mms.modules {
		var module Module
		switch moduleType {
		case "&":
			module = &ConjunctionModule{rememberedInputs: make(map[string]bool)}
		case "%":
			module = &FlipFlopModule{}
		case "broadcaster":
			module = &BroadcasterModule{}
		default:
			panic("Invalid module type")
		}
		modules[name] = module
	}
	for source, dests := range mms.connections {
		for _, dest := range dests {
			// Initialize input connections
			_, exists := modules[dest]
			if exists {
				modules[dest].addConnection(source)
			}
		}
	}
	return ModuleMachine{
		modules:     modules,
		connections: mms.connections,
	}
}

type ModuleMachine struct {
	modules     map[string]Module
	connections map[string][]string
}

type Module interface {
	// Returns: sent a pulse? , what pulse did it send?
	receivePulse(pulse Pulse) (bool, bool)
	addConnection(name string)
}

type ConjunctionModule struct {
	rememberedInputs map[string]bool
}

func (cm *ConjunctionModule) receivePulse(pulse Pulse) (bool, bool) {
	cm.rememberedInputs[pulse.source] = pulse.level
	for _, val := range cm.rememberedInputs {
		if !val {
			return true, true
		}
	}
	return true, false
}

func (cm *ConjunctionModule) addConnection(name string) {
	cm.rememberedInputs[name] = false
}

type FlipFlopModule struct {
	memory bool
}

func (ffm *FlipFlopModule) receivePulse(pulse Pulse) (bool, bool) {
	if pulse.level {
		return false, false
	} else {
		ffm.memory = !ffm.memory
		return true, ffm.memory
	}
}
func (ffm *FlipFlopModule) addConnection(name string) {} // No-op

type BroadcasterModule struct{}

func (bm *BroadcasterModule) receivePulse(pulse Pulse) (bool, bool) {
	return true, pulse.level
}

func (bm *BroadcasterModule) addConnection(name string) {} // No-op

type Pulse struct {
	source string
	level  bool
}

type SimulateResult struct {
	totalHighPulsesSent int
	totalLowPulsesSent  int
	sentLowPulseToRx    bool
}

func (mm ModuleMachine) simulate(initialPulse Pulse) SimulateResult {
	highPulseCount := 0
	lowPulseCount := 0
	sentLowPulseToRx := false
	pulses := queue.New()
	pulses.Enqueue(initialPulse)
	for pulses.Len() != 0 {
		pulse := pulses.Dequeue().(Pulse)
		desintationModuleNames := mm.connections[pulse.source]
		if pulse.level {
			highPulseCount += len(desintationModuleNames)
		} else {
			lowPulseCount += len(desintationModuleNames)
		}
		for _, destinationName := range desintationModuleNames {
			module, exists := mm.modules[destinationName]
			if !exists {
				continue
			}
			if destinationName == "rx" {
				sentLowPulseToRx = true
			}
			sent, level := module.receivePulse(pulse)
			if sent {
				pulses.Enqueue(Pulse{source: destinationName, level: level})
			}
		}
	}
	return SimulateResult{
		totalHighPulsesSent: highPulseCount,
		totalLowPulsesSent:  lowPulseCount,
		sentLowPulseToRx:    sentLowPulseToRx,
	}
}