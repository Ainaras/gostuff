package races_with_tracks

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/Ainaras/gostuff/dataloader"
)

var (
	racesBySeries map[string][]Race
	oncePerSeries map[string]*sync.Once
)

func getRaces(series string) []Race {
	if racesBySeries == nil {
		racesBySeries = make(map[string][]Race)
	}
	if oncePerSeries == nil {
		oncePerSeries = make(map[string]*sync.Once)
	}

	once, exists := oncePerSeries[series]
	if !exists {
		once = &sync.Once{}
		oncePerSeries[series] = once
	}

	once.Do(func() {
		races := []Race{}
		data := dataloader.LoadData(fmt.Sprintf("%s/races_with_tracks.txt", series))
		lines := strings.Split(data, "\n")

		for _, line := range lines {
			columns := strings.Split(strings.TrimSpace(line), "\t")
			if len(columns) != 3 {
				continue
			}

			date, err := time.Parse("2006-01-02", columns[0])
			if err != nil {
				continue
			}

			match := Race{
				Date:  date,
				Race:  columns[1],
				Track: columns[2],
			}
			races = append(races, match)
		}

		racesBySeries[series] = races
	})

	return racesBySeries[series]
}

func getFilteredRaces(series string, startYear int16, endYear int16) []Race {
	races := getRaces(series)

	var filteredRaces []Race
	for _, race := range races {
		if race.Date.Year() >= int(startYear) && race.Date.Year() <= int(endYear) {
			filteredRaces = append(filteredRaces, race)
		}
	}
	return filteredRaces
}
