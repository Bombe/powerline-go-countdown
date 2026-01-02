package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"testing/synctest"
	"time"
)

func TestConfigurationWithoutDeadlinesIsTurnedIntoEmptyPowerlineSegments(t *testing.T) {
	configuration := NewConfiguration()
	powerlineSegments := CreatePowerlineSegments(configuration)
	if len(powerlineSegments) != 0 {
		t.Fatal("powerline segments must be empty")
	}
}

func TestConfigurationWithTwoDeadlineIsTurnedIntoTwoPowerlineSegments(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		configuration := NewConfiguration()
		configuration.Deadlines = append(configuration.Deadlines, Deadline{Date: "2000-03-04", Occasion: "Some Point", Symbol: "x"})
		configuration.Deadlines = append(configuration.Deadlines, Deadline{Date: "2000-04-05", Occasion: "Other Point", Symbol: "y"})
		powerlineSegments := CreatePowerlineSegments(configuration)
		if len(powerlineSegments) != 2 {
			t.Fatal("unexpected number of powerline segments:", len(powerlineSegments))
		}
		if powerlineSegments[0].Content != fmt.Sprintf("x %d", 63) {
			t.Fatal("content is", powerlineSegments[0].Content)
		}
		if powerlineSegments[1].Content != fmt.Sprintf("y %d", 95) {
			t.Fatal("content is", powerlineSegments[1].Content)
		}
	})
}

func TestPowerlineSegmentsAreSortedBySmallestDistanceFirst(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		configuration := NewConfiguration()
		configuration.Deadlines = append(configuration.Deadlines, Deadline{Date: "2001-03-04", Occasion: "Some Point", Symbol: "z"})
		configuration.Deadlines = append(configuration.Deadlines, Deadline{Date: "2000-03-04", Occasion: "Some Point", Symbol: "x"})
		configuration.Deadlines = append(configuration.Deadlines, Deadline{Date: "2000-04-05", Occasion: "Other Point", Symbol: "y"})
		powerlineSegments := CreatePowerlineSegments(configuration)
		if len(powerlineSegments) != 3 {
			t.Fatal("unexpected number of powerline segments:", len(powerlineSegments))
		}
		if powerlineSegments[0].Content != fmt.Sprintf("x %d", 63) {
			t.Fatal("content is", powerlineSegments[0].Content)
		}
		if powerlineSegments[1].Content != fmt.Sprintf("y %d", 95) {
			t.Fatal("content is", powerlineSegments[1].Content)
		}
		if powerlineSegments[2].Content != fmt.Sprintf("z %d", 428) {
			t.Fatal("content is", powerlineSegments[2].Content)
		}
	})
}

func TestDeadlinesInThePastAreIgnored(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		configuration := NewConfiguration()
		configuration.Deadlines = append(configuration.Deadlines, Deadline{Date: "2001-03-04", Occasion: "Some Point", Symbol: "z"})
		configuration.Deadlines = append(configuration.Deadlines, Deadline{Date: "1999-03-04", Occasion: "Some Point", Symbol: "x"})
		configuration.Deadlines = append(configuration.Deadlines, Deadline{Date: "2000-04-05", Occasion: "Other Point", Symbol: "y"})
		powerlineSegments := CreatePowerlineSegments(configuration)
		if len(powerlineSegments) != 2 {
			t.Fatal("unexpected number of powerline segments:", len(powerlineSegments))
		}
		if powerlineSegments[0].Content != fmt.Sprintf("y %d", 95) {
			t.Fatal("content is", powerlineSegments[0].Content)
		}
		if powerlineSegments[1].Content != fmt.Sprintf("z %d", 428) {
			t.Fatal("content is", powerlineSegments[1].Content)
		}
	})
}

func TestDeadlinesThatCanNotBeParsedAreIgnored(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		configuration := NewConfiguration()
		configuration.Deadlines = append(configuration.Deadlines, Deadline{Date: "2001-03-04", Occasion: "Some Point", Symbol: "z"})
		configuration.Deadlines = append(configuration.Deadlines, Deadline{Date: "2000-38-145", Occasion: "Some Point", Symbol: "x"})
		configuration.Deadlines = append(configuration.Deadlines, Deadline{Date: "2000-04-05", Occasion: "Other Point", Symbol: "y"})
		powerlineSegments := CreatePowerlineSegments(configuration)
		if len(powerlineSegments) != 2 {
			t.Fatal("unexpected number of powerline segments:", len(powerlineSegments))
		}
		if powerlineSegments[0].Content != fmt.Sprintf("y %d", 95) {
			t.Fatal("content is", powerlineSegments[0].Content)
		}
		if powerlineSegments[1].Content != fmt.Sprintf("z %d", 428) {
			t.Fatal("content is", powerlineSegments[1].Content)
		}
	})
}

func TestDistanceUntilTimeOnSameDateAsNowIsZero(t *testing.T) {
	verifyThatDistanceToDateIsCalculatedCorrectly(t, "2000-01-01", 0, 43200*time.Second)
}

func TestDistanceUntilTimeOnTomorrowIsOne(t *testing.T) {
	verifyThatDistanceToDateIsCalculatedCorrectly(t, "2000-01-02", 1, time.Duration(0))
}

func TestOmittedTimestampCountsUntilMidnight(t *testing.T) {
	verifyThatDistanceToDateIsCalculatedCorrectly(t, "2000-01-02", 1, time.Duration(0))
	verifyThatDistanceToDateIsCalculatedCorrectly(t, "2000-01-02", 1, 86399*time.Second)
	verifyThatDistanceToDateIsCalculatedCorrectly(t, "2000-01-02", 0, 86400*time.Second)
}

func verifyThatDistanceToDateIsCalculatedCorrectly(t *testing.T, date string, expectedDistance int, delay time.Duration) {
	synctest.Test(t, func(t *testing.T) {
		configuration := NewConfiguration()
		time.Sleep(delay)
		configuration.Deadlines = append(configuration.Deadlines, Deadline{Date: date, Occasion: "Some Point", Symbol: "x"})
		powerlineSegments := CreatePowerlineSegments(configuration)
		if len(powerlineSegments) != 1 {
			t.Fatal("unexpected number of powerline segments:", len(powerlineSegments))
		}
		if powerlineSegments[0].Content != fmt.Sprintf("x %d", expectedDistance) {
			t.Fatal("content is", powerlineSegments[0].Content)
		}
	})
}

func TestOmittedSymbolDoesNotCauseASpaceToBePrependedToNumberOfDays(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		configuration := NewConfiguration()
		configuration.Deadlines = append(configuration.Deadlines, Deadline{Date: "2000-01-02", Occasion: "Some Point"})
		powerlineSegments := CreatePowerlineSegments(configuration)
		if len(powerlineSegments) != 1 {
			t.Fatal("unexpected number of powerline segments:", len(powerlineSegments))
		}
		if powerlineSegments[0].Content != "1" {
			t.Fatal("content is", powerlineSegments[0].Content)
		}
	})
}

func TestThreeDigitHexColorIsTranslatedCorrectly(t *testing.T) {
	expectedColor := json.Number("31")
	verifyThatColorIsTranslatedCorrectly(t, "#48c", &expectedColor)
}

func TestSixDigitHexColorIsTranslatedCorrectly(t *testing.T) {
	expectedColor := json.Number("52")
	verifyThatColorIsTranslatedCorrectly(t, "#7a1b2a", &expectedColor)
}

func TestTerminalColorIsNotTranslated(t *testing.T) {
	expectedColor := json.Number("246")
	verifyThatColorIsTranslatedCorrectly(t, "246", &expectedColor)
}

func TestUnparseableColorIsReplacedByEmptyColor(t *testing.T) {
	verifyThatColorIsTranslatedCorrectly(t, "no-color", nil)
}

func verifyThatColorIsTranslatedCorrectly(t *testing.T, inputColor string, expectedColor *json.Number) {
	synctest.Test(t, func(t *testing.T) {
		configuration := NewConfiguration()
		configuration.Deadlines = append(configuration.Deadlines, Deadline{Date: "2001-01-01", Occasion: "Some Point", Symbol: "x", Color: inputColor})
		powerlineSegments := CreatePowerlineSegments(configuration)
		if len(powerlineSegments) != 1 {
			t.Fatal("unexpected number of powerline segments:", len(powerlineSegments))
		}
		if (expectedColor == nil && powerlineSegments[0].Color != nil) || (expectedColor != nil && (*powerlineSegments[0].Color != *expectedColor)) {
			t.Fatal("color is", powerlineSegments[0].Color)
		}
	})
}

func TestGrayscaleColorsAreTranslatedToDifferentScale(t *testing.T) {
	colors := make(map[string]*json.Number)
	colors["#000"] = toJson("16")
	colors["#111"] = toJson("232")
	colors["#222"] = toJson("234")
	colors["#888"] = toJson("244")
	colors["#ededed"] = toJson("254")
	colors["#f0f0f0"] = toJson("255")
	colors["#fff"] = toJson("231")
	colors["invalid"] = toJson(nil)
	for inputColor, expectedColor := range colors {
		t.Run(fmt.Sprintf("%s -> %s", inputColor, expectedColor), func(t *testing.T) {
			verifyThatColorIsTranslatedCorrectly(t, inputColor, expectedColor)
		})
	}
}

func TestGrayscaleBackgroundColorsAreTranslatedToDifferentScale(t *testing.T) {
	colors := make(map[string]*json.Number)
	colors["#000"] = toJson("16")
	colors["#111"] = toJson("232")
	colors["#222"] = toJson("234")
	colors["#888"] = toJson("244")
	colors["#ededed"] = toJson("254")
	colors["#f0f0f0"] = toJson("255")
	colors["#fff"] = toJson("231")
	colors["invalid"] = toJson(nil)
	for inputColor, expectedColor := range colors {
		t.Run(fmt.Sprintf("%s -> %s", inputColor, expectedColor), func(t *testing.T) {
			verifyThatBackgroundColorIsTranslatedCorrectly(t, inputColor, expectedColor)
		})
	}
}

func toJson(color any) *json.Number {
	if color == nil {
		return nil
	}
	jsonNumber := json.Number(color.(string))
	return &jsonNumber
}

func TestThreeDigitHexBackgroundColorIsTranslatedCorrectly(t *testing.T) {
	expectedColor := json.Number("31")
	verifyThatBackgroundColorIsTranslatedCorrectly(t, "#48c", &expectedColor)
}

func TestSixDigitHexBackgroundColorIsTranslatedCorrectly(t *testing.T) {
	expectedColor := json.Number("52")
	verifyThatBackgroundColorIsTranslatedCorrectly(t, "#7a1b2a", &expectedColor)
}

func TestTerminalBackgroundColorIsNotTranslated(t *testing.T) {
	expectedColor := json.Number("246")
	verifyThatBackgroundColorIsTranslatedCorrectly(t, "246", &expectedColor)
}

func TestUnparseableBackgroundColorIsReplacedByEmptyColor(t *testing.T) {
	verifyThatBackgroundColorIsTranslatedCorrectly(t, "no-color", nil)
}

func verifyThatBackgroundColorIsTranslatedCorrectly(t *testing.T, inputColor string, expectedColor *json.Number) {
	synctest.Test(t, func(t *testing.T) {
		configuration := NewConfiguration()
		configuration.Deadlines = append(configuration.Deadlines, Deadline{Date: "2001-01-01", Occasion: "Some Point", Symbol: "x", BackgroundColor: inputColor})
		powerlineSegments := CreatePowerlineSegments(configuration)
		if len(powerlineSegments) != 1 {
			t.Fatal("unexpected number of powerline segments:", len(powerlineSegments))
		}
		if (expectedColor == nil && powerlineSegments[0].BackgroundColor != nil) || (expectedColor != nil && (*powerlineSegments[0].BackgroundColor != *expectedColor)) {
			t.Fatal("color is", powerlineSegments[0].BackgroundColor)
		}
	})
}
