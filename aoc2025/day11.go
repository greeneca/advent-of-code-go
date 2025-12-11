package aoc2025

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Device struct {
	name string
	conn []*Device
}

func day11Part1(data []string) string {
	devices := map[string]*Device{}
	for _, line := range data {
		if line == "" { continue }
		deviceStrs := strings.Split(line, " ")
		curDeviceName := deviceStrs[0][:len(deviceStrs[0])-1]
		deviceStrs = deviceStrs[1:]
		curDevice, exists := devices[curDeviceName]
		if !exists {
			curDevice = &Device{name: curDeviceName}
			devices[curDeviceName] = curDevice
		}
		for _, connName := range deviceStrs {
			connDevice, exists := devices[connName]
			if !exists {
				connDevice = &Device{name: connName}
				devices[connName] = connDevice
			}
			curDevice.conn = append(curDevice.conn, connDevice)
		}
	}
	startDevice := devices["you"]
	count := traverse(startDevice)
	return strconv.Itoa(count)
}

func traverse(device *Device) int {
	if device.name == "out" {
		return 1
	}
	count := 0
	for _, conn := range device.conn {
		count += traverse(conn)
	}
	return count
}

func day11Part2(data []string) string {
	devices := map[string]*Device{}
	for _, line := range data {
		if line == "" { continue }
		deviceStrs := strings.Split(line, " ")
		curDeviceName := deviceStrs[0][:len(deviceStrs[0])-1]
		deviceStrs = deviceStrs[1:]
		curDevice, exists := devices[curDeviceName]
		if !exists {
			curDevice = &Device{name: curDeviceName}
			devices[curDeviceName] = curDevice
		}
		for _, connName := range deviceStrs {
			connDevice, exists := devices[connName]
			if !exists {
				connDevice = &Device{name: connName}
				devices[connName] = connDevice
			}
			curDevice.conn = append(curDevice.conn, connDevice)
		}
	}
	startDevice := devices["svr"]
	cache := map[string]int{}
	count := traverseSVR(startDevice, []string{}, &cache)
	return strconv.Itoa(count)
}

func traverseSVR(device *Device, visited []string, cache *map[string]int) int {
	key := fmt.Sprintf("%s|%t|%t", device.name, slices.Contains(visited, "fft"), slices.Contains(visited, "dac"))
	if val, exists := (*cache)[key]; exists {
		return val
	}
	if device.name == "out" {
		if slices.Contains(visited, "fft") && slices.Contains(visited, "dac") {
			return 1
		}
		return 0
	}
	if slices.Contains(visited, device.name) {
		return 0
	}
	count := 0
	for _, conn := range device.conn {
		count += traverseSVR(conn, append(visited, device.name), cache)
	}
	(*cache)[key] = count
	return count
}

