package main

import (
	"fmt"
	"os"
)

func main() {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		userConfigDir = "."
	}
	configuration := ReadFrom(userConfigDir + "/powerline-go/countdown.json")
	if configuration == nil {
		configuration = NewConfiguration()
	}
	powerlineSegments := CreatePowerlineSegments(configuration)
	powerlineSegmentsJson := ToPowerlineJson(powerlineSegments)
	fmt.Print(powerlineSegmentsJson)
}
