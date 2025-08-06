package races_with_tracks

import "sort"

func uniqueStrings(input []string) []string {
	seen := make(map[string]bool)
	result := []string{}
	for _, v := range input {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

func sortRacesByDate(races []Race) []Race {
	sort.Slice(races, func(i, j int) bool {
		return races[i].Date.Before(races[j].Date)
	})

	return races
}
