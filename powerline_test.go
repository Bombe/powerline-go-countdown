package main

import "testing"

func TestPowerlineStructureIsMarshalledCorrectly(t *testing.T) {
	output := ToPowerlineJson([]PowerlineSegment{{Content: "✅ success"}, {Content: "❌ failure"}})
	if output != "[{\"Content\":\"✅ success\"},{\"Content\":\"❌ failure\"}]" {
		t.Fatal("output was", output)
	}
}

func TestColorFieldIsTranslatedAsForeground(t *testing.T) {
	output := ToPowerlineJson([]PowerlineSegment{{Content: "✅ success", Color: "123"}, {Content: "❌ failure"}})
	if output != "[{\"Content\":\"✅ success\",\"Foreground\":\"123\"},{\"Content\":\"❌ failure\"}]" {
		t.Fatal("output was", output)
	}
}

func TestColorFieldIsOmittedIfEmpty(t *testing.T) {
	output := ToPowerlineJson([]PowerlineSegment{{Content: "✅ success", Color: ""}, {Content: "❌ failure"}})
	if output != "[{\"Content\":\"✅ success\"},{\"Content\":\"❌ failure\"}]" {
		t.Fatal("output was", output)
	}
}
