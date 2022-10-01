package main

import (
	"fmt"
	"strings"
)

func printMenu() {
	fmt.Println("Calendar")
	fmt.Println("========")
	fmt.Println("Choose your option:")
	fmt.Println("1) Print Year")
	fmt.Println("2) Print Month")
	fmt.Println("3) Print Day of year")
}

func printMonth(month string, numberOfDays, firstDay int) {
	fmt.Println("--------------------")
	fmt.Println("       " + strings.ToUpper(month))
	fmt.Println("--------------------")
	fmt.Println(" S  M  T  W  T  F  S")
	fmt.Println("--------------------")

	firstDay = firstDay - 1
	if firstDay < 0 {
		firstDay = 6
	}

	dayOfWeekCounter := 0

	for i := 0; i < firstDay; i++ {
		fmt.Print("   ")
		dayOfWeekCounter++
	}

	for d := 0; d < numberOfDays; d++ {
		if dayOfWeekCounter > 6 {
			fmt.Println("")
			dayOfWeekCounter = 0
		}
		if (d + 1) < 10 {
			fmt.Printf(" %d ", (d + 1))
		} else {
			fmt.Printf("%d ", (d + 1))
		}
		dayOfWeekCounter++
	}

	fmt.Println("")
}

func printYear(year int) {
	for monthNumber := 1; monthNumber < 13; monthNumber++ {
		_, firstDay := getDayOfTheWeek(year, monthNumber, 1)
		monthName, monthDays := monthValues(year, monthNumber)
		printMonth(monthName, monthDays, int(firstDay))
	}
}
