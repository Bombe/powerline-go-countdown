package main

import (
	"io"
	"os"
	"reflect"
	"testing"
)

func TestCreatingNewConfigurationReturnsDefaultConfiguration(t *testing.T) {
	var configuration = NewConfiguration()
	if configuration == nil {
		t.Fatal("new configuration must not be nil")
	}
	if len(configuration.Deadlines) != 0 {
		t.Fatal("new configuration must not have deadlines")
	}
}

func TestReadingConfigurationFromNonExistingFileReturnsNil(t *testing.T) {
	var configuration = ReadFrom("/not/existing")
	if configuration != nil {
		t.Fatal("reading from non-existing file must return nil")
	}
}

func TestReadingConfigurationFromEmptyFileReturnsNewConfiguration(t *testing.T) {
	tempFile, err := os.CreateTemp("", "configuration-*.json")
	if err != nil {
		t.Fatal("could not create temp file")
	}
	defer os.Remove(tempFile.Name())
	tempFile.Close()
	configuration := ReadFrom(tempFile.Name())
	if configuration == nil {
		t.Fatal("configuration must not be nil")
	}
	if len(configuration.Deadlines) != 0 {
		t.Fatal("configuration must not have deadlines")
	}
}

func TestConfigurationCanBeUnmarshalledCorrectly(t *testing.T) {
	tempFile, err := os.CreateTemp("", "configuration-*.json")
	if err != nil {
		t.Fatal("could not create temp file")
	}
	defer os.Remove(tempFile.Name())
	_, err = io.WriteString(tempFile, "{\"deadlines\": [{\"date\": \"2026-01-01 00:00:00\", \"occasion\": \"New Year 2026\", \"symbol\": \"🎆\", \"color\": \"#fff\"},{\"date\": \"2026-06-07 00:00:00\",\"occasion\": \"June 7th\",\"symbol\": \"🥂\"}]}")
	if err != nil {
		t.Fatal("could not write to temp file")
	}
	tempFile.Close()
	configuration := ReadFrom(tempFile.Name())
	if configuration == nil {
		t.Fatal("configuration must not be nil")
	}
	wantedConfiguration := &Configuration{
		Deadlines: []Deadline{
			{
				Date:     "2026-01-01 00:00:00",
				Occasion: "New Year 2026",
				Symbol:   "🎆",
				Color:    "#fff",
			},
			{
				Date:     "2026-06-07 00:00:00",
				Occasion: "June 7th",
				Symbol:   "🥂",
				Color:    "",
			},
		},
	}
	if !reflect.DeepEqual(configuration, wantedConfiguration) {
		t.Fatal("configuration must be", wantedConfiguration, "but was", configuration)
	}
}
