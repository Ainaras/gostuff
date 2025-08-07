package races_with_tracks

import (
	"reflect"
	"testing"
)

func TestGetAllRaces(t *testing.T) {
	races := GetAllRaces("f1", 2002, 2003)
	if races == nil {
		t.Error("GetAllRaces returned nil slice")
	}

	if len(races) == 0 {
		t.Error("GetAllRaces returned empty slice")
	}

	if len(races) != 33 { // Assuming
		t.Errorf("GetAllRaces returned %d races, expected 33", len(races))
	}

	races2023 := GetAllRaces("f1", 2023, 2023)
	races2024 := GetAllRaces("f1", 2024, 2024)
	racesBoth := GetAllRaces("f1", 2023, 2024)

	if len(racesBoth) != len(races2023)+len(races2024) {
		t.Errorf("Combined years should include all races from individual years, but got %d. Expected %d + %d", len(racesBoth), len(races2023), len(races2024))
	}

	for _, race := range races2023 {
		if race.Date.Year() != 2023 {
			t.Errorf("GetAllRaces(2023, 2023) returned race from year %d", race.Date.Year())
		}
	}

	for _, race := range races2024 {
		if race.Date.Year() != 2024 {
			t.Errorf("GetAllRaces(2024, 2024) returned race from year %d", race.Date.Year())
		}
	}
}

func TestGetAllRacesGroupedByYear(t *testing.T) {
	seasons := GetAllRacesGroupedByYear("f1", 2023, 2024)

	if seasons == nil {
		t.Error("GetAllRacesGroupedByYear returned nil map")
	}
	if len(seasons) == 0 {
		t.Error("GetAllRacesGroupedByYear returned empty map")
	}
	if len(seasons) != 2 {
		t.Errorf("GetAllRacesGroupedByYear returned %d seasons, expected 2", len(seasons))
	}
}

func TestGetDistinctRaceNames(t *testing.T) {
	raceNames := GetDistinctRaceNames("f1", 2020, 2025)

	if len(raceNames) == 0 {
		t.Skip("No race names found - this test requires data")
	}

	// Test that all names are unique
	seen := make(map[string]bool)
	for _, name := range raceNames {
		if seen[name] {
			t.Errorf("Race name '%s' appears more than once in distinct results", name)
		}
		seen[name] = true
	}

	names2003 := GetDistinctRaceNames("f1", 2003, 2003)

	if len(names2003) != 16 {
		t.Error("Expected 16 distinct race names for 2003")
	}
}

func TestGetDistinctTrackNames(t *testing.T) {
	trackNames := GetDistinctTrackNames("f1", 2003, 2003)

	if trackNames == nil {
		t.Error("GetDistinctTrackNames returned nil slice")
	}

	if len(trackNames) != 16 {
		t.Error("Expected 16 distinct track names for 2003")
	}

	narrowTracks := GetDistinctTrackNames("f1", 2023, 2023)
	wideTracks := GetDistinctTrackNames("f1", 2020, 2025)

	if len(wideTracks) < len(narrowTracks) {
		t.Error("Wider date range should return at least as many distinct track names")
	}
}

func TestGetRacesWithAssignedNumbers(t *testing.T) {
	numberedRaces := GetRacesWithAssignedNumbers("f1", 2023, 2024)

	if len(numberedRaces) == 0 {
		t.Skip("No races found for 2023-2024")
	}

	for i := 1; i < len(numberedRaces); i++ {
		if numberedRaces[i].Race.Date.Before(numberedRaces[i-1].Race.Date) {
			t.Error("Races with assigned numbers are not sorted by date")
			break
		}
	}

	yearCounts := make(map[int]uint8)
	for _, raceWithNumber := range numberedRaces {
		year := raceWithNumber.Race.Date.Year()
		yearCounts[year]++

		if raceWithNumber.RaceNumberInYear != yearCounts[year] {
			t.Errorf("Race in year %d has number %d, expected %d",
				year, raceWithNumber.RaceNumberInYear, yearCounts[year])
		}

		if raceWithNumber.RaceNumberInYear == 0 {
			t.Error("Race number should never be 0")
		}
	}

}

func TestGetRacesWithAssignedNumbersSingleYear(t *testing.T) {
	numberedRaces := GetRacesWithAssignedNumbers("f1", 2023, 2023)

	if len(numberedRaces) == 0 {
		t.Skip("No races found for 2023 - this test requires data")
	}

	for _, raceWithNumber := range numberedRaces {
		if raceWithNumber.Race.Date.Year() != 2023 {
			t.Errorf("Expected only 2023 races, got race from %d", raceWithNumber.Race.Date.Year())
		}
	}

	expectedNumbers := make([]uint8, len(numberedRaces))
	for i := range expectedNumbers {
		expectedNumbers[i] = uint8(i + 1)
	}

	actualNumbers := make([]uint8, len(numberedRaces))
	for i, raceWithNumber := range numberedRaces {
		actualNumbers[i] = raceWithNumber.RaceNumberInYear
	}

	if !reflect.DeepEqual(actualNumbers, expectedNumbers) {
		t.Errorf("Expected race numbers %v, got %v", expectedNumbers, actualNumbers)
	}
}

func TestEmptyResults(t *testing.T) {
	// Test with a year range that should have no results
	emptyRaces := GetAllRaces("f1", 1900, 1900)
	if len(emptyRaces) != 0 {
		t.Errorf("Expected no races for year 1900, got %d", len(emptyRaces))
	}

	emptySeasons := GetAllRacesGroupedByYear("f1", 1900, 1900)
	if len(emptySeasons) != 0 {
		t.Errorf("Expected no seasons for year 1900, got %d", len(emptySeasons))
	}

	emptyRaceNames := GetDistinctRaceNames("f1", 1900, 1900)
	if len(emptyRaceNames) != 0 {
		t.Errorf("Expected no race names for year 1900, got %d", len(emptyRaceNames))
	}

	emptyTrackNames := GetDistinctTrackNames("f1", 1900, 1900)
	if len(emptyTrackNames) != 0 {
		t.Errorf("Expected no track names for year 1900, got %d", len(emptyTrackNames))
	}

	emptyNumbered := GetRacesWithAssignedNumbers("f1", 1900, 1900)
	if len(emptyNumbered) != 0 {
		t.Errorf("Expected no numbered races for year 1900, got %d", len(emptyNumbered))
	}
}
