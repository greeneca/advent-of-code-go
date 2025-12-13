package aoc2025

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/stat/combin"
)

type Machine struct {
	InitLights []bool
	Lights  []bool
	Joltage []int
	Buttons [][]int
	ButtonsPresed  []int
	FuncSwitch bool
}

func (m *Machine) PressButton(index int) {
	m.ButtonsPresed = append(m.ButtonsPresed, index)
	if !m.FuncSwitch {
		for _, lightIndex := range m.Buttons[index] {
			m.Lights[lightIndex] = !m.Lights[lightIndex]
		}
	} else {
		for _, lightIndex := range m.Buttons[index] {
			m.Joltage[lightIndex] -= 1
		}
	}
}

func (m *Machine) CheckLights() bool {
	for i, light := range m.Lights {
		if light != m.InitLights[i] {
			return false
		}
	}
	return true
}

func (m *Machine) CheckJoltage() int {
	matching := true
	for _, joltage := range m.Joltage {
		if joltage < 0 {
			return -1
		} else if joltage > 0 {
			matching = false
		}
	}
	if matching {
		return 1
	}
	return 0
}

func (m *Machine) StateKey() string {
	stateKey := ""
	for _, light := range m.Lights {
		if light {
			stateKey += "#"
		} else {
			stateKey += "."
		}
	}
	return stateKey
}

func copyMachine(machine Machine) Machine {
	newLights := make([]bool, len(machine.Lights))
	copy(newLights, machine.Lights)
	newJoltage := make([]int, len(machine.Joltage))
	copy(newJoltage, machine.Joltage)
	newButtonsPresed := make([]int, len(machine.ButtonsPresed))
	copy(newButtonsPresed, machine.ButtonsPresed)
	newMachine := Machine{
		InitLights: machine.InitLights,
		Lights:     newLights,
		Joltage:    newJoltage,
		Buttons:    machine.Buttons,
		ButtonsPresed: newButtonsPresed,
		FuncSwitch: machine.FuncSwitch,
	}
	return newMachine
}


func day10Part1(data []string) string {
	sum := 0
	for _, machine := range data {
		if machine == "" { continue }
		parts := strings.Split(machine, " ")
		lights := []bool{}
		for _, char := range parts[0][1:len(parts[0])-1] {
			if char == '#' {
				lights = append(lights, true)
			} else {
				lights = append(lights, false)
			}
		}
		buttons := [][]int{}
		for _, button := range parts[1:len(parts)-1] {
			buttonindexes := strings.Split(button[1:len(button)-1], ",")


			indices := []int{}
			for _, indexStr := range buttonindexes {
				index, _ := strconv.Atoi(indexStr)
				indices = append(indices, index)
			}
			buttons = append(buttons, indices)
		}
		machine := Machine{
			InitLights: lights,
			Lights:     make([]bool, len(lights)),
			Buttons:    buttons,
		}
		sum += len(findLightButtonPresses(machine, false)[0])
	}
	return strconv.Itoa(sum)
}

func findLightButtonPresses(machine Machine, all bool) [][]int {
	buttonCombos := [][]int{}
	for i := range machine.Buttons {
		buttons := combin.Combinations(len(machine.Buttons), i+1)
		for _, buttonSet := range buttons {
			newMachine := copyMachine(machine)
			for _, buttonIndex := range buttonSet {
				newMachine.PressButton(buttonIndex)
			}
			if newMachine.CheckLights() {
				if !all {
					return [][]int{buttonSet}
				}
				buttonCombos = append(buttonCombos, buttonSet)
			}
		}
	}
	return buttonCombos
}

func day10Part2(data []string) string {
	sum := 0
	for _, machine := range data {
		if machine == ""  || machine[0:2] == "//" { continue }
		parts := strings.Split(machine, " ")
		joltages := []int{}
		joltage := strings.SplitSeq(parts[len(parts)-1][1:len(parts[len(parts)-1])-1], ",")
		for val := range joltage {
			intVal, _ := strconv.Atoi(val)
			joltages = append(joltages, intVal)
		}
		buttons := [][]int{}
		for _, button := range parts[1:len(parts)-1] {
			buttonindexes := strings.Split(button[1:len(button)-1], ",")
			indices := []int{}
			for _, indexStr := range buttonindexes {
				index, _ := strconv.Atoi(indexStr)
				indices = append(indices, index)
			}
			buttons = append(buttons, indices)
		}
		machine := Machine{
			InitLights: make([]bool, len(joltages)),
			Lights: make([]bool, len(joltages)),
			Joltage: joltages,
			Buttons:    buttons,
		}
		cache := map[string]int{}
		value := findJoltageButtonPresses(machine, &cache)
		fmt.Println("Value for machine:", value)
		sum += value
	}
	return strconv.Itoa(sum)
}
func findJoltageButtonPresses(machine Machine, cache *map[string]int) int {
	key := cacheKey(machine)
	println("Finding joltage presses for machine with key:", key)
	check := machine.CheckJoltage()
	if check == 1 {
		return 0
	}
	if check == -1 {
		return -1
	}
	if val, exists := (*cache)[key]; exists {
		return val
	}
	if allEven(machine.Joltage) {
		for i := range machine.Joltage {
			machine.Joltage[i] /= 2
		}
		value := 2*findJoltageButtonPresses(copyMachine(machine), cache)
		if value < 0 {
			return -1
		}
		(*cache)[key] = value
		return value
	}
	for i, joltage := range machine.Joltage {
		if joltage%2 != 0 {
			machine.InitLights[i] = true
		} else {
			machine.InitLights[i] = false
		}
	}
	buttonSets := findLightButtonPresses(copyMachine(machine), true)
	min := math.MaxInt32
	for _, buttons := range buttonSets {
		newMachine := copyMachine(machine)
		count := len(buttons)
		newMachine.FuncSwitch = true
		for _, button := range buttons {
			newMachine.PressButton(button)
		}
		newMachine.FuncSwitch = false
		presses := findJoltageButtonPresses(newMachine, cache)
		if presses >= 0  {
			count += presses
			if count < min {
				min = count
			}
		}
	}
	if min != math.MaxInt32 {
		(*cache)[key] = min
		return min
	}
	return -1
}

func cacheKey(machine Machine) string {
	key := ""
	for _, j := range machine.Joltage {
		key += strconv.Itoa(j) + ","
	}
	return key
}

func allEven(joltages []int) bool {
	for _, joltage := range joltages {
		if joltage%2 != 0 {
			return false
		}
	}
	return true
}
