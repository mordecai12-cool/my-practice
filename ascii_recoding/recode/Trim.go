package main

import (
	"strings"
)

func TrimArtRight(row []string) []string {
	result := make([]string, len(row))
		for i, r := range row{
		    result[i] = strings.TrimRight(r, " ")
		 }
		return result
}