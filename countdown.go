package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"slices"
	"strconv"
	"time"
)

func terminalColorGradient(color uint8) uint8 {
	if color < 95 {
		return 0
	} else if color < 135 {
		return 1
	} else if color < 175 {
		return 2
	} else if color < 215 {
		return 3
	} else if color < 255 {
		return 4
	}
	return 5
}

func terminalColorFor(red, green, blue uint8) uint8 {
	return 16 + 36*terminalColorGradient(red) + 6*terminalColorGradient(green) + terminalColorGradient(blue)
}

func convertColorToTerminalColor(color string) *json.Number {
	if len(color) == 4 && color[0] == '#' {
		r, _ := hex.DecodeString("0" + color[1:2])
		g, _ := hex.DecodeString("0" + color[2:3])
		b, _ := hex.DecodeString("0" + color[3:4])
		number := json.Number(strconv.Itoa(int(terminalColorFor(r[0]*17, g[0]*17, b[0]*17))))
		return &number
	} else if len(color) == 7 && color[0] == '#' {
		r, _ := hex.DecodeString(color[1:3])
		g, _ := hex.DecodeString(color[3:5])
		b, _ := hex.DecodeString(color[5:7])
		number := json.Number(strconv.Itoa(int(terminalColorFor(r[0], g[0], b[0]))))
		return &number
	}
	_, err := strconv.Atoi(color)
	if err != nil {
		return nil
	}
	jsonNumber := json.Number(color)
	return &jsonNumber
}

func CreatePowerlineSegments(configuration *Configuration) (powerlineSegments []PowerlineSegment) {
	powerlineSegments = []PowerlineSegment{}
	now := time.Now()
	now = time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), time.UTC)
	slices.SortFunc(configuration.Deadlines, func(first, second Deadline) int {
		if first.Date == second.Date {
			return 0
		} else if first.Date < second.Date {
			return -1
		}
		return 1
	})
	for _, deadline := range configuration.Deadlines {
		date, err := time.ParseInLocation("2006-01-02 03:04:05", deadline.Date, time.UTC)
		if err != nil {
			date, err = time.ParseInLocation("2006-01-02", deadline.Date, time.UTC)
		}
		if err != nil {
			/* skip this deadline. */
			continue
		}
		distance := date.UnixMilli() - now.UnixMilli()
		if distance < 0 {
			continue
		}
		content := fmt.Sprintf("%s %d", deadline.Symbol, distance/86400000)
		powerlineSegments = append(powerlineSegments, PowerlineSegment{Content: content, Color: convertColorToTerminalColor(deadline.Color)})
	}
	return
}
