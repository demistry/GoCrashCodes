package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	result := timeConversion("07:05:50PM")
	fmt.Printf(result)
}

func timeConversion(twelveHourFormat string) string {
	stringArr := strings.Split(twelveHourFormat, ":")
	hourInt, _ := strconv.Atoi(stringArr[0])
	minuteStr := stringArr[1]
	secondsString := strings.ToLower(stringArr[2])
	var numberOfSeconds string
	var amOrPm string
	if strings.Contains(secondsString, "pm") { //PM
		numberOfSeconds = strings.Split(secondsString, "pm")[0]
		amOrPm = "pm"
	} else if strings.Contains(secondsString, "am") {
		numberOfSeconds = strings.Split(secondsString, "am")[0]
		amOrPm = "am"
	} else {
		return "Invalid Time Format, please check input...\n"
	}
	if amOrPm == "pm" {
		if hourInt > 0 && hourInt < 12 {
			hourInt = 12 + hourInt
		}
	} else {
		if hourInt == 12 {
			hourInt = 00
		}
	}
	hourStr := fmt.Sprintf("%02d", hourInt)
	var arrayOfTimeIntegers = []string{hourStr, minuteStr, numberOfSeconds + "\n"}
	// return "24 hour format is " + strings.Join(arrayOfTimeIntegers, ":")
	return strings.Join(arrayOfTimeIntegers, ":")
}
