// package main

// import (
// 	"strings"
// 	"testing"
// )

// // buildGoodBanner constructs a perfectly valid banner map in memory.
// // No file needed — all 95 printable ASCII chars, each with exactly 8 lines.
// func buildGoodBanner() map[rune][]string {
// 	banner := make(map[rune][]string)
// 	for r := rune(32); r <= 126; r++ {
// 		art := make([]string, 8)
// 		for i := range art {
// 			art[i] = string(r) + "row"
// 		}
// 		banner[r] = art
// 	}
// 	return banner
// }

// func TestValidateBanner_GoodMap(t *testing.T) {
// 	err := ValidateBanner(buildGoodBanner())
// 	if err != nil {
// 		t.Errorf("expected nil for valid banner, got: %v", err)
// 	}
// }

// func TestValidateBanner_Nil(t *testing.T) {
// 	err := ValidateBanner(nil)
// 	if err == nil {
// 		t.Error("expected error for nil banner, got nil")
// 	}
// }

// func TestValidateBanner_WrongEntryCount(t *testing.T) {
// 	banner := buildGoodBanner()
// 	delete(banner, 'A')
// 	delete(banner, 'B')
// 	err := ValidateBanner(banner)
// 	if err == nil {
// 		t.Error("expected error for banner with 93 entries, got nil")
// 	}
// }

// func TestValidateBanner_MissingSpace(t *testing.T) {
// 	banner := buildGoodBanner()
// 	delete(banner, ' ')
// 	err := ValidateBanner(banner)
// 	if err == nil {
// 		t.Error("expected error for missing space character")
// 	}
// }

// func TestValidateBanner_MissingTilde(t *testing.T) {
// 	// ~ is ASCII 126 — highest valid character
// 	banner := buildGoodBanner()
// 	delete(banner, '~')
// 	err := ValidateBanner(banner)
// 	if err == nil {
// 		t.Error("expected error for missing tilde (~) character")
// 	}
// }

// func TestValidateBanner_TooFewLines(t *testing.T) {
// 	banner := buildGoodBanner()
// 	banner['A'] = []string{"only", "six", "lines", "here", "not", "eight"}
// 	err := ValidateBanner(banner)
// 	if err == nil {
// 		t.Error("expected error when character 'A' has 6 lines instead of 8")
// 	}
// 	if err != nil && !strings.Contains(err.Error(), "A") {
// 		t.Errorf("error message should mention the offending character 'A', got: %v", err)
// 	}
// }

// func TestValidateBanner_TooManyLines(t *testing.T) {
// 	banner := buildGoodBanner()
// 	banner['Z'] = make([]string, 10)
// 	err := ValidateBanner(banner)
// 	if err == nil {
// 		t.Error("expected error when character 'Z' has 10 lines instead of 8")
// 	}
// }

// func TestValidateBanner_ZeroLinesForChar(t *testing.T) {
// 	banner := buildGoodBanner()
// 	banner['!'] = []string{}
// 	err := ValidateBanner(banner)
// 	if err == nil {
// 		t.Error("expected error when character '!' has 0 lines")
// 	}
// }

// func TestValidateBanner_EmptyMap(t *testing.T) {
// 	err := ValidateBanner(map[rune][]string{})
// 	if err == nil {
// 		t.Error("expected error for empty map, got nil")
// 	}
// }

// func TestValidateBanner_ExtraCharacterDoesNotMakeItValid(t *testing.T) {
// 	// A map with 95 correct entries + 1 extra non-ASCII rune still has the wrong count
// 	banner := buildGoodBanner()
// 	banner[rune(200)] = make([]string, 8) // add a non-printable-ASCII char
// 	err := ValidateBanner(banner)
// 	if err == nil {
// 		t.Error("expected error: 96 entries should fail the count check")
// 	}
// }

// func TestValidateBanner_ErrorMessageIsDescriptive(t *testing.T) {
// 	banner := buildGoodBanner()
// 	banner['M'] = []string{"only", "three"}
// 	err := ValidateBanner(banner)
// 	if err == nil {
// 		t.Fatal("expected non-nil error")
// 	}
// 	if len(err.Error()) < 5 {
// 		t.Errorf("error message too short to be descriptive: %q", err.Error())
// 	}
// }
// package main

// import (
// 	"reflect"
// 	"testing"
// )

// // makeArt builds a simple 8-line art slice with an identifying label.
// func makeArt(label string) []string {
// 	art := make([]string, 8)
// 	for i := range art {
// 		art[i] = label
// 	}
// 	return art
// }

// func TestMergeBanners_BaseOnlyEntries(t *testing.T) {
// 	base := map[rune][]string{'A': makeArt("base-A")}
// 	priority := map[rune][]string{}
// 	result := MergeBanners(base, priority)
// 	if !reflect.DeepEqual(result['A'], makeArt("base-A")) {
// 		t.Errorf("entry only in base must appear in result")
// 	}
// }

// func TestMergeBanners_PriorityOnlyEntries(t *testing.T) {
// 	base := map[rune][]string{}
// 	priority := map[rune][]string{'B': makeArt("priority-B")}
// 	result := MergeBanners(base, priority)
// 	if !reflect.DeepEqual(result['B'], makeArt("priority-B")) {
// 		t.Errorf("entry only in priority must appear in result")
// 	}
// }

// func TestMergeBanners_PriorityWinsOnConflict(t *testing.T) {
// 	base := map[rune][]string{'A': makeArt("base-A")}
// 	priority := map[rune][]string{'A': makeArt("priority-A")}
// 	result := MergeBanners(base, priority)
// 	if !reflect.DeepEqual(result['A'], makeArt("priority-A")) {
// 		t.Errorf("priority entry must overwrite base entry for same rune")
// 	}
// }

// func TestMergeBanners_BothContributeDistinctKeys(t *testing.T) {
// 	base := map[rune][]string{'A': makeArt("base-A")}
// 	priority := map[rune][]string{'B': makeArt("priority-B")}
// 	result := MergeBanners(base, priority)
// 	if _, ok := result['A']; !ok {
// 		t.Error("'A' from base is missing in result")
// 	}
// 	if _, ok := result['B']; !ok {
// 		t.Error("'B' from priority is missing in result")
// 	}
// }

// func TestMergeBanners_DoesNotModifyBase(t *testing.T) {
// 	base := map[rune][]string{'A': makeArt("base-A")}
// 	priority := map[rune][]string{'A': makeArt("priority-A")}
// 	MergeBanners(base, priority)
// 	if !reflect.DeepEqual(base['A'], makeArt("base-A")) {
// 		t.Error("MergeBanners must not modify the base map")
// 	}
// }

// func TestMergeBanners_DoesNotModifyPriority(t *testing.T) {
// 	base := map[rune][]string{}
// 	priority := map[rune][]string{'A': makeArt("priority-A")}
// 	MergeBanners(base, priority)
// 	if !reflect.DeepEqual(priority['A'], makeArt("priority-A")) {
// 		t.Error("MergeBanners must not modify the priority map")
// 	}
// }

// func TestMergeBanners_ResultIsNewMap(t *testing.T) {
// 	base := map[rune][]string{'A': makeArt("base-A")}
// 	priority := map[rune][]string{}
// 	result := MergeBanners(base, priority)
// 	// Mutating result must not affect base
// 	result['A'] = makeArt("mutated")
// 	if reflect.DeepEqual(base['A'], makeArt("mutated")) {
// 		t.Error("mutating the result must not affect the base map")
// 	}
// }

// func TestMergeBanners_BothEmpty(t *testing.T) {
// 	result := MergeBanners(map[rune][]string{}, map[rune][]string{})
// 	if result == nil {
// 		t.Error("result must not be nil even when both inputs are empty")
// 	}
// 	if len(result) != 0 {
// 		t.Errorf("expected empty result, got %d entries", len(result))
// 	}
// }

// func TestMergeBanners_ResultLength(t *testing.T) {
// 	base := map[rune][]string{
// 		'A': makeArt("A"),
// 		'B': makeArt("B"),
// 	}
// 	priority := map[rune][]string{
// 		'B': makeArt("B-override"),
// 		'C': makeArt("C"),
// 	}
// 	result := MergeBanners(base, priority)
// 	// Unique keys: A, B, C = 3
// 	if len(result) != 3 {
// 		t.Errorf("expected 3 entries in merged result, got %d", len(result))
// 	}
// }

// func TestMergeBanners_NilBaseActsAsEmpty(t *testing.T) {
// 	priority := map[rune][]string{'A': makeArt("A")}
// 	// Should not panic on nil base
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Errorf("MergeBanners panicked on nil base: %v", r)
// 		}
// 	}()
// 	result := MergeBanners(nil, priority)
// 	if result == nil {
// 		t.Error("result must not be nil")
// 	}
// }
// package main

// import (
// 	"strings"
// 	"testing"
// )

// func TestPadArtRows_PadsShortRows(t *testing.T) {
// 	input := []string{"hi", "there", "a", "b", "c", "d", "e", "f"}
// 	result := PadArtRows(input, 8)
// 	for i, row := range result {
// 		if len(row) != 8 {
// 			t.Errorf("row %d: expected length 8, got %d (%q)", i, len(row), row)
// 		}
// 	}
// }

// func TestPadArtRows_PaddingIsSpaces(t *testing.T) {
// 	input := []string{"ab", "", "", "", "", "", "", ""}
// 	result := PadArtRows(input, 5)
// 	if result[0] != "ab   " {
// 		t.Errorf("expected \"ab   \", got %q", result[0])
// 	}
// 	// All padding characters must be spaces
// 	if strings.TrimLeft(result[0][2:], " ") != "" {
// 		t.Errorf("padding must be spaces only, got %q", result[0][2:])
// 	}
// }

// func TestPadArtRows_DoesNotTruncate(t *testing.T) {
// 	// Row wider than target width must come back unchanged
// 	input := []string{"hello world", "", "", "", "", "", "", ""}
// 	result := PadArtRows(input, 5)
// 	if result[0] != "hello world" {
// 		t.Errorf("wider row must not be truncated: got %q", result[0])
// 	}
// }

// func TestPadArtRows_ExactWidthUnchanged(t *testing.T) {
// 	input := []string{"abcd", "", "", "", "", "", "", ""}
// 	result := PadArtRows(input, 4)
// 	if result[0] != "abcd" {
// 		t.Errorf("row at exact width must be unchanged: got %q", result[0])
// 	}
// }

// func TestPadArtRows_EmptyRowPaddedToWidth(t *testing.T) {
// 	input := []string{"", "", "", "", "", "", "", ""}
// 	result := PadArtRows(input, 4)
// 	for i, row := range result {
// 		if row != "    " {
// 			t.Errorf("row %d: empty row should pad to 4 spaces, got %q", i, row)
// 		}
// 	}
// }

// func TestPadArtRows_LeadingSpacesPreserved(t *testing.T) {
// 	// Leading spaces are art — must not be touched
// 	input := []string{"  _  ", "", "", "", "", "", "", ""}
// 	result := PadArtRows(input, 8)
// 	if !strings.HasPrefix(result[0], "  _  ") {
// 		t.Errorf("leading spaces removed: got %q", result[0])
// 	}
// }

// func TestPadArtRows_LengthAlwaysPreserved(t *testing.T) {
// 	input := []string{"a", "bb", "ccc", "dddd", "eeeee", "ffffff", "ggggggg", "hhhhhhhh"}
// 	result := PadArtRows(input, 6)
// 	if len(result) != len(input) {
// 		t.Errorf("number of rows changed: got %d, want %d", len(result), len(input))
// 	}
// }

// func TestPadArtRows_DoesNotModifyInput(t *testing.T) {
// 	input := []string{"hi", "a", "b", "c", "d", "e", "f", "g"}
// 	originals := make([]string, len(input))
// 	copy(originals, input)
// 	PadArtRows(input, 10)
// 	for i := range input {
// 		if input[i] != originals[i] {
// 			t.Errorf("row %d: input was mutated — must not modify input slice", i)
// 		}
// 	}
// }

// func TestPadArtRows_ZeroWidthDoesNotPanic(t *testing.T) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Errorf("PadArtRows panicked with width=0: %v", r)
// 		}
// 	}()
// 	input := []string{"hi", "", "", "", "", "", "", ""}
// 	result := PadArtRows(input, 0)
// 	if len(result) != 8 {
// 		t.Errorf("expected 8 rows returned, got %d", len(result))
// 	}
// }

// func TestPadArtRows_NegativeWidthDoesNotPanic(t *testing.T) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Errorf("PadArtRows panicked with width=-5: %v", r)
// 		}
// 	}()
// 	input := []string{"hi", "", "", "", "", "", "", ""}
// 	PadArtRows(input, -5)
// }

// func TestPadArtRows_AllRowsSameWidthAfterPadding(t *testing.T) {
// 	input := []string{"a", "bb", "ccc", "d", "ee", "f", "ggg", "hh"}
// 	result := PadArtRows(input, 10)
// 	for i, row := range result {
// 		if len(row) != 10 {
// 			t.Errorf("row %d: expected length 10 after padding, got %d (%q)", i, len(row), row)
// 		}
// 	}
// }
package main

import (
	"strings"
	"testing"
)

func TestTrimArtRows_LengthUnchanged(t *testing.T) {
	input := []string{"a  ", "b", "c   ", "d", "e  ", "f", "g   ", "h"}
	result := TrimArtRight(input)
	if len(result) != len(input) {
		t.Errorf("expected %d rows, got %d", len(input), len(result))
	}
}

func TestTrimArtRows_AllEmptyRows(t *testing.T) {
	input := []string{"", "", "", "", "", "", "", ""}
	result := TrimArtRight(input)
	for i, row := range result {
		if row != "" {
			t.Errorf("row %d: expected empty string, got %q", i, row)
		}
	}
}

func TestTrimArtRows_AllSpaceRows(t *testing.T) {
	input := []string{"   ", "  ", " ", "    ", "     ", "  ", "   ", " "}
	result := TrimArtRight(input)
	for i, row := range result {
		if row != "" {
			t.Errorf("row %d: all-space row should become empty string, got %q", i, row)
		}
	}
}

func TestTrimArtRows_NoTrailingSpaces(t *testing.T) {
	// Nothing to trim — rows should come back identical
	input := []string{"_", "| |", "|_|", "", " _", "| |", "|_|", ""}
	result := TrimArtRight(input)
	for i := range input {
		if result[i] != input[i] {
			t.Errorf("row %d: no trailing spaces, must be unchanged. got %q, want %q",
				i, result[i], input[i])
		}
	}
}

func TestTrimArtRows_DoesNotModifyInput(t *testing.T) {
	input := []string{"hi   ", "there  ", "a   ", "b  ", "c ", "d  ", "e  ", "f   "}
	originals := make([]string, len(input))
	copy(originals, input)
	TrimArtRight(input)
	for i := range input {
		if input[i] != originals[i] {
			t.Errorf("row %d: input was modified — must return new slice. got %q", i, input[i])
		}
	}
}

func TestTrimArtRows_ReturnsNewSlice(t *testing.T) {
	input := []string{"a  ", "b  ", "c  ", "d  ", "e  ", "f  ", "g  ", "h  "}
	result := TrimArtRight(input)
	// Mutate result — original must be unchanged
	result[0] = "MUTATED"
	if input[0] == "MUTATED" {
		t.Error("TrimArtRows must return a new slice, not a reference to the input")
	}
}

func TestTrimArtRows_MidRowSpacesPreserved(t *testing.T) {
	// Internal spaces in art must never be touched
	input := []string{"| _ |   ", "| | |  ", "|___|", "", "", "", "", ""}
	result := TrimArtRight(input)
	if !strings.Contains(result[0], "| _ |") {
		t.Errorf("row 0: internal spaces removed — must be preserved. got %q", result[0])
	}
	if !strings.Contains(result[1], "| | |") {
		t.Errorf("row 1: internal spaces removed — must be preserved. got %q", result[1])
	}
}

func TestTrimArtRows_EmptySlice(t *testing.T) {
	result := TrimArtRight([]string{})
	if result == nil {
		t.Error("must not return nil for empty input slice")
	}
	if len(result) != 0 {
		t.Errorf("expected empty slice, got length %d", len(result))
	}
}