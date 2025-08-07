# Races With Tracks Package

A Go package for managing and querying race data with track information.

## Public Functions

### `GetAllRaces(series string, startYear int16, endYear int16) []Race`

Returns all races for a specified series within a given year range.

**Parameters:**
- `series`: The racing series to filter by
- `startYear`: The starting year (inclusive)
- `endYear`: The ending year (inclusive)

**Returns:**
- `[]Race`: A slice of Race structs containing race data

**Example:**
```go
races := GetAllRaces("f1", 2020, 2023)
```

### `GetAllRacesGroupedByYear(series string, startYear int16, endYear int16) map[int]Season`

Returns races grouped by year as seasons.

**Parameters:**
- `series`: The racing series to filter by
- `startYear`: The starting year (inclusive)
- `endYear`: The ending year (inclusive)

**Returns:**
- `map[int]Season`: A map where keys are years and values are Season structs

**Example:**
```go
seasons, err := GetAllRacesGroupedByYear("f1", 2020, 2023)
```

### `GetDistinctRaceNames(series string, startYear int16, endYear int16) []string`

Returns a list of unique race names for the specified criteria.

**Parameters:**
- `series`: The racing series to filter by
- `startYear`: The starting year (inclusive)
- `endYear`: The ending year (inclusive)

**Returns:**
- `[]string`: A slice of unique race names

**Example:**
```go
raceNames := GetDistinctRaceNames("f1", 2020, 2023)
```

### `GetDistinctTrackNames(series string, startYear int16, endYear int16) []string`

Returns a list of unique track names for the specified criteria.

**Parameters:**
- `series`: The racing series to filter by
- `startYear`: The starting year (inclusive)
- `endYear`: The ending year (inclusive)

**Returns:**
- `[]string`: A slice of unique track names

**Example:**
```go
trackNames := GetDistinctTrackNames("f1", 2020, 2023)
```

### `GetRacesWithAssignedNumbers(series string, startYear int16, endYear int16) []RaceWithNumber`

Returns races with their sequential number within each year, sorted by date.

**Parameters:**
- `series`: The racing series to filter by
- `startYear`: The starting year (inclusive)
- `endYear`: The ending year (inclusive)

**Returns:**
- `[]RaceWithNumber`: A slice of RaceWithNumber structs containing race data and their number within the year

**Example:**
```go
numberedRaces := GetRacesWithAssignedNumbers("f1", 2023, 2023)
```

## Data Types

### `Race`
```go
type Race struct {
    Date  time.Time
    Race  string
    Track string
}
```

### `Season`
```go
type Season struct {
    Name  string
    Races []Race
}
```

### `RaceWithNumber`
```go
type RaceWithNumber struct {
    Race             Race
    RaceNumberInYear uint8
}
```
