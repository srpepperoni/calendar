package main

import (
	"fmt"
	"math"
)

func main() {
	printMenu()
	option := getParametersForSelectedOption()

	switch option {
	case "1":
		year := getYear()
		printYear(year)
	case "2":
		year, month := getParametersForMonthOption()
		_, firstDay := getDayOfTheWeek(year, month, 1)
		monthName, monthDays := monthValues(year, month)
		printMonth(monthName, monthDays, int(firstDay))
	case "3":
		year, month, day := getParametersForDayOfWeekOption()
		dayOfWeek, _ := getDayOfTheWeek(year, month, day)
		fmt.Println(dayOfWeek)
	}
}

func getParametersForMonthOption() (int, int) {
	year := getYear()
	month := getMont()

	return year, month
}

func getParametersForDayOfWeekOption() (int, int, int) {
	year := getYear()
	month := getMont()
	day := getDay(year, month)

	return year, month, day
}

func getDayOfTheWeek(year, month, day int) (string, float64) {
	if month == 1 || month == 2 {
		month = month + 12
		year = year - 1
	}

	N := day + 2*month + (3 * (month + 1) / 5) + year + (year / 4) - (year / 100) + (year / 400) + 2
	mod := math.Mod(float64(N), 7)

	switch mod {
	case 0:
		return "Saturday", mod
	case 1:
		return "Sunday", mod
	case 2:
		return "Monday", mod
	case 3:
		return "Tuesday", mod
	case 4:
		return "Wednesday", mod
	case 5:
		return "Thursday", mod
	case 6:
		return "Friday", mod
	}

	return "", -1
}
