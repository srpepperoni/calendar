package main

import (
	"math"
	"runtime"
	"strings"
)

func isValidYear(year int) bool {
	if year > 0 && year < 3000 {
		return true
	}
	return false
}

func isValidMonth(month int) bool {
	if month > 0 && month < 13 {
		return true
	}
	return false
}

func isValidDay(year, month, day int) bool {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		if day > 0 && day < 32 {
			return true
		}
	case 4, 6, 9, 11:
		if day > 0 && day < 31 {
			return true
		}
	case 2:
		mod := math.Mod(float64(year), 4)
		if mod == 0 && day > 0 && day < 30 {
			return true
		} else if day > 0 && day < 29 {
			return true
		}
	}
	return false
}

func monthValues(year, monthNumber int) (string, int) {
	mod := math.Mod(float64(year), 4)
	switch monthNumber {
	case 1:
		return "January", 31
	case 2:
		if mod == 0 {
			return "February", 29
		}
		return "February", 28
	case 3:
		return "March", 31
	case 4:
		return "April", 30
	case 5:
		return "May", 31
	case 6:
		return "June", 30
	case 7:
		return "July", 31
	case 8:
		return "August", 31
	case 9:
		return "September", 30
	case 10:
		return "October", 31
	case 11:
		return "November", 30
	case 12:
		return "December", 31
	}

	return "", 0
}

func cleanStdInput(input string) string {
	var cleanInput string

	if runtime.GOOS == "windows" {
		cleanInput = strings.TrimRight(input, "\r\n")
	} else {
		cleanInput = strings.TrimRight(input, "\n")
	}
	return cleanInput
}
