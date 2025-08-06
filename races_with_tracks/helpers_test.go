package races_with_tracks

import (
	"reflect"
	"testing"
	"time"
)

func TestUniqueStrings(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "empty slice",
			input:    []string{},
			expected: []string{},
		},
		{
			name:     "single element",
			input:    []string{"hello"},
			expected: []string{"hello"},
		},
		{
			name:     "no duplicates",
			input:    []string{"a", "b", "c"},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "with duplicates",
			input:    []string{"a", "b", "a", "c", "b"},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "all duplicates",
			input:    []string{"same", "same", "same"},
			expected: []string{"same"},
		},
		{
			name:     "empty strings",
			input:    []string{"", "hello", "", "world", ""},
			expected: []string{"", "hello", "world"},
		},
		{
			name:     "consecutive duplicates",
			input:    []string{"a", "a", "b", "b", "c"},
			expected: []string{"a", "b", "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := uniqueStrings(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("uniqueStrings(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestUniqueStringsNilInput(t *testing.T) {
	// Test with nil input
	var input []string = nil
	result := uniqueStrings(input)

	if result == nil {
		t.Error("uniqueStrings should return empty slice, not nil")
	}

	if len(result) != 0 {
		t.Errorf("uniqueStrings(nil) should return empty slice, got %v", result)
	}
}

func TestSortRacesByDate(t *testing.T) {
	tests := []struct {
		name     string
		input    []Race
		expected []Race
	}{
		{
			name:     "empty slice",
			input:    []Race{},
			expected: []Race{},
		},
		{
			name: "single race",
			input: []Race{
				{Date: time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC), Race: "Monaco GP", Track: "Monaco"},
			},
			expected: []Race{
				{Date: time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC), Race: "Monaco GP", Track: "Monaco"},
			},
		},
		{
			name: "already sorted races",
			input: []Race{
				{Date: time.Date(2023, 3, 5, 0, 0, 0, 0, time.UTC), Race: "Bahrain GP", Track: "Bahrain"},
				{Date: time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC), Race: "Monaco GP", Track: "Monaco"},
				{Date: time.Date(2023, 7, 10, 0, 0, 0, 0, time.UTC), Race: "British GP", Track: "Silverstone"},
			},
			expected: []Race{
				{Date: time.Date(2023, 3, 5, 0, 0, 0, 0, time.UTC), Race: "Bahrain GP", Track: "Bahrain"},
				{Date: time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC), Race: "Monaco GP", Track: "Monaco"},
				{Date: time.Date(2023, 7, 10, 0, 0, 0, 0, time.UTC), Race: "British GP", Track: "Silverstone"},
			},
		},
		{
			name: "unsorted races",
			input: []Race{
				{Date: time.Date(2023, 7, 10, 0, 0, 0, 0, time.UTC), Race: "British GP", Track: "Silverstone"},
				{Date: time.Date(2023, 3, 5, 0, 0, 0, 0, time.UTC), Race: "Bahrain GP", Track: "Bahrain"},
				{Date: time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC), Race: "Monaco GP", Track: "Monaco"},
			},
			expected: []Race{
				{Date: time.Date(2023, 3, 5, 0, 0, 0, 0, time.UTC), Race: "Bahrain GP", Track: "Bahrain"},
				{Date: time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC), Race: "Monaco GP", Track: "Monaco"},
				{Date: time.Date(2023, 7, 10, 0, 0, 0, 0, time.UTC), Race: "British GP", Track: "Silverstone"},
			},
		},
		{
			name: "reverse sorted races",
			input: []Race{
				{Date: time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC), Race: "Abu Dhabi GP", Track: "Yas Marina"},
				{Date: time.Date(2023, 6, 20, 0, 0, 0, 0, time.UTC), Race: "Canadian GP", Track: "Montreal"},
				{Date: time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC), Race: "Season Opener", Track: "Test Track"},
			},
			expected: []Race{
				{Date: time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC), Race: "Season Opener", Track: "Test Track"},
				{Date: time.Date(2023, 6, 20, 0, 0, 0, 0, time.UTC), Race: "Canadian GP", Track: "Montreal"},
				{Date: time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC), Race: "Abu Dhabi GP", Track: "Yas Marina"},
			},
		},
		{
			name: "races with same date",
			input: []Race{
				{Date: time.Date(2023, 5, 15, 14, 0, 0, 0, time.UTC), Race: "Race 2", Track: "Track B"},
				{Date: time.Date(2023, 5, 15, 10, 0, 0, 0, time.UTC), Race: "Race 1", Track: "Track A"},
				{Date: time.Date(2023, 5, 15, 18, 0, 0, 0, time.UTC), Race: "Race 3", Track: "Track C"},
			},
			expected: []Race{
				{Date: time.Date(2023, 5, 15, 10, 0, 0, 0, time.UTC), Race: "Race 1", Track: "Track A"},
				{Date: time.Date(2023, 5, 15, 14, 0, 0, 0, time.UTC), Race: "Race 2", Track: "Track B"},
				{Date: time.Date(2023, 5, 15, 18, 0, 0, 0, time.UTC), Race: "Race 3", Track: "Track C"},
			},
		},
		{
			name: "races across multiple years",
			input: []Race{
				{Date: time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC), Race: "2024 Opener", Track: "Australia"},
				{Date: time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC), Race: "2022 Finale", Track: "Abu Dhabi"},
				{Date: time.Date(2023, 6, 15, 0, 0, 0, 0, time.UTC), Race: "2023 Mid Season", Track: "Canada"},
			},
			expected: []Race{
				{Date: time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC), Race: "2022 Finale", Track: "Abu Dhabi"},
				{Date: time.Date(2023, 6, 15, 0, 0, 0, 0, time.UTC), Race: "2023 Mid Season", Track: "Canada"},
				{Date: time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC), Race: "2024 Opener", Track: "Australia"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sortRacesByDate(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("sortRacesByDate() failed for %s", tt.name)
				t.Errorf("Input: %v", tt.input)
				t.Errorf("Expected: %v", tt.expected)
				t.Errorf("Got: %v", result)
			}
		})
	}
}

func TestSortRacesByDateNilInput(t *testing.T) {
	// Test with nil input
	var input []Race = nil
	result := sortRacesByDate(input)

	// The current implementation returns the same slice (nil in this case)
	// This documents the current behavior
	if result != nil {
		t.Errorf("sortRacesByDate(nil) returned %v, expected nil", result)
	}
}
