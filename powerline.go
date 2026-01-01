package main

import "encoding/json"

type PowerlineSegment struct {
	Content string
}

func ToPowerlineJson(segments []PowerlineSegment) (output string) {
	bytes, _ := json.Marshal(segments)
	output = string(bytes)
	return
}
