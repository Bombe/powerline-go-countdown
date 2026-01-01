package main

import (
	"encoding/json"
	"testing"
)

func TestPowerlineStructureIsMarshalledCorrectly(t *testing.T) {
	output := ToPowerlineJson([]PowerlineSegment{{Content: "✅ success"}, {Content: "❌ failure"}})
	if output != "[{\"Content\":\"✅ success\"},{\"Content\":\"❌ failure\"}]" {
		t.Fatal("output was", output)
	}
}

func TestColorFieldIsTranslatedAsForeground(t *testing.T) {
	color := json.Number("123")
	output := ToPowerlineJson([]PowerlineSegment{{Content: "✅ success", Color: &color}, {Content: "❌ failure"}})
	if output != "[{\"Content\":\"✅ success\",\"Foreground\":123},{\"Content\":\"❌ failure\"}]" {
		t.Fatal("output was", output)
	}
}

func TestColorFieldIsOmittedIfEmpty(t *testing.T) {
	output := ToPowerlineJson([]PowerlineSegment{{Content: "✅ success", Color: nil}, {Content: "❌ failure"}})
	if output != "[{\"Content\":\"✅ success\"},{\"Content\":\"❌ failure\"}]" {
		t.Fatal("output was", output)
	}
}
