package main

import (
	"encoding/json"
)

type PowerlineSegment struct {
	Content string
	Color   string `json:"Foreground,omitempty"`
}

func ToPowerlineJson(segments []PowerlineSegment) (output string) {
	bytes, _ := json.Marshal(segments)
	output = string(bytes)
	return
}
