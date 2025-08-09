package races_with_tracks

import "time"

type Season struct {
	Name  string
	Races []RaceWithNumber
}

type Race struct {
	Date  time.Time
	Race  string
	Track string
}

type RaceWithNumber struct {
	Race             Race
	RaceNumberInYear uint8
}
