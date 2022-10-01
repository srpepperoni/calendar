package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getParametersForSelectedOption() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter option: ")
	option, _ := reader.ReadString('\n')
	option = cleanStdInput(option)
	return option
}

func getYear() int {
	validYear := false
	var yearNumber int

	for !validYear {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter year: ")
		year, _ := reader.ReadString('\n')
		year = cleanStdInput(year)
		yearNumber, _ = strconv.Atoi(year)
		validYear = isValidYear(yearNumber)
		if !validYear {
			fmt.Print("Invalid Year! Value must be between 1 and 2999")
		}
	}

	return yearNumber
}

func getMont() int {
	validMonth := false
	var monthNumber int

	for !validMonth {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter month: ")
		month, _ := reader.ReadString('\n')
		month = cleanStdInput(month)
		monthNumber, _ = strconv.Atoi(month)
		validMonth = isValidMonth(monthNumber)
		if !validMonth {
			fmt.Print("Invalid Month! Value must be between 1 and 12")
		}
	}

	return monthNumber
}

func getDay(year, month int) int {
	validDay := false
	var dayNumber int

	for !validDay {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter month: ")
		day, _ := reader.ReadString('\n')
		day = cleanStdInput(day)
		dayNumber, _ = strconv.Atoi(day)
		validDay = isValidDay(year, month, dayNumber)
		if !validDay {
			fmt.Print("Invalid Day! Value must be between 1 and (28,29,30 or 31)")
		}
	}

	return dayNumber
}
