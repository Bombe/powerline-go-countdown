package main

import (
	"encoding/json"
)

type PowerlineSegment struct {
	Content         string
	Color           *json.Number `json:"Foreground,omitempty"`
	BackgroundColor *json.Number `json:"Background,omitempty"`
}

func ToPowerlineJson(segments []PowerlineSegment) (output string) {
	bytes, _ := json.Marshal(segments)
	output = string(bytes)
	return
}
