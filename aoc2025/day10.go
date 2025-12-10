package aoc2025

import (
	"strconv"
	"strings"
)

type Machine struct {
	InitLights []bool
	InitJoltage []int
	Lights  []bool
	Joltage []int
	Buttons [][]int
	Presses int
	FuncSwitch bool
}

func (m *Machine) PressButton(index int) {
	m.Presses += 1
	if !m.FuncSwitch {
		for _, lightIndex := range m.Buttons[index] {
			m.Lights[lightIndex] = !m.Lights[lightIndex]
		}
	} else {
		for _, lightIndex := range m.Buttons[index] {
			m.Joltage[lightIndex] += 1
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
	for i, joltage := range m.Joltage {
		if joltage > m.InitJoltage[i] {
			return -1
		} else if joltage < m.InitJoltage[i] {
			matching = false
		}
	}
	if matching {
		return 1
	}
	return 0
}

func copyMachine(machine Machine) Machine {
	newLights := make([]bool, len(machine.Lights))
	copy(newLights, machine.Lights)
	newJoltage := make([]int, len(machine.Joltage))
	copy(newJoltage, machine.Joltage)
	newMachine := Machine{
		InitLights: machine.InitLights,
		InitJoltage: machine.InitJoltage,
		Lights:     newLights,
		Joltage:    newJoltage,
		Buttons:    machine.Buttons,
		Presses:    machine.Presses,
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
		sum += findLightButtonPresses(machine)
	}
	return strconv.Itoa(sum)
}

func findLightButtonPresses(machine Machine) int {
	machines := []Machine{machine}
	for true {
		newMachines := []Machine{}
		for _, m := range machines {
			for i := range m.Buttons {
				newMachine := copyMachine(m)
				newMachine.PressButton(i)
				if newMachine.CheckLights() {
					return newMachine.Presses
				}
				newMachines = append(newMachines, newMachine)
			}
		}
		machines = newMachines
	}
	return -1
}

func day10Part2(data []string) string {
	sum := 0
	for _, machine := range data {
		if machine == "" { continue }
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
			InitJoltage: joltages,
			Joltage:     make([]int, len(joltages)),
			Buttons:    buttons,
			FuncSwitch: true,
		}
		sum += findJoltageButtonPresses(machine)
	}
	return strconv.Itoa(sum)
}
func findJoltageButtonPresses(machine Machine) int {
	machines := []Machine{machine}
	for true {
		newMachines := []Machine{}
		for _, m := range machines {
			for i := range m.Buttons {
				newMachine := copyMachine(m)
				newMachine.PressButton(i)
				result := newMachine.CheckJoltage()
				if result == 1 {
					return newMachine.Presses
				} else if result == 0 {
					newMachines = append(newMachines, newMachine)
				}
			}
		}
		machines = newMachines
	}
	return -1
}
