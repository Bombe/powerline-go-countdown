package main

import (
	"fmt"
	"slices"
	"time"
)

func CreatePowerlineSegments(configuration *Configuration) (powerlineSegments []PowerlineSegment) {
	powerlineSegments = []PowerlineSegment{}
	now := time.Now()
	slices.SortFunc(configuration.Deadlines, func(first, second Deadline) int {
		if first.Date == second.Date {
			return 0
		} else if first.Date < second.Date {
			return -1
		}
		return 1
	})
	for _, deadline := range configuration.Deadlines {
		date, err := time.Parse("2006-01-02 03:04:05", deadline.Date)
		if err != nil {
			date, err = time.Parse("2006-01-02", deadline.Date)
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
		powerlineSegments = append(powerlineSegments, PowerlineSegment{Content: content})
	}
	return
}
