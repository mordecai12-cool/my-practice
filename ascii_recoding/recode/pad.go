package main

import (
	"strings"
)

func PadArtRows(rows []string, width int) []string {
	result := make([]string, len(rows))
	for i, row := range rows {
		padding := width - len(row)
		if padding > 0 {
			result[i] = row+strings.Repeat(" ", padding)
		}else {
			if padding <= 0 {
				result[i] = row
			}
		}
	}
	return result
}
//For each row, compute padding := width - len(row). If padding > 0, append strings.Repeat(" ", padding) to the row. If padding <= 0, keep the row unchanged.