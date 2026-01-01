package main

import (
	"encoding/json"
	"io"
	"os"
)

type Deadline struct {
	Date     string
	Occasion string
	Symbol   string
	Color    string
}

type Configuration struct {
	Deadlines []Deadline
}

func NewConfiguration() (configuration *Configuration) {
	configuration = &Configuration{
		Deadlines: []Deadline{},
	}
	return
}

func ReadFrom(path string) (configuration *Configuration) {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()
	buffer, err := io.ReadAll(file)
	if err != nil {
		return nil
	}
	configuration = NewConfiguration()
	_ = json.Unmarshal(buffer, configuration)
	return
}
