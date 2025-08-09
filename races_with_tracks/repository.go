package races_with_tracks

import (
	"strconv"
)

func GetAllRaces(series string, startYear int16, endYear int16) []Race {
	var filteredRaces = getFilteredRaces(series, startYear, endYear)

	// copying the slice to avoid modifying the original slice
	out := make([]Race, len(filteredRaces))
	copy(out, filteredRaces)
	return out
}

func GetAllRacesGroupedByYear(series string, startYear int16, endYear int16) map[int]Season {
	races := GetRacesWithAssignedNumbers(series, startYear, endYear)

	var seasons = make(map[int]Season)
	var year int
	for _, race := range races {
		year = race.Race.Date.Year()
		val, ok := seasons[year]
		if ok {
			val.Races = append(val.Races, race)
			seasons[year] = val
		} else {
			s := Season{}
			s.Name = strconv.Itoa(race.Race.Date.Year())
			s.Races = []RaceWithNumber{race}
			seasons[year] = s
		}
	}
	return seasons
}

func GetDistinctRaceNames(series string, startYear int16, endYear int16) []string {
	races := getFilteredRaces(series, startYear, endYear)

	gps := make([]string, len(races))
	for i, race := range races {
		gps[i] = race.Race
	}
	return uniqueStrings(gps)
}

func GetDistinctTrackNames(series string, startYear int16, endYear int16) []string {
	races := getFilteredRaces(series, startYear, endYear)

	tracks := make([]string, len(races))
	for i, race := range races {
		tracks[i] = race.Track
	}
	return uniqueStrings(tracks)
}

func GetRacesWithAssignedNumbers(series string, startYear int16, endYear int16) []RaceWithNumber {
	races := GetAllRaces(series, startYear, endYear)

	sortedRaces := sortRacesByDate(races)

	yearCount := make(map[int]uint8)
	var result []RaceWithNumber

	for _, race := range sortedRaces {
		year := race.Date.Year()
		yearCount[year]++
		raceWithNumber := RaceWithNumber{
			Race:             race,
			RaceNumberInYear: yearCount[year],
		}
		result = append(result, raceWithNumber)
	}

	return result
}
