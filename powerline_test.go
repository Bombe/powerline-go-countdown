package main

import "testing"

func TestPowerlineStructureIsMarshalledCorrectly(t *testing.T) {
	output := ToPowerlineJson([]PowerlineSegment{{Content: "✅ success"}, {Content: "❌ failure"}})
	if output != "[{\"Content\":\"✅ success\"},{\"Content\":\"❌ failure\"}]" {
		t.Fatal("output was", output)
	}
}
